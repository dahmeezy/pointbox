package main

import (
	"fmt"
	"log"
	"net/http"
)

type Member struct {
	Name        string
	Points      int
	LastUpdated string
}

var members = make(map[int]Member)

var id int

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)

	log.Println("Starting Server on Port:4040")
	log.Fatal(http.ListenAndServe(":4040", mux))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	members[id] = Member{
		Name:   "Zainab",
		Points: 1500,
	}
	id++
	for _, val := range members {
		fmt.Fprintf(w, "%s: %d", val.Name, val.Points)
	}
}
