package http

import "net/http"

// Ping используется для проверки доступности сервиса
func (h *RecordHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := h.recordUsecase.Ping(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
