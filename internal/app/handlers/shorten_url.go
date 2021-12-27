package handlers

import (
	"io"
	"net/http"
	"strconv"
)

func (s *Server) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		url := string(b)
		id, err := s.Storage.Put(url)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultURL := "http://" + r.Host + "/" + strconv.FormatUint(id, 10)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte(resultURL))
	}
}
