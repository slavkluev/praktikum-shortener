package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/slavkluev/praktikum-shortener/internal/app/handlers"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"log"
	"net/http"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

func main() {
	var cfg = Config{
		ServerAddress:   "localhost:8080",
		BaseURL:         "http://localhost:8080",
		FileStoragePath: "db.txt",
	}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := storages.CreateFileStorage(cfg.FileStoragePath)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	handler := handlers.NewHandler(storage, cfg.BaseURL)
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
