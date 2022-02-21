package handlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"net/http"
	"strconv"
)

func (h *Handler) GetOriginalURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawID := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(rawID, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		record, err := h.Storage.Get(r.Context(), id)

		if err != nil {
			if errors.Is(storages.ErrDeleted, err) {
				http.Error(w, err.Error(), http.StatusGone)
				return
			}

			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", record.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
