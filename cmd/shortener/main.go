package main

import (
	"github.com/slavkluev/praktikum-shortener/internal/app/handlers"
	"github.com/slavkluev/praktikum-shortener/internal/app/storage"
	"log"
	"net/http"
)

func main() {
	handler := handlers.NewHandler(storage.CreateShortener(1000))
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
