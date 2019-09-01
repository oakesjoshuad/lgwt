package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	PostScore(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.postScore(w)
	case http.MethodGet:
		p.getScore(w, r)
	}
}

func (p *PlayerServer) getScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) postScore(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}
	if name == "Floyd" {
		return 10
	}
	return 0
}
