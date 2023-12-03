package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	httpDelivery "github.com/slavkluev/praktikum-shortener/internal/app/record/delivery/http"
	"github.com/slavkluev/praktikum-shortener/internal/app/record/delivery/http/middleware"
	recordMemoryRepo "github.com/slavkluev/praktikum-shortener/internal/app/record/repository/memory"
	recordUcase "github.com/slavkluev/praktikum-shortener/internal/app/record/usecase"
)

const (
	// Базовый адрес
	baseURL = "http://localhost:8080"

	// Адрес сервера
	serverAddress = "localhost:8080"
)

func Example() {
	recordRepository := recordMemoryRepo.NewMemoryRecordRepository()

	timeoutContext := time.Duration(5) * time.Second
	recordUsecase := recordUcase.NewRecordUsecase(recordRepository, timeoutContext)

	router := chi.NewRouter()

	authenticator := middleware.NewAuthenticator([]byte("secret key"))
	gzipEncoder := middleware.GzipEncoder{}
	gzipDecoder := middleware.GzipDecoder{}

	// Подключаем Middlewares
	router.Use(authenticator.Handle)
	router.Use(gzipEncoder.Handle)
	router.Use(gzipDecoder.Handle)

	httpDelivery.NewRecordHandler(baseURL, router, recordUsecase)

	server := &http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	server.ListenAndServe()
}
