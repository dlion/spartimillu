package server

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("should call the client to forward the request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		mockClient := newMockClient()
		spartimillu := NewSpartimilluServer(mockClient)
		mockClient.On("ForwardRequest", mock.Anything).Return(&http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			Body:       io.NopCloser(bytes.NewBufferString("dummy body")),
			Request:    req,
		})

		spartimillu.ServeHTTP(resp, req)

		mockClient.AssertExpectations(t)
		assert.Equal(t, "dummy body", resp.Body.String(), "got %q, want %q", resp.Body.String(), "dummy body")
	})

	t.Run("should call the client to do an health check", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
		resp := httptest.NewRecorder()
		mockClient := newMockClient()
		spartimillu := NewSpartimilluServer(mockClient)
		mockClient.On("HealthCheck").Return(&http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			Request:    req,
		})

		spartimillu.HealthCheck()

		mockClient.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, resp.Code, "got %q, want %q", resp.Code, http.StatusOK)
	})
}

type MockClient struct {
	mock.Mock
}

func newMockClient() *MockClient { return &MockClient{} }

func (m *MockClient) ForwardRequest(req http.Request) *http.Response {
	args := m.Called(req)
	return args.Get(0).(*http.Response)
}

func (m *MockClient) HealthCheck() {
	m.Called()
}
