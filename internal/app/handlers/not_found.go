package handlers

import (
	"net/http"
)

// ShowNotFoundPage используется для отображения ошибки, когда страница не найдена
func (h *Handler) ShowNotFoundPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
