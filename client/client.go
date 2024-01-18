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
	conf SpartimilluClientConf
}

func NewSpartimilluClient(conf SpartimilluClientConf) *SpartimilluClient {
	return &SpartimilluClient{conf: conf}
}

func (s *SpartimilluClient) ForwardRequest(req http.Request) *http.Response {
	switch req.Method {
	case http.MethodGet:
		return sendGetRequestToAnotherServer(s.conf.address + req.RequestURI)
	case http.MethodPost:
		return sendPostRequestToAnotherServer(s.conf.address+req.RequestURI, req)
	}
	return nil
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
