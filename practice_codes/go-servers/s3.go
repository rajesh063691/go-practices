package main

import (
	"log"
	"net/http"
)

func main_3() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		header := rw.Header()
		header.Set("Content-Type", "application/json")
		header.Set("Date", "02/03/2022")

		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"status":"FAILURE"}`))
	})

	log.Fatal(http.ListenAndServe(":9000", mux))
}
