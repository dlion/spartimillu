package main

import (
	"log"
	"net/http"
	"spartimillu/client"
	"spartimillu/server"
	"time"
)

func main() {
	spartimilluClient := client.NewSpartimilluClient(client.NewSpartimilluClientConf([]string{
		"http://localhost:8080",
		"http://localhost:8081",
	}, "/healthcheck"))
	spartimilluServer := server.NewSpartimilluServer(spartimilluClient)

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				spartimilluServer.HealthCheck()
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":80", spartimilluServer))
}
