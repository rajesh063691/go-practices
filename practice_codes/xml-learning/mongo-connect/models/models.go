package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// MovieDetails ...
type Movie struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Movie   string             `json:"movie" bson:"movie"`
	Watched bool               `json:"watched" bson:"watched"`
}
