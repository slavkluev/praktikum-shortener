package handlers

import (
	"github.com/go-chi/chi/v5"
)

type storage interface {
	Get(id uint64) (string, error)
	Put(data string) (uint64, error)
}

type Handler struct {
	*chi.Mux
	Storage storage
	BaseURL string
}

func NewHandler(storage storage, baseURL string) *Handler {
	h := &Handler{
		Mux:     chi.NewMux(),
		Storage: storage,
		BaseURL: baseURL,
	}
	h.Get("/{id}", h.GetOriginalURL())
	h.Post("/", h.ShortenURL())
	h.Post("/api/shorten", h.APIShortenURL())
	h.NotFound(h.ShowNotFoundPage())

	return h
}
