package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	mux.HandleFunc("POST /delete/{ID}", deleteMember)
	mux.HandleFunc("POST /points/add/1", addPoint)

	log.Println("Starting Server on Port:4040 http://localhost:4040")
	log.Fatal(http.ListenAndServe(":4040", mux))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	for _, val := range members {
		fmt.Fprintf(w, "%s: %d\n", val.Name, val.Points)
	}

}

func addMember(w http.ResponseWriter, r *http.Request) {
	now := string(time.Now().Format("02-01-2006 03:04PM"))
	name := r.PathValue("ID")
	members[id] = Member{

		Name:        name,
		Points:      1500,
		LastUpdated: now,
	}
	id++
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	i,_ := strconv.Atoi(r.PathValue("ID"))
	delete(members, i)
}

func addPoint(w http.ResponseWriter, r *http.Request){

}