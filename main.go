package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	mux.HandleFunc("POST /add/{ID}", addMember)

	log.Println("Starting Server on Port:4040 http://localhost:4040")
	log.Fatal(http.ListenAndServe(":4040", mux))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	for _, val := range members {
		fmt.Fprintf(w, "%s: %d\n", val.Name, val.Points)
	}

}

func addMember(w http.ResponseWriter, r *http.Request) {
	now :=string(time.Now().Format("02-01-2006 03:04PM"))
	name := r.PathValue("ID")
	members[id] = Member{

		Name:   name,
		Points: 1500,
		LastUpdated: now,

	}
	id++
}
