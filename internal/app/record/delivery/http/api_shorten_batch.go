package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
)

// BatchRequest хранит данные одного элемента в запросе на сокращение
type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginURL     string `json:"original_url"`
}

// BatchResponse хранит данные одного элемента в ответе на запрос
type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

// APIShortenBatch используется для множественного сокращения длинных URL до нескольких символов
func (h *RecordHandler) APIShortenBatch(w http.ResponseWriter, r *http.Request) {
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

	var batchRequests []BatchRequest
	if err := json.Unmarshal(b, &batchRequests); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	batchRecords := make(map[string]*domain.Record)
	for _, batchRequest := range batchRequests {
		batchRecords[batchRequest.CorrelationID] = &domain.Record{
			User: user,
			URL:  batchRequest.OriginURL,
		}
	}

	records := make([]*domain.Record, 0, len(batchRecords))
	for _, record := range batchRecords {
		records = append(records, record)
	}

	err = h.recordUsecase.StoreBatch(r.Context(), records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var batchResponses []BatchResponse
	for correlationID, batchRecord := range batchRecords {
		batchResponses = append(batchResponses, BatchResponse{
			CorrelationID: correlationID,
			ShortURL:      h.createShortURL(batchRecord),
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
