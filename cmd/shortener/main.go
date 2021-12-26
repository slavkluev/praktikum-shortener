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

		resultUrl := r.Host + "/" + strconv.FormatUint(id, 10)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(201)
		w.Write([]byte(resultUrl))
	case http.MethodGet:
		rawId := strings.TrimPrefix(r.URL.Path, "/")
		id, err := strconv.ParseUint(rawId, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		originUrl, err := s.shortener.GetOriginURL(id)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Location", originUrl)
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
		Addr:    ":8080",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
