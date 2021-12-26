package main

import (
	"github.com/slavkluev/praktikum-shortener/internal/app"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ShortenerHandler struct {
	shortener *app.Shortener
}

func (s ShortenerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		b, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		url := string(b)
		id, err := s.shortener.ShortenURL(url)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultURL := "http://" + r.Host + "/" + strconv.FormatUint(id, 10)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte(resultURL))
	case http.MethodGet:
		rawID := strings.TrimPrefix(r.URL.Path, "/")
		id, err := strconv.ParseUint(rawID, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		originURL, err := s.shortener.GetOriginURL(id)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", originURL)
		w.WriteHeader(307)
	default:
		http.Error(w, "Bad request", 400)
		return
	}
}

func main() {
	shortener := app.CreateShortener(1000)
	handler := &ShortenerHandler{shortener: &shortener}
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
