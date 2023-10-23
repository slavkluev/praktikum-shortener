package handlers

import (
	"encoding/json"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"io"
	"net/http"
	"strconv"
)

type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginURL     string `json:"original_url"`
}

type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (h *Handler) APIShortenBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var batchRequests []BatchRequest
		if err := json.Unmarshal(b, &batchRequests); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userCookie, err := r.Cookie("user_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var batchRecords []storages.BatchRecord
		for _, batchRequest := range batchRequests {
			batchRecords = append(batchRecords, storages.BatchRecord{
				User:          userCookie.Value,
				URL:           batchRequest.OriginURL,
				CorrelationID: batchRequest.CorrelationID,
			})
		}

		batchRecords, err = h.Storage.PutRecords(r.Context(), batchRecords)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var batchResponses []BatchResponse
		for _, batchRecord := range batchRecords {
			batchResponses = append(batchResponses, BatchResponse{
				CorrelationID: batchRecord.CorrelationID,
				ShortURL:      h.BaseURL + "/" + strconv.FormatUint(batchRecord.ID, 10),
			})
		}

		res, err := json.Marshal(batchResponses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}