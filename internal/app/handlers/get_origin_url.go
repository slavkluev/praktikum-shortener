package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rawID := strings.TrimPrefix(r.URL.Path, "/")
		id, err := strconv.ParseUint(rawID, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		originURL, err := s.Storage.Get(id)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", originURL)
		w.WriteHeader(307)
	}
}
