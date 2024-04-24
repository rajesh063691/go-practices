package router

import (
	"mongoapi/controller"
	database "mongoapi/database_helpher"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// initialize database connection
	database.CreateDbConnection()

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controller.GetMyOneMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteMyOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteMyAllMovie).Methods("DELETE")

	return router
}
