package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
)

// Storage интерфейс хранилища
type Storage interface {
	Get(ctx context.Context, id uint64) (storages.Record, error)
	GetByOriginURL(ctx context.Context, originURL string) (storages.Record, error)
	GetByUser(ctx context.Context, userID string) ([]storages.Record, error)
	Put(ctx context.Context, record storages.Record) (uint64, error)
	PutRecords(ctx context.Context, records []storages.BatchRecord) ([]storages.BatchRecord, error)
	Ping(ctx context.Context) error
	DeleteRecords(ctx context.Context, ids []uint64) error
}

// Middleware интерфейс посредника
type Middleware interface {
	Handle(next http.HandlerFunc) http.HandlerFunc
}

// Handler обработчик эндпоинта
type Handler struct {
	*chi.Mux
	Storage Storage
	BaseURL string
}

// NewHandler создание нового обработчика
func NewHandler(storage Storage, baseURL string, middlewares []Middleware) *Handler {
	h := &Handler{
		Mux:     chi.NewMux(),
		Storage: storage,
		BaseURL: baseURL,
	}

	h.Get("/ping", applyMiddlewares(h.Ping(), middlewares))
	h.Get("/{id}", applyMiddlewares(h.GetOriginalURL(), middlewares))
	h.Get("/api/user/urls", applyMiddlewares(h.GetAllUrls(), middlewares))
	h.Delete("/api/user/urls", applyMiddlewares(h.DeleteUrls(), middlewares))
	h.Post("/", applyMiddlewares(h.ShortenURL(), middlewares))
	h.Post("/api/shorten", applyMiddlewares(h.APIShortenURL(), middlewares))
	h.Post("/api/shorten/batch", applyMiddlewares(h.APIShortenBatch(), middlewares))
	h.NotFound(applyMiddlewares(h.ShowNotFoundPage(), middlewares))

	return h
}

func applyMiddlewares(handler http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware.Handle(handler)
	}

	return handler
}
