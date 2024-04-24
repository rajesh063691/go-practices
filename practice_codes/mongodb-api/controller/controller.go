package controller

import (
	"encoding/json"
	"log"
	database "mongoapi/database_helpher"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Actual controllers - file
// get Single Movie
func GetMyOneMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	id := params["id"]
	oneMovie := database.GetOneMovie(id)

	json.NewEncoder(w).Encode(oneMovie)
}

// get All Movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	allMovies := database.GetAllMovies()
	err := json.NewEncoder(w).Encode(allMovies)
	if err != nil {
		log.Fatal(err)
	}
}

// create a movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	intObjID := database.InsertOneMovie(movie)
	// convert interface into primitive object id
	oid := intObjID.(primitive.ObjectID)
	// //cast interface to string
	// strObjID := fmt.Sprintf("%v", intObjID)
	// // convert string to primitive object id
	// id, err := primitive.ObjectIDFromHex(strObjID)
	// if err != nil {
	// 	fmt.Printf("strObjID=%v", strObjID)
	// }
	movie.ID = oid
	json.NewEncoder(w).Encode(movie)
}

// mark the movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	movieId := params["id"]
	database.UpdateOneMovie(movieId)

	json.NewEncoder(w).Encode(movieId)
}

//delete one movie
func DeleteMyOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	movieId := params["id"]
	database.DeleteOneMovie(movieId)
	json.NewEncoder(w).Encode(movieId)
}

//delete all movie
func DeleteMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	_ = database.DeleteAllMovie()
	json.NewEncoder(w).Encode("All movies deleted successfully!")
}
