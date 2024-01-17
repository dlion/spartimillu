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

func (s SpartimilluClient) ForwardRequest(req http.Request) *http.Response {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Can't read the request body")
	}

	url := s.conf.scheme + s.conf.ip + req.RequestURI
	resp, err := http.Post(url, http.DetectContentType(bodyBytes), req.Body)
	if err != nil {
		log.Fatal("Can't read the response body")
	}
	return resp
}
