package main

import ("time"
"net/http"
)

type Data struct{
	Name string
	Points int
	LastUpdated string
}

var Members = make(map[int]Data)

func main() {

	mux:=http.NewServeMux()

	mux.HandleFunc("/{$}",homeHandler)


	
}

func homeHandler(w http.ResponseWriter, r *http.Request){

	
}