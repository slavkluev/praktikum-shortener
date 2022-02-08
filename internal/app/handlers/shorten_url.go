package handlers

import (
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) ShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		userCookie, err := r.Cookie("user_id")

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		url := string(b)
		id, err := h.Storage.Put(r.Context(), storages.Record{
			User: userCookie.Value,
			URL:  url,
		})

		var pge *pgconn.PgError
		if err != nil && errors.As(err, &pge) && pge.Code == pgerrcode.UniqueViolation {
			record, err := h.Storage.GetByOriginURL(r.Context(), url)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			resultURL := h.BaseURL + "/" + strconv.FormatUint(record.ID, 10)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(409)
			w.Write([]byte(resultURL))
			return
		}

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultURL := h.BaseURL + "/" + strconv.FormatUint(id, 10)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte(resultURL))
	}
}
