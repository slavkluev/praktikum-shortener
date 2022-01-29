package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Storage interface {
	Get(id uint64) (string, error)
	Put(data string) (uint64, error)
}

type Middleware interface {
	Handle(next http.HandlerFunc) http.HandlerFunc
}

type Handler struct {
	*chi.Mux
	Storage Storage
	BaseURL string
}

func NewHandler(storage Storage, baseURL string, middlewares []Middleware) *Handler {
	h := &Handler{
		Mux:     chi.NewMux(),
		Storage: storage,
		BaseURL: baseURL,
	}

	h.Get("/{id}", applyMiddlewares(h.GetOriginalURL(), middlewares))
	h.Post("/", applyMiddlewares(h.ShortenURL(), middlewares))
	h.Post("/api/shorten", applyMiddlewares(h.APIShortenURL(), middlewares))
	h.NotFound(applyMiddlewares(h.ShowNotFoundPage(), middlewares))

	return h
}

func applyMiddlewares(handler http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware.Handle(handler)
	}

	return handler
}
