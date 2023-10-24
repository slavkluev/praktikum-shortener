package handlers

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
)

// ShortenURL используется для сокращения длинных URL до нескольких символов
func (h *Handler) ShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userCookie, err := r.Cookie("user_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		url := string(b)
		id, err := h.Storage.Put(r.Context(), storages.Record{
			User: userCookie.Value,
			URL:  url,
		})

		if err != nil {
			var pge *pgconn.PgError
			if errors.As(err, &pge) && pge.Code == pgerrcode.UniqueViolation {
				record, err := h.Storage.GetByOriginURL(r.Context(), url)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				resultURL := h.BaseURL + "/" + strconv.FormatUint(record.ID, 10)
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte(resultURL))
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resultURL := h.BaseURL + "/" + strconv.FormatUint(id, 10)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resultURL))
	}
}
