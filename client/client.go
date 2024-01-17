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
	return sendPostRequestToAnotherServer(s.conf.scheme, s.conf.ip, req)
}

func sendPostRequestToAnotherServer(scheme, ip string, req http.Request) *http.Response {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Can't read the request's body")
	}
	resp, err := http.Post(scheme+ip+req.RequestURI, http.DetectContentType(bodyBytes), req.Body)
	if err != nil {
		log.Fatal("Can't read the response body")
	}

	return resp
}
