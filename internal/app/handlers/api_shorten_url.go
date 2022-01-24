package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (h *Handler) ApiShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		request := Request{}
		if err := json.Unmarshal(b, &request); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		id, err := h.Storage.Put(request.Url)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultURL := "http://" + r.Host + "/" + strconv.FormatUint(id, 10)
		response := Response{Result: resultURL}

		res, err := json.Marshal(response)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(res)
	}
}
