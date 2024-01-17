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
		return sendGetRequestToAnotherServer(s.conf.scheme, s.conf.ip, req)
	case http.MethodPost:
		return sendPostRequestToAnotherServer(s.conf.scheme, s.conf.ip, req)
	}
	return nil
}

func sendGetRequestToAnotherServer(scheme, ip string, req http.Request) *http.Response {
	url := scheme + ip + req.RequestURI
	body, err := http.Get(url)
	if err != nil {
		log.Fatal("Can't read the response body from the GET request")
	}
	return body
}

func sendPostRequestToAnotherServer(scheme, ip string, req http.Request) *http.Response {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Can't read the request's body")
	}
	url := scheme + ip + req.RequestURI
	resp, err := http.Post(url, http.DetectContentType(bodyBytes), req.Body)
	if err != nil {
		log.Fatal("Can't read the response body from the POST request")
	}

	return resp
}
