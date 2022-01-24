package handlers

import (
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) ShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		url := string(b)
		id, err := h.Storage.Put(url)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultURL := "http://" + h.BaseURL + "/" + strconv.FormatUint(id, 10)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte(resultURL))
	}
}
