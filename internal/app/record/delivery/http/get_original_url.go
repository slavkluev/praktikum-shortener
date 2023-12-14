package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetOriginalURL используется для получения оригинального URL по ID
func (h *RecordHandler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(rawID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	record, err := h.recordUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if record.Deleted {
		http.Error(w, "Gone", http.StatusGone)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Location", record.URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
