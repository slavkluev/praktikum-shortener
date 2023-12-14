package http

import (
	"errors"
	"io"
	"net/http"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
)

// ShortenURL используется для сокращения длинных URL до нескольких символов
func (h *RecordHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := string(b)
	record := &domain.Record{
		User: user,
		URL:  url,
	}

	err = h.recordUsecase.Store(r.Context(), record)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrConflict):
			resultURL := h.createShortURL(record)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(resultURL))
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	resultURL := h.createShortURL(record)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resultURL))
}
