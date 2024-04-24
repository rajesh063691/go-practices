package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// slice of books
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(r)
	id := params["id"]

	//loop for each of the books to find the id params

	for _, item := range books {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func postBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))

	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			//delete the id first from book
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			//append new collection and set the id as passed through param
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	msg := `{"message":"Delete Request Processed successfully"}`

	json.NewEncoder(w).Encode(msg)
}

func main() {
	books = append(books, Book{ID: "1", Isbn: "ISBN-1", Title: "Book-1", Author: &Author{FirstName: "Author1-FirstName", LastName: "Author1-LastName"}})
	books = append(books, Book{ID: "2", Isbn: "ISBN-2", Title: "Book-2", Author: &Author{FirstName: "Author2-FirstName", LastName: "Author2-LastName"}})
	books = append(books, Book{ID: "3", Isbn: "ISBN-3", Title: "Book-3", Author: &Author{FirstName: "Author3-FirstName", LastName: "Author3-LastName"}})

	//Init Router
	r := mux.NewRouter()

	//Route Handlers/Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", postBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
