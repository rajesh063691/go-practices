package db_connect

import (
	"context"
	"fmt"
	"log"
	"time"
	"xml-learning/mongo-connect/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb+srv://rajesh_mongo:rajesh_mongo@cluster0.zab56.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func GetMongoClient() (client *mongo.Client, err error) {
	options := options.Client().ApplyURI(mongoURI)
	client, err = mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	//ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// if we close client here, receiving functions willnot be able to receive
	//defer client.Disconnect(ctx)

	return client, nil
}

func GetAllMovieList(ctx context.Context, database *mongo.Database, showAll bool) (movies []models.Movie) {
	//find one Record

	collection := database.Collection("watchlist")
	if !showAll {
		filter := primitive.M{"watched": true}
		result := collection.FindOne(ctx, filter)
		var movie models.Movie
		err := result.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%+v\n", movie)
		movies = append(movies, movie)
		return movies
	}

	// finc All record
	cursor, err := collection.Find(ctx, primitive.M{})
	if err != nil {
		log.Fatal("Could not connect to collection", err)
	}
	for cursor.Next(ctx) {
		var movie models.Movie
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(ctx)

	return movies
}

func AddMovie(ctx context.Context, database *mongo.Database, movie *models.Movie) (ID primitive.ObjectID, err error) {
	//Add one Record
	dbName := database.Name()
	fmt.Println("Database Name: ", dbName)
	collection := database.Collection("watchlist")
	movie.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, movie)
	if err != nil {
		return ID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
