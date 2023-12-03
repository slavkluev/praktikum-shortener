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

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
	httpDelivery "github.com/slavkluev/praktikum-shortener/internal/app/record/delivery/http"
	"github.com/slavkluev/praktikum-shortener/internal/app/record/delivery/http/middleware"
	recordMemoryRepo "github.com/slavkluev/praktikum-shortener/internal/app/record/repository/memory"
	recordPostgresRepo "github.com/slavkluev/praktikum-shortener/internal/app/record/repository/postgres"
	recordUcase "github.com/slavkluev/praktikum-shortener/internal/app/record/usecase"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	router := chi.NewRouter()

	authenticator := middleware.NewAuthenticator([]byte("secret key"))
	gzipEncoder := middleware.GzipEncoder{}
	gzipDecoder := middleware.GzipDecoder{}

	router.Use(authenticator.Handle)
	router.Use(gzipEncoder.Handle)
	router.Use(gzipDecoder.Handle)

	var recordRepository domain.RecordRepository
	if cfg.DatabaseDSN != "" {
		db, err := sql.Open("pgx", cfg.DatabaseDSN)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		recordRepository, err = recordPostgresRepo.NewPostgresRecordRepository(db)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		recordRepository = recordMemoryRepo.NewMemoryRecordRepository()
	}

	timeoutContext := time.Duration(5) * time.Second
	recordUsecase := recordUcase.NewRecordUsecase(recordRepository, timeoutContext)

	httpDelivery.NewRecordHandler(cfg.BaseURL, router, recordUsecase)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
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

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("finished")
}
