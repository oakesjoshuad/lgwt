package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

type InMemoryPS struct{}

func (i *InMemoryPS) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPS) PostScore(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPS{}}

	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
