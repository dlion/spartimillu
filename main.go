package main

import (
	"log"
	"net/http"
	"spartimillu/client"
	"spartimillu/server"
)

func main() {
	spartimilluClient := client.NewSpartimilluClient(client.NewSpartimilluClientConf([]string{
		"http://localhost:8080",
		"http://localhost:8081",
	}))
	spartimilluServer := server.NewSpartimilluServer(spartimilluClient)
	log.Fatal(http.ListenAndServe(":80", spartimilluServer))
}
