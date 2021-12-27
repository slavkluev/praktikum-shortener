package handlers

import (
	storages "github.com/slavkluev/praktikum-shortener/internal/app/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_ShortenUrl(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		shortURL    string
	}
	tests := []struct {
		name    string
		request string
		body    string
		storage storage
		want    want
	}{
		{
			name: "simple test #1",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  201,
				shortURL:    "http://example.com/1002",
			},
			request: "/",
			body:    "test1.ru",
		},
		{
			name: "empty body #1",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  201,
				shortURL:    "http://example.com/1002",
			},
			request: "/",
			body:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			server := Server{Storage: tt.storage}
			server.ServeHTTP(w, request)
			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			shortURL, err := ioutil.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, tt.want.shortURL, string(shortURL))
		})
	}
}
