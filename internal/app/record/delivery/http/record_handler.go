package http

import (
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
)

// RecordHandler обработчик эндпоинтов
type RecordHandler struct {
	baseURL       string
	recordUsecase domain.RecordUsecase
}

// NewRecordHandler создание RecordHandler
func NewRecordHandler(baseURL string, router chi.Router, recordUsecase domain.RecordUsecase) {
	handler := &RecordHandler{
		baseURL:       baseURL,
		recordUsecase: recordUsecase,
	}
	router.Get("/ping", handler.Ping)
	router.Get("/{id}", handler.GetOriginalURL)
	router.Get("/api/user/urls", handler.GetAllUrls)
	router.Delete("/api/user/urls", handler.DeleteUrls)
	router.Post("/", handler.ShortenURL)
	router.Post("/api/shorten", handler.APIShortenURL)
	router.Post("/api/shorten/batch", handler.APIShortenBatch)
	router.NotFound(handler.NotFound)
}

func (h *RecordHandler) createShortURL(record *domain.Record) string {
	return h.baseURL + "/" + strconv.FormatUint(record.ID, 10)
}
