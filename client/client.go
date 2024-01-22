package client

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client interface {
	ForwardRequest(req http.Request) *http.Response
	HealthCheck()
}

type SpartimilluClient struct {
	conf           SpartimilluClientConf
	counter        int
	healthyServers map[string]bool
	mu             sync.Mutex
}

func NewSpartimilluClient(conf SpartimilluClientConf) *SpartimilluClient {
	return &SpartimilluClient{conf: conf, healthyServers: make(map[string]bool)}
}

func (s *SpartimilluClient) ForwardRequest(req http.Request) *http.Response {
	for {
		if len(s.healthyServers) == 0 {
			s.HealthCheck()
		}

		s.mu.Lock()
		index := s.counter % len(s.conf.addresses)
		address := s.conf.addresses[index]
		s.counter++

		if s.healthyServers[address] == true {
			s.mu.Unlock()
			switch req.Method {
			case http.MethodGet:
				return sendGetRequestToAnotherServer(address + req.RequestURI)
			case http.MethodPost:
				return sendPostRequestToAnotherServer(address+req.RequestURI, req)
			}
		}
		s.mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *SpartimilluClient) HealthCheck() {
	for _, address := range s.conf.addresses {
		resp, err := http.Get(address)

		s.mu.Lock()
		if err == nil && resp.StatusCode == http.StatusOK {
			s.healthyServers[address] = true
		} else {
			s.healthyServers[address] = false
		}
		s.mu.Unlock()
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
