package handlers

import (
	"github.com/slavkluev/praktikum-shortener/internal/app/middlewares"
	"github.com/slavkluev/praktikum-shortener/internal/app/storages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	storage := &storages.SimpleStorage{
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
	}
	handler := NewHandler(storage, "test.ru", []Middleware{
		middlewares.GzipEncoder{},
		middlewares.GzipDecoder{},
		middlewares.NewAuthenticator([]byte("secret key")),
	})
	assert.Implements(t, (*http.Handler)(nil), handler)
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	require.NoError(t, err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp, string(respBody)
}
