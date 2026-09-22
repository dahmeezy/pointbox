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

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", homeHandler)

	log.Println("Starting Server on Port:4000")
	log.Fatal(http.ListenAndServe(":4000",mux))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	for _, val := range Members {
		fmt.Fprintf(w, "%s: %d", val.Name, val.Points)
	}
}
