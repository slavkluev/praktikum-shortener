package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/slavkluev/praktikum-shortener/internal/app/handlers"
	"github.com/slavkluev/praktikum-shortener/internal/app/storage"
	"log"
	"net/http"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
}

func main() {
	var cfg = Config{
		ServerAddress: "localhost:8080",
		BaseURL:       "localhost:8080",
	}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandler(storage.CreateShortener(1000), cfg.BaseURL)
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
