package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

type InMemoryPS struct {
	store map[string]int
}

func NewInMemoryPS() *InMemoryPS {
	return &InMemoryPS{map[string]int{}}
}

func (i *InMemoryPS) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPS) PostScore(name string) {
	i.store[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPS()}

	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
