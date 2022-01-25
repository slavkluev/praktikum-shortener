package handlers

import (
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandler_GetOriginalUrl(t *testing.T) {
	file, err := ioutil.TempFile("", "db")
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
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
				},
				File: file,
			},
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  500,
				redirectURL: "",
			},
			path: "/1009",
		},
		{
			name: "empty id #3",
			storage: &storages.SimpleStorage{
				Start: 1002,
				Urls: map[uint64]string{
					1000: "test1.ru",
					1001: "test2.ru",
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
			handler := NewHandler(tt.storage, "test.ru")
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
