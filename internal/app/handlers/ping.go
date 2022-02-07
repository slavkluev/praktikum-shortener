package handlers

import (
	"net/http"
)

func (h *Handler) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.Db.Ping(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(200)
	}
}
