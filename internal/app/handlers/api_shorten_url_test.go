package handlers

import (
	storages "github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_ApiShortenUrl(t *testing.T) {
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
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
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
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  500,
				id:          "",
			},
			path: "/api/shorten",
			body: "{",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewHandler(tt.storage, "test.ru")
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
