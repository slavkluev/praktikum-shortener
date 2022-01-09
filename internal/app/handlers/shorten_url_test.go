package handlers

import (
	storages "github.com/slavkluev/praktikum-shortener/internal/app/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_ShortenUrl(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		id          string
	}
	tests := []struct {
		name    string
		path    string
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
				id:          "1002",
			},
			path: "/",
			body: "test1.ru",
		},
		{
			name: "empty body #2",
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
				id:          "1002",
			},
			path: "/",
			body: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(tt.storage)
			ts := httptest.NewServer(handler)
			defer ts.Close()

			resp, body := testRequest(t, ts, http.MethodPost, tt.path)
			defer resp.Body.Close()

			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
			assert.Contains(t, body, tt.want.id)
		})
	}
}
