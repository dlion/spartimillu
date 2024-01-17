package server

import (
	"fmt"
	"net/http"
)

type Spartimillu struct{}

func (s *Spartimillu) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("%s / %s\n", r.Method, r.Proto)
	fmt.Printf("%s / %s\n", r.Method, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.Header.Get("User-Agent"))
	fmt.Printf("Accept: %+v\n", r.Header.Get("Accept"))

	fmt.Fprint(w, "ok")
}
