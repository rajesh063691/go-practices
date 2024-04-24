package database

import (
	"context"
	"fmt"
	"log"
	"mongoapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://rajesh_mongo:rajesh_mongo@cluster0.zab56.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbName = "netflix"

const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongodb
func CreateDbConnection() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("collection instance is ready")
}

// MONGODB Helpers - file

// insert one record
func InsertOneMovie(movie model.Netflix) interface{} {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id: ", inserted.InsertedID)
	return inserted.InsertedID
}

// update one record
func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one record
func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with id: and  deleted count: ", movieId, deleteCount)
}

// delete all record from mongodb
func DeleteAllMovie() int64 {
	filter := bson.M{}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with deleted count: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount

}

// get all movies from database
func GetAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie primitive.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// get single movie

func GetOneMovie(movieId string) model.Netflix {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	var movie model.Netflix
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)
	err = result.Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie fetched successfully")

	return movie
}
