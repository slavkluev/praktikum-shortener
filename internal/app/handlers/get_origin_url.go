package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetOriginalURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawID := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(rawID, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		record, err := h.Storage.Get(r.Context(), id)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", record.URL)
		w.WriteHeader(307)
	}
}
