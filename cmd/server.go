package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j15r/bookmark-server/cmd/bookmarks"
	"github.com/j15r/bookmark-server/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func main() {
	println("Bookmark server is starting...")
	r := gin.Default()

	// enable database connection
	println("Connecting to database...")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("bookmarks")
	bookmarksCollection := database.Collection("bookmarks")
	bookmarkIndex := mongo.IndexModel{
		Keys: bson.M{
			"id":      bsonx.String("text"),
			"url":     bsonx.String("text"),
			"desc":    bsonx.String("text"),
			"path":    bsonx.String("text"),
			"created": bsonx.DateTime(time.Now().Unix()),
			"updated": bsonx.DateTime(time.Now().Unix()),
			"userId":  bsonx.ObjectID("id"),
		},
		Options: {
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		},
	}

	_, err = bookmarksCollection.Indexes().CreateMany(ctx, bookmarkIndex)
	if err != nil {
		panic(err)
	}

	dbBookmarks := &db.DB{client: client, collection: bookmarksCollection}
	if err != nil {
		panic(err)
	}

	// r.GET("/api/bookmarks", CheckAuth(dbLogins)(BookmarkGetHandler(dbBookmarks)))
	r.GET("/api/bookmarks", bookmarks.GetBookmarksHandler(ctx, dbBookmarks))
	// r.OPTIONS("/api/bookmarks", bookmarks.GetBookmarksHandler(dbBookmarks))
	// r.GET("/api/bookmarks/{id:[a-zA-Z0-9]*}", bookmarks.GetBookmarkHandler(dbBookmarks))
	// r.POST("/api/bookmarks", bookmarks.PostBookmarkHandler(dbBookmarks))
	// r.PUT("/api/bookmarks", bookmarks.PutBookmarkHandler(dbBookmarks))
	// r.DELETE("/api/bookmarks/{id:[a-zA-Z0-9]*}", bookmarks.DeleteBookmarkHandler(dbBookmarks))

	r.Run()
}
