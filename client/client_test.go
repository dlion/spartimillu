package client

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	server, address := startTestServer("ok")
	defer server.Close()

	t.Run("should forward a POST request to a specific server", func(t *testing.T) {
		conf := NewSpartimilluClientConf(address)
		client := NewSpartimilluClient(conf)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("hello world for the test server")))

		resp := client.ForwardRequest(*req)

		body := getBodyFromResp(t, resp)
		assert.Equal(t, http.MethodPost, resp.Request.Method, "got %v, wanted %v", resp.Request.Method, http.MethodPost)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, "ok", body, "got %v, wanted %v", body, resp)
	})

	t.Run("should forward a GET request to a specific server", func(t *testing.T) {
		conf := NewSpartimilluClientConf(address)
		client := NewSpartimilluClient(conf)
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		resp := client.ForwardRequest(*req)

		body := getBodyFromResp(t, resp)
		assert.Equal(t, http.MethodGet, resp.Request.Method, "got %v, wanted %v", resp.Request.Method, http.MethodGet)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, "ok", body, "got %v, wanted %v", body, resp)

	})
}

func startTestServer(bodyResponse string) (*httptest.Server, string) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, bodyResponse)
	}))
	return server, server.URL
}

func getBodyFromResp(t *testing.T, resp *http.Response) string {
	t.Helper()

	bodyBytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	return string(bodyBytes)
}
