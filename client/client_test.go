package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("should forward a GET request to a specific server", func(t *testing.T) {
		server, address := startTestServer(t, "ok")
		defer server.Close()
		conf := NewSpartimilluClientConf([]string{address}, "/healthcheck")
		client := NewSpartimilluClient(conf)
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		resp := client.ForwardRequest(*req)

		body := getBodyFromResp(t, resp)
		assert.Equal(t, http.MethodGet, resp.Request.Method, "got %v, wanted %v", resp.Request.Method, http.MethodGet)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, "ok", body, "got %v, wanted %v", body, "ok")
	})

	t.Run("should forward a GET request to any server using a round robin algorithm", func(t *testing.T) {
		server1, address1 := startTestServer(t, "server1")
		defer server1.Close()
		server2, address2 := startTestServer(t, "server2")
		defer server2.Close()
		server3, address3 := startTestServer(t, "server3")
		defer server3.Close()
		conf := NewSpartimilluClientConf([]string{address1, address2, address3}, "/healthcheck")
		client := NewSpartimilluClient(conf)
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		resp := client.ForwardRequest(*req)

		body := getBodyFromResp(t, resp)
		assert.Equal(t, http.MethodGet, resp.Request.Method, "got %v, wanted %v", resp.Request.Method, http.MethodGet)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, "server1", body, "got %v, wanted %v", body, "server1")

		resp = client.ForwardRequest(*req)

		body = getBodyFromResp(t, resp)
		assert.Equal(t, http.MethodGet, resp.Request.Method, "got %v, wanted %v", resp.Request.Method, http.MethodGet)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, "server2", body, "got %v, wanted %v", body, "server2")

		resp = client.ForwardRequest(*req)

		body = getBodyFromResp(t, resp)
		assert.Equal(t, http.MethodGet, resp.Request.Method, "got %v, wanted %v", resp.Request.Method, http.MethodGet)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, "server3", body, "got %v, wanted %v", body, "server3")
	})

	t.Run("should perform a healthcheck towards a server", func(t *testing.T) {
		server1, address1 := startTestServer(t, "server1")
		defer server1.Close()
		server2, address2 := startTestServer(t, "server2")
		defer server2.Close()
		server3, address3 := startTestServer(t, "server3")
		defer server3.Close()

		conf := NewSpartimilluClientConf([]string{address1, address2, address3}, "/healthcheck")
		client := NewSpartimilluClient(conf)
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		client.HealthCheck()
		resp := client.ForwardRequest(*req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, server1.URL, resp.Request.URL.Scheme+"://"+resp.Request.Host, "got %v, wanted %v", resp.Request.URL.Scheme+"://"+resp.Request.Host, server1.URL)

		resp = client.ForwardRequest(*req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, server2.URL, resp.Request.URL.Scheme+"://"+resp.Request.Host, "got %v, wanted %v", resp.Request.URL.Scheme+"://"+resp.Request.Host, server2.URL)

		resp = client.ForwardRequest(*req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, server3.URL, resp.Request.URL.Scheme+"://"+resp.Request.Host, "got %v, wanted %v", resp.Request.URL.Scheme+"://"+resp.Request.Host, server3.URL)

		server1.Close()

		client.HealthCheck()
		resp = client.ForwardRequest(*req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, server2.URL, resp.Request.URL.Scheme+"://"+resp.Request.Host, "got %v, wanted %v", resp.Request.URL.Scheme+"://"+resp.Request.Host, server2.URL)

		server2.Close()
		client.HealthCheck()
		resp = client.ForwardRequest(*req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, server3.URL, resp.Request.URL.Scheme+"://"+resp.Request.Host, "got %v, wanted %v", resp.Request.URL.Scheme+"://"+resp.Request.Host, server3.URL)

		resp = client.ForwardRequest(*req)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "got %v, wanted %v", resp.StatusCode, http.StatusOK)
		assert.Equal(t, server3.URL, resp.Request.URL.Scheme+"://"+resp.Request.Host, "got %v, wanted %v", resp.Request.URL.Scheme+"://"+resp.Request.Host, server3.URL)
	})
}

func startTestServer(t *testing.T, bodyResponse string) (*httptest.Server, string) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%s has been called\n", bodyResponse)
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
