package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/slavkluev/praktikum-shortener/internal/app/handlers"
	"github.com/slavkluev/praktikum-shortener/internal/app/middlewares"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
)

const shutdownTimeout = 5 * time.Second

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// Config хранит конфигурацию сервиса
type Config struct {
	ServerAddress       string `env:"SERVER_ADDRESS"`
	BaseURL             string `env:"BASE_URL"`
	FileStoragePath     string `env:"FILE_STORAGE_PATH"`
	FileStorageSyncTime int    `env:"FILE_STORAGE_SYNC_TIME"`
	DatabaseDSN         string `env:"DATABASE_DSN"`
	EnableHTTPS         bool   `env:"ENABLE_HTTPS"`
	CertFile            string `env:"CERT_FILE"`
	KeyFile             string `env:"KEY_FILE"`
}

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n\n", buildVersion, buildDate, buildCommit)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	cfg := parseVariables()

	var storage handlers.Storage
	var err error
	if cfg.DatabaseDSN != "" {
		db, err := sql.Open("pgx", cfg.DatabaseDSN)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		storage, err = storages.CreateDatabaseStorage(db)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		storage, err = storages.CreateSimpleStorage(cfg.FileStoragePath, cfg.FileStorageSyncTime)
		if err != nil {
			log.Fatal(err)
		}
	}

	mws := []handlers.Middleware{
		middlewares.GzipEncoder{},
		middlewares.GzipDecoder{},
		middlewares.NewAuthenticator([]byte("secret key")),
	}

	handler := handlers.NewHandler(storage, cfg.BaseURL, mws)
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}

	go func(cfg Config) {
		if cfg.EnableHTTPS {
			err := server.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("listen and serve tls: %v", err)
			}
		} else {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("listen and serve: %v", err)
			}
		}
	}(cfg)

	log.Printf("listening on %s", cfg.ServerAddress)
	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}

	longShutdown := make(chan struct{}, 1)

	go func() {
		time.Sleep(3 * time.Second)
		longShutdown <- struct{}{}
	}()

	select {
	case <-shutdownCtx.Done():
		log.Fatal(err)
	case <-longShutdown:
		log.Println("finished")
	}
}

func parseVariables() Config {
	var cfg = Config{
		ServerAddress:       "localhost:8080",
		BaseURL:             "http://localhost:8080",
		FileStoragePath:     "db.txt",
		FileStorageSyncTime: 5,
		CertFile:            "server.pem",
		KeyFile:             "server.key",
	}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "File storage path")
	flag.IntVar(&cfg.FileStorageSyncTime, "t", cfg.FileStorageSyncTime, "File storage sync time")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "Database DSN")
	flag.BoolVar(&cfg.EnableHTTPS, "s", cfg.EnableHTTPS, "Enable HTTPS")
	flag.StringVar(&cfg.CertFile, "c", cfg.CertFile, "Cert file")
	flag.StringVar(&cfg.KeyFile, "k", cfg.KeyFile, "Key file")
	flag.Parse()

	return cfg
}
