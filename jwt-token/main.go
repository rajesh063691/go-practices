package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/login", Login).Methods("POST")
	router.HandleFunc("/v1/home", Home).Methods("GET")
	router.HandleFunc("/v1/refresh", Refresh).Methods("GET")

	err := http.ListenAndServe(":4445", router)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}

}
