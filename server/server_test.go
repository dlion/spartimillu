package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	spartimillu := &Spartimillu{}

	t.Run("log the request and returns ok", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		spartimillu.ServeHTTP(resp, req)

		got := resp.Body.String()
		want := "ok"
		assert.Equal(t, want, got, "got %q, want %q", got, want)

	})
}
