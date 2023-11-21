package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	ServerAddress       string `env:"SERVER_ADDRESS" json:"server_address"`
	BaseURL             string `env:"BASE_URL" json:"base_url"`
	FileStoragePath     string `env:"FILE_STORAGE_PATH" json:"file_storage_path"`
	FileStorageSyncTime int    `env:"FILE_STORAGE_SYNC_TIME" json:"file_storage_sync_time"`
	DatabaseDSN         string `env:"DATABASE_DSN" json:"database_dsn"`
	EnableHTTPS         bool   `env:"ENABLE_HTTPS" json:"enable_https"`
	CertFile            string `env:"CERT_FILE" json:"cert_file"`
	KeyFile             string `env:"KEY_FILE" json:"key_file"`
	Config              string `env:"CONFIG"`
}

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n\n", buildVersion, buildDate, buildCommit)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	cfg, err := parseVariables()
	if err != nil {
		log.Fatal(err)
	}

	var storage handlers.Storage
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
	}(*cfg)

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

func parseVariables() (*Config, error) {
	cfg := &Config{
		ServerAddress:       "localhost:8080",
		BaseURL:             "http://localhost:8080",
		FileStoragePath:     "db.txt",
		FileStorageSyncTime: 5,
		CertFile:            "server.pem",
		KeyFile:             "server.key",
	}

	path := getConfigPath()
	if path != "" {
		err := loadConfigFromFile(cfg, path)
		if err != nil {
			return nil, err
		}
	}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "File storage path")
	flag.IntVar(&cfg.FileStorageSyncTime, "t", cfg.FileStorageSyncTime, "File storage sync time")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "Database DSN")
	flag.BoolVar(&cfg.EnableHTTPS, "s", cfg.EnableHTTPS, "Enable HTTPS")
	flag.StringVar(&cfg.CertFile, "cert", cfg.CertFile, "Cert file")
	flag.StringVar(&cfg.KeyFile, "key", cfg.KeyFile, "Key file")
	flag.StringVar(&cfg.Config, "c", cfg.Config, "Config")
	flag.StringVar(&cfg.Config, "config", cfg.Config, "Config")
	flag.Parse()

	return cfg, nil
}

func getConfigPath() string {
	path := os.Getenv("CONFIG")

	for i := 1; i < len(os.Args); i += 2 {
		if os.Args[i] == "-c" || os.Args[i] == "-config" {
			path = os.Args[i+1]
		}
	}

	return path
}

func loadConfigFromFile(cfg *Config, filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, cfg)
}
