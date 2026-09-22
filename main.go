package main

import (
	"fmt"
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

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	for _,val:=range Members{
		fmt.Fprintf(w,"%s: %d",val.Name,val.Points)
	}
}
