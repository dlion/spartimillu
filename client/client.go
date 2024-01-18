package client

import (
	"io"
	"log"
	"net/http"
)

type Client interface {
	ForwardRequest(req http.Request) *http.Response
}

type SpartimilluClient struct {
	conf    SpartimilluClientConf
	counter int
}

func NewSpartimilluClient(conf SpartimilluClientConf) *SpartimilluClient {
	return &SpartimilluClient{conf: conf}
}

func (s *SpartimilluClient) ForwardRequest(req http.Request) *http.Response {
	var resp *http.Response

	serverIndex := s.counter % len(s.conf.addresses)

	switch req.Method {
	case http.MethodGet:
		resp = sendGetRequestToAnotherServer(s.conf.addresses[serverIndex] + req.RequestURI)
	case http.MethodPost:
		resp = sendPostRequestToAnotherServer(s.conf.addresses[serverIndex]+req.RequestURI, req)
	}

	s.counter++

	return resp
}

func sendGetRequestToAnotherServer(url string) *http.Response {
	body, err := http.Get(url)
	if err != nil {
		log.Fatal("Can't read the response body from the GET request")
	}
	return body
}

func sendPostRequestToAnotherServer(url string, req http.Request) *http.Response {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Can't read the request's body")
	}
	resp, err := http.Post(url, http.DetectContentType(bodyBytes), req.Body)
	if err != nil {
		log.Fatal("Can't read the response body from the POST request")
	}

	return resp
}
