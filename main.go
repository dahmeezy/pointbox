package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name        string
	Points      int
	LastUpdated string
}

var Members = make(map[int]Data)

var id int

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)

	log.Println("Starting Server on Port:4040")
	log.Fatal(http.ListenAndServe(":4040", mux))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	Members[id] = Data{
		Name:   "Zainab",
		Points: 1500,
	}
	id++
	for _, val := range Members {
		fmt.Fprintf(w, "%s: %d", val.Name, val.Points)
	}
}
