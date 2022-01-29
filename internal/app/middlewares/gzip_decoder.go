package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
)

type GzipDecoder struct{}

func (g GzipDecoder) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") != "gzip" {
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
	}
}
