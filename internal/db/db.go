package db

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
}
