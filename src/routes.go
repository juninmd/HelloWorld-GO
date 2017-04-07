package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"log"
)

type People struct {
	ID string `json:"id"`
}

var people People

func helloWorld(w http.ResponseWriter, req *http.Request) {
	people.ID = "1"
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/teste", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
