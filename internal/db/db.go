package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func GetDBClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database")

	options := options.Index()
	options.SetUnique(true)
	options.SetBackground(true)
	options.SetSparse(true)
	bookmarkIndex := mongo.IndexModel{
		Keys: bson.M{
			"id":      bsonx.String("id"),
			"url":     bsonx.String("url"),
			"desc":    bsonx.String("desc"),
			"path":    bsonx.String("path"),
			"created": bsonx.DateTime(time.Now().Unix()),
			"updated": bsonx.DateTime(time.Now().Unix()),
			"userId":  bsonx.ObjectID(primitive.NewObjectID()),
		},
		Options: options,
	}

	_, err = client.Database("bookmarks").Collection("bookmarks").Indexes().CreateOne(ctx, bookmarkIndex)
	if err != nil {
		log.Fatal("Could not create index:", err)
		panic(err)
	}

	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}
