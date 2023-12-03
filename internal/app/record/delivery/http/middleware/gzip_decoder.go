package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// GzipDecoder является Middleware для расшифровки Gzip
type GzipDecoder struct{}

// Handle обработка Middleware
func (g GzipDecoder) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		r.Body = gz

		next.ServeHTTP(w, r)
	})
}
