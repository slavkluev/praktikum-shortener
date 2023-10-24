package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetOriginalURL используется для получения оригинального URL по ID
func (h *Handler) GetOriginalURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawID := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(rawID, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		record, err := h.Storage.Get(r.Context(), id)

		if record.Deleted {
			http.Error(w, "Gone", http.StatusGone)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", record.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
