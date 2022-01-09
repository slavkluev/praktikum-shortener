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
}

func NewHandler(storage storage) *Handler {
	h := &Handler{
		Mux:     chi.NewMux(),
		Storage: storage,
	}
	h.Get("/{id}", h.GetOriginalURL())
	h.Post("/", h.ShortenURL())
	h.NotFound(h.ShowNotFoundPage())

	return h
}
