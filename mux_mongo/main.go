package main

import (
	"context"
	"crypto/tls"
	"log"
	"mux_mongo/usecase"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("env fileloaded successfully")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("error in connecting to mongodb: %v", err)
	}

	// ping the mongodb
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	log.Println("MongoDB Connection is successful!")
}

func main() {
	// disconnect the client
	defer mongoClient.Disconnect(context.Background())
	// create the service instance
	coll := mongoClient.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	service := &usecase.EmployeeService{
		MongoDBCollection: coll,
	}

	// create the mux instance
	r := mux.NewRouter()
	r.HandleFunc("/v1/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/employee", service.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/v1/employee/{id}", service.GetEmpByID).Methods(http.MethodGet)
	r.HandleFunc("/v1/employee/{id}", service.UpdateEmpByID).Methods(http.MethodPut)
	r.HandleFunc("/v1/employee", service.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/v1/employee/{id}", service.DeleteEmpByID).Methods(http.MethodDelete)
	r.HandleFunc("/v1/employee", service.DeleteAll).Methods(http.MethodDelete)

	log.Println("Server is running on port 4444")
	// one way
	// err := http.ListenAndServe(":4444", r)
	// if err != nil {
	// 	log.Fatalf("error in starting the server: %v", err)
	// }

	// 2nd way
	server := &http.Server{
		Addr:      ":4444",
		TLSConfig: &tls.Config{},
		Handler:   r,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error in starting the server: %v", err)
	}

}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("running..."))
}
