package handlers

import (
	"net/http"
)

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad request", 400)
}
