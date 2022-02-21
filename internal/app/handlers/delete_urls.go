package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func (h *Handler) DeleteUrls() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var rawIds []string
		if err := json.Unmarshal(b, &rawIds); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var ids []uint64
		for _, rawID := range rawIds {
			id, err := strconv.ParseUint(rawID, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ids = append(ids, id)
		}

		userCookie, err := r.Cookie("user_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		outCh := make(chan uint64)

		go func() {
			var idsToDelete []uint64
			for id := range outCh {
				idsToDelete = append(idsToDelete, id)
			}
			_ = h.Storage.DeleteRecords(r.Context(), idsToDelete)
		}()

		go func() {
			wg := &sync.WaitGroup{}
			for _, id := range ids {
				wg.Add(1)
				go func(id uint64) {
					record, err := h.Storage.Get(r.Context(), id)
					if err != nil {
						wg.Done()
						return
					}

					if record.User == userCookie.Value {
						outCh <- id
					}
					wg.Done()
				}(id)
			}

			wg.Wait()
			close(outCh)
		}()

		w.WriteHeader(http.StatusAccepted)
	}
}
