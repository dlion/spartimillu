package main

import (
	"log"
	"net/http"
	"spartimillu/client"
	"spartimillu/server"
)

func main() {
	spartimilluClient := client.NewSpartimilluClient(client.NewSpartimilluClientConf("http://localhost:8080"))
	spartimilluServer := server.NewSpartimilluServer(spartimilluClient)
	log.Fatal(http.ListenAndServe(":80", spartimilluServer))
}
