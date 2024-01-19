package server

import (
	"fmt"
	"io"
	"net/http"
	"spartimillu/client"
)

type SpartimilluServer struct {
	client client.Client
}

func NewSpartimilluServer(client client.Client) *SpartimilluServer {
	return &SpartimilluServer{client: client}
}

func (s *SpartimilluServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("%s / %s\n", r.Method, r.Proto)
	fmt.Printf("%s / %s\n", r.Method, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.Header.Get("User-Agent"))
	fmt.Printf("Accept: %+v\n", r.Header.Get("Accept"))

	resp := s.client.ForwardRequest(*r)

	fmt.Printf("Response from server: %s %s\n\n", resp.Proto, resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading the response body", http.StatusInternalServerError)
	}

	stringBody := string(body)
	fmt.Fprint(w, stringBody)
	fmt.Println(stringBody)
}

func (s *SpartimilluServer) HealthCheck() *http.Response {
	fmt.Printf("Performing Health Check\n")

	resp := s.client.HealthCheck()

	fmt.Printf("Response from server %s: %s %s\n\n", resp.Request.Host, resp.Proto, resp.Status)

	return resp
}
