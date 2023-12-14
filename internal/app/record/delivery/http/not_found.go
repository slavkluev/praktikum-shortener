package http

import "net/http"

// NotFound используется для отображения ошибки, когда страница не найдена
func (h *RecordHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
