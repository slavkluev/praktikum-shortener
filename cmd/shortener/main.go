package main

import (
	"github.com/slavkluev/praktikum-shortener/internal/app/handlers"
	"github.com/slavkluev/praktikum-shortener/internal/app/storage"
	"log"
	"net/http"
)

func main() {
	myServer := &handlers.Server{
		Storage: storage.CreateShortener(1000),
	}
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: myServer,
	}
	log.Fatal(server.ListenAndServe())
}
