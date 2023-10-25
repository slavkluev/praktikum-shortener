package handlers

import (
	"log"

	"github.com/go-chi/chi/v5"

	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
)

const (
	// Адрес сервиса
	baseURL = "http://localhost:8080"

	// Путь к файлу для хранения сокращенных URL
	fileStoragePath = "db.txt"

	// Частота синхронизации данных в файле (в секундах)
	fileStorageSyncTime = 5
)

func Example() {
	// Создаем хранилище, которое будет использоваться для хранения сокращенных URL
	storage, err := storages.CreateSimpleStorage(fileStoragePath, fileStorageSyncTime)
	if err != nil {
		log.Fatal(err)
	}

	h := &Handler{
		Mux:     chi.NewMux(),
		Storage: storage,
		BaseURL: baseURL,
	}

	// Слайс с Middlewares
	var mws []Middleware

	// Добавляем обработчики на эндпоинты
	h.Get("/ping", applyMiddlewares(h.Ping(), mws))
	h.Get("/{id}", applyMiddlewares(h.GetOriginalURL(), mws))
	h.Get("/api/user/urls", applyMiddlewares(h.GetAllUrls(), mws))
	h.Delete("/api/user/urls", applyMiddlewares(h.DeleteUrls(), mws))
	h.Post("/", applyMiddlewares(h.ShortenURL(), mws))
	h.Post("/api/shorten", applyMiddlewares(h.APIShortenURL(), mws))
	h.Post("/api/shorten/batch", applyMiddlewares(h.APIShortenBatch(), mws))
	h.NotFound(applyMiddlewares(h.ShowNotFoundPage(), mws))
}
