package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/slavkluev/praktikum-shortener/internal/app/middlewares"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
)

func TestHandler_GetOriginalUrl(t *testing.T) {
	file, err := os.CreateTemp("", "db")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	type want struct {
		contentType string
		statusCode  int
		redirectURL string
	}
	tests := []struct {
		name    string
		path    string
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
				contentType: "text/plain; charset=utf-8",
				statusCode:  307,
				redirectURL: "test2.ru",
			},
			path: "/1001",
		},
		{
			name: "wrong id #2",
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
				statusCode:  404,
				redirectURL: "",
			},
			path: "/1009",
		},
		{
			name: "empty id #3",
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
				contentType: "",
				statusCode:  405,
				redirectURL: "",
			},
			path: "/",
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

			resp, _ := testRequest(t, ts, http.MethodGet, tt.path, strings.NewReader(""))
			defer resp.Body.Close()

			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
			assert.Equal(t, tt.want.redirectURL, resp.Header.Get("Location"))
		})
	}
}
