package main

import (
	"log"
	"net/http"

	"github.com/rapando/go-web-api/api"
	"github.com/rapando/go-web-api/store"
)

func main() {
	server := &api.PlayerServer{Store: &store.InMemoryPlayerStore{}}
	log.Println("starting api")
	log.Fatal(http.ListenAndServe(":5000", server))
}
