package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
)

// Request запрос на сокращение
type Request struct {
	URL string `json:"url"`
}

// Response ответ на запрос на сокращение
type Response struct {
	Result string `json:"result"`
}

// APIShortenURL используется для сокращения длинных URL до нескольких символов
func (h *Handler) APIShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		request := Request{}
		if err := json.Unmarshal(b, &request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userCookie, err := r.Cookie("user_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.Storage.Put(r.Context(), storages.Record{
			User: userCookie.Value,
			URL:  request.URL,
		})

		if err != nil {
			var pge *pgconn.PgError
			if errors.As(err, &pge) && pge.Code == pgerrcode.UniqueViolation {
				record, err := h.Storage.GetByOriginURL(r.Context(), request.URL)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				res, err := h.formatResult(record.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				w.Write(res)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res, err := h.formatResult(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func (h *Handler) formatResult(id uint64) ([]byte, error) {
	resultURL := h.BaseURL + "/" + strconv.FormatUint(id, 10)
	response := Response{Result: resultURL}
	return json.Marshal(response)
}
