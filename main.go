package main

import (
	"log"
	"net/http"
	"spartimillu/server"
)

func main() {
	spartimillu := &server.Spartimillu{}
	log.Fatal(http.ListenAndServe(":80", spartimillu))
}
