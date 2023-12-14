package http

import (
	"encoding/json"
	"net/http"
)

// StatsResponse хранит статистические данные
type StatsResponse struct {
	UrlsCount  uint64 `json:"urls"`
	UsersCount uint64 `json:"users"`
}

// Stats используется для получения статистики
func (h *RecordHandler) Stats(w http.ResponseWriter, r *http.Request) {
	statistic, err := h.recordUsecase.GetStatistic(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statsResponse := StatsResponse{
		UrlsCount:  statistic.UrlsCount,
		UsersCount: statistic.UsersCount,
	}

	res, err := json.Marshal(statsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
