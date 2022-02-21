package handlers

import (
	"net/http"
)

func (h *Handler) ShowNotFoundPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
