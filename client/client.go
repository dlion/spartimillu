package client

import (
	"io"
	"log"
	"net/http"
)

type Client interface {
	ForwardRequest(req http.Request) *http.Response
	HealthCheck()
}

type SpartimilluClient struct {
	conf           SpartimilluClientConf
	counter        int
	healthyServers map[string]bool
}

func NewSpartimilluClient(conf SpartimilluClientConf) *SpartimilluClient {
	return &SpartimilluClient{conf: conf, healthyServers: make(map[string]bool)}
}

func (s *SpartimilluClient) ForwardRequest(req http.Request) *http.Response {
	if len(s.healthyServers) == 0 {
		s.HealthCheck()
	}

	index := s.counter % len(s.conf.addresses)
	address := s.conf.addresses[index]
	s.counter++

	if s.healthyServers[address] == true {
		switch req.Method {
		case http.MethodGet:
			return sendGetRequestToAnotherServer(address + req.RequestURI)
		case http.MethodPost:
			return sendPostRequestToAnotherServer(address+req.RequestURI, req)
		}
	}

	return s.ForwardRequest(req)
}

func (s *SpartimilluClient) HealthCheck() {
	for _, address := range s.conf.addresses {
		resp, err := http.Get(address)
		if err == nil && resp.StatusCode == http.StatusOK {
			s.healthyServers[address] = true
		} else {
			s.healthyServers[address] = false
		}
	}
}

func sendGetRequestToAnotherServer(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Can't read the response resp from the GET request")
	}
	return resp
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
