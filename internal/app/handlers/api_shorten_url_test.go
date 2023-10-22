package handlers

import (
	"github.com/slavkluev/praktikum-shortener/internal/app/middlewares"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandler_ApiShortenUrl(t *testing.T) {
	file, err := os.CreateTemp("", "db")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	type want struct {
		contentType string
		statusCode  int
		id          string
	}
	tests := []struct {
		name    string
		path    string
		body    string
		storage Storage
		want    want
	}{
		{
			name: "simple test #1",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Records: map[uint64]storages.Record{
					1000: {
						ID:   1000,
						User: "user",
						URL:  "test1.ru",
					},
					1001: {
						ID:   1001,
						User: "user",
						URL:  "test2.ru",
					},
				},
				File: file,
			},
			want: want{
				contentType: "application/json",
				statusCode:  201,
				id:          "1002",
			},
			path: "/api/shorten",
			body: "{\"url\": \"test1.ru\"}",
		},
		{
			name: "empty json #2",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Records: map[uint64]storages.Record{
					1000: {
						ID:   1000,
						User: "user",
						URL:  "test1.ru",
					},
					1001: {
						ID:   1001,
						User: "user",
						URL:  "test2.ru",
					},
				},
				File: file,
			},
			want: want{
				contentType: "application/json",
				statusCode:  201,
				id:          "1002",
			},
			path: "/api/shorten",
			body: "{}",
		},
		{
			name: "wrong json #3",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Records: map[uint64]storages.Record{
					1000: {
						ID:   1000,
						User: "user",
						URL:  "test1.ru",
					},
					1001: {
						ID:   1001,
						User: "user",
						URL:  "test2.ru",
					},
				},
				File: file,
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  400,
				id:          "",
			},
			path: "/api/shorten",
			body: "{",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(tt.storage, "test.ru", []Middleware{
				middlewares.GzipEncoder{},
				middlewares.GzipDecoder{},
				middlewares.NewAuthenticator([]byte("secret key")),
			})
			ts := httptest.NewServer(handler)
			defer ts.Close()

			resp, body := testRequest(t, ts, http.MethodPost, tt.path, strings.NewReader(tt.body))
			defer resp.Body.Close()

			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
			assert.Contains(t, body, tt.want.id)
		})
	}
}
