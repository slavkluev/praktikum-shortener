package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

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
	ServerAddress       string `mapstructure:"server_address"`
	BaseURL             string `mapstructure:"base_url"`
	FileStoragePath     string `mapstructure:"file_storage_path"`
	FileStorageSyncTime int    `mapstructure:"file_storage_sync_time"`
	DatabaseDSN         string `mapstructure:"database_dsn"`
	EnableHTTPS         bool   `mapstructure:"enable_https"`
	CertFile            string `mapstructure:"cert_file"`
	KeyFile             string `mapstructure:"key_file"`
}

func initializeViper() error {
	viper.AutomaticEnv()

	pflag.StringP("config", "c", "config.json", "Config path")
	pflag.StringP("server_address", "a", "localhost:8080", "Server address")
	pflag.StringP("base_url", "b", "http://localhost:8080", "Base URL")
	pflag.StringP("file_storage_path", "f", "db.txt", "File storage path")
	pflag.IntP("file_storage_sync_time", "t", 5, "File storage sync time")
	pflag.StringP("database_dsn", "d", "", "Database DSN")
	pflag.BoolP("enable_https", "s", false, "Enable HTTPS")
	pflag.StringP("cert_file", "", "server.pem", "Cert file")
	pflag.StringP("key_file", "", "server.key", "Key file")

	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return err
	}

	path := viper.GetString("config")
	base := strings.Split(filepath.Base(path), ".")

	viper.SetConfigName(base[0])
	viper.SetConfigType(base[1])
	viper.AddConfigPath(filepath.Dir(path))

	return viper.ReadInConfig()
}

func main() {
	err := initializeViper()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n\n", buildVersion, buildDate, buildCommit)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

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
