package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

/*
type Bookmark struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url     string             `json:"url" bson:"url"`
	Desc    string             `json:"desc" bson:"desc"`
	Path    string             `json:"path" bson:"path"`
	Created time.Time          `json:"created" bson:"created"`
	Updated time.Time          `json:"updated" bson:"updated"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId"`
}

type Login struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Created  time.Time          `json:"created" bson:"created"`
	Updated  time.Time          `json:"updated" bson:"updated"`
}
*/

func main() {
	println("Bookmark server is starting...")
	r := gin.Default()

	// enable database connection
	println("Connecting to database...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	r.Run()
}
