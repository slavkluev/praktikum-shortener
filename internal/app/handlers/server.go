package handlers

import (
	"net/http"
)

type storage interface {
	Get(id uint64) (string, error)
	Put(data string) (uint64, error)
}

type Server struct {
	Storage storage
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.ShortenUrl(w, r)
	case http.MethodGet:
		s.GetOriginalUrl(w, r)
	default:
		s.NotFound(w, r)
	}
}
