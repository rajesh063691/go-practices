package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// create the mux instance
// create the handler
// create the server
// start the server
type Teachers struct {
	Id      int
	Name    string
	Subject string
}

type Students struct {
	Id     int
	Name   string
	Degree string
}

func main_mux() {

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/teachers", teachersHandler)
	// create new type which implements Handler interface
	student := Students{}
	mux.Handle("/v1/students", student)

	// create the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

// function to gives the list of teachers
func teachersHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received for teachers = %+v \n", req)
	teacherData := []Teachers{
		{Id: 1, Name: "John", Subject: "Maths"},
		{Id: 2, Name: "Doe", Subject: "Science"},
		{Id: 3, Name: "Smith", Subject: "English"},
	}

	data, err := json.Marshal(teacherData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write(data)
	log.Printf("Response successfully sent for teachers =%+v \n", req)
}

func (s Students) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request received for students = %+v \n", req)
	studentData := []Students{
		{Id: 1, Name: "John", Degree: "B.Tech"},
		{Id: 2, Name: "Doe", Degree: "M.Tech"},
		{Id: 3, Name: "Smith", Degree: "MCA"},
	}

	data, err := json.Marshal(studentData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write(data)
	log.Printf("Response successfully sent for students= %+v \n", req)
}
