package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllUrls() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userCookie, err := r.Cookie("user_id")

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		records, err := h.Storage.GetByUser(userCookie.Value)

		if err != nil {
			http.Error(w, err.Error(), 204)
			return
		}

		res, err := json.Marshal(records)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}
}
