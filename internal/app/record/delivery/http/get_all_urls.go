package http

import (
	"encoding/json"
	"net/http"
)

// ShortenURL хранит связь сокращенного URL и оригинального URL
type ShortenURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// GetAllUrls используется для получения всех сокращенных URL у пользователя
func (h *RecordHandler) GetAllUrls(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	records, err := h.recordUsecase.GetByUserID(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(records) == 0 {
		http.Error(w, "Not found", http.StatusNoContent)
		return
	}

	var shortenUrls []ShortenURL
	for _, record := range records {
		shortenUrls = append(shortenUrls, ShortenURL{
			ShortURL:    h.createShortURL(&record),
			OriginalURL: record.URL,
		})
	}

	res, err := json.Marshal(shortenUrls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
