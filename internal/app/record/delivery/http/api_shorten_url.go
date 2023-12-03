package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
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
func (h *RecordHandler) APIShortenURL(w http.ResponseWriter, r *http.Request) {
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

	request := Request{}
	if err := json.Unmarshal(b, &request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record := &domain.Record{
		User: user,
		URL:  request.URL,
	}

	err = h.recordUsecase.Store(r.Context(), record)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrConflict):
			res, err := h.createResponse(record)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			w.Write(res)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res, err := h.createResponse(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *RecordHandler) createResponse(record *domain.Record) ([]byte, error) {
	response := Response{
		Result: h.createShortURL(record),
	}
	return json.Marshal(response)
}
