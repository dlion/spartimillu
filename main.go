package main

import (
	"log"
	"net/http"
	"spartimillu/client"
	"spartimillu/server"
	"time"
)

func main() {
	const seconds = 1 * time.Second

	spartimilluClient := client.NewSpartimilluClient(client.NewSpartimilluClientConf([]string{
		"http://localhost:8080",
		"http://localhost:8081",
	}, "/healthcheck"))
	spartimilluServer := server.NewSpartimilluServer(spartimilluClient)

	go doEvery(seconds, spartimilluServer.HealthCheck)

	log.Fatal(http.ListenAndServe(":80", spartimilluServer))
}

func doEvery(d time.Duration, f func()) {
	ticker := time.Tick(d)
	for range ticker {
		go f()
	}
}
