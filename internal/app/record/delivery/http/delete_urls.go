package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// DeleteUrls используется для удаления сокращенных URL
func (h *RecordHandler) DeleteUrls(w http.ResponseWriter, r *http.Request) {
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

	err = h.recordUsecase.DeleteBatch(r.Context(), ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
