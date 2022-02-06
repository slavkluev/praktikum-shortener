package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ShortenURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func (h *Handler) GetAllUrls() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userCookie, err := r.Cookie("user_id")

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		records, err := h.Storage.GetByUser(userCookie.Value)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}

		shortenUrls := make([]ShortenURL, 0)

		for _, record := range records {
			shortenUrls = append(shortenUrls, ShortenURL{
				ShortURL:    h.BaseURL + "/" + strconv.FormatUint(record.ID, 10),
				OriginalURL: record.URL,
			})
		}

		res, err := json.Marshal(shortenUrls)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}
}
