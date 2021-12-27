package handlers

import (
	storages "github.com/slavkluev/praktikum-shortener/internal/app/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_GetOriginalUrl(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		redirectURL string
	}
	tests := []struct {
		name    string
		request string
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
				statusCode:  307,
				redirectURL: "test2.ru",
			},
			request: "/1001",
		},
		{
			name: "wrong id #2",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  500,
				redirectURL: "",
			},
			request: "/1009",
		},
		{
			name: "empty id #3",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  500,
				redirectURL: "",
			},
			request: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			w := httptest.NewRecorder()

			server := Server{Storage: tt.storage}
			server.ServeHTTP(w, request)
			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			assert.Equal(t, tt.want.redirectURL, result.Header.Get("Location"))
		})
	}
}
