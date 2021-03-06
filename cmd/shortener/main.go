package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/slavkluev/praktikum-shortener/internal/app/handlers"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
	flag.Parse()

	storage, err := storages.CreateSimpleStorage(cfg.FileStoragePath, 5)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandler(storage, cfg.BaseURL)
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-c
		storage.Close()
		server.Close()
	}()

	log.Fatal(server.ListenAndServe())
}
