package main

import (
	"context"
	"fmt"
	"log"
	"xml-learning/mongo-connect/db_connect"
	"xml-learning/mongo-connect/models"
)

var dbName = "netflix"

func main() {
	fmt.Println("<=======MongoDb Connection Tutorials=========>")
	ctx := context.Background()
	client, err := db_connect.GetMongoClient()
	if err != nil {
		log.Fatal("Could not connect to MongoDB Database")
	}
	defer client.Connect(ctx)
	fmt.Println("Connection success!!!")

	database := client.Database(dbName)

	// get single or All Movies
	// movieLIst := db_connect.GetAllMovieList(ctx, database, true)
	// fmt.Printf("Movie List: %+v\n", movieLIst)

	// Insert one Movie
	movie := models.Movie{
		Movie:   "Testing By Rajesh",
		Watched: true,
	}
	id, err := db_connect.AddMovie(ctx, database, &movie)
	if err != nil {
		log.Fatalf("Error while inserting movie record, err: %v", err)
	}

	fmt.Printf("Record Inserted with ID: %v\n", id)

}
