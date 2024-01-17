package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Received request from %s\n", request.RemoteAddr)
		fmt.Printf("%s / %s\n", request.Method, request.Proto)
		fmt.Printf("%s / %s\n", request.Method, request.Proto)
		fmt.Printf("Host: %s\n", request.Host)
		fmt.Printf("User-Agent: %s\n", request.Header.Get("User-Agent"))
		fmt.Printf("Accept: %+v\n\n", request.Header.Get("Accept"))
		fmt.Printf("Replied with a hello message\n")
		fmt.Fprintf(writer, "Hello From Backend Server")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error listening and serve")
	}
}
