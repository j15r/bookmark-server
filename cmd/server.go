package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j15r/bookmark-server/cmd/bookmarks"
	"github.com/j15r/bookmark-server/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

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

	bookmarkCollection := client.Database("bookmarks").Collection("bookmarks")
	bookmarkIndex := []mongo.IndexModel{
		{
			Keys:   bsonx.Doc{{Key: "id", Value: bsonx.String("text")}},
			Keys:   bsonx.Doc{{Key: "url", Value: bsonx.String("text")}},
			Keys:   bsonx.Doc{{Key: "desc", Value: bsonx.String("text")}},
			Keys:   bsonx.Doc{{Key: "path", Value: bsonx.String("text")}},
			Keys:   bsonx.Doc{{Key: "created", Value: bsonx.DateTime()}},
			Keys:   bsonx.Doc{{Key: "updated", Value: bsonx.DateTime()}},
			UserId: bsonx.Doc{{Key: "userId", Value: bsonx.ObjectID()}},
			Keys:   []string{"id", "url", "desc", "path", "created", "updated", "userid"},
		},
	}

	/* index options
	{
		unique:     true,
		dropDups:   true,
		background: true,
		sparse:     true
	}
	*/

	_, err = bookmarkCollection.Indexes().CreateMany(nil, bookmarkIndex)
	if err != nil {
		panic(err)
	}

	dbBookmarks := &db.DB{client: client, collection: bookmarkCollection}
	if err != nil {
		panic(err)
	}

	// r.GET("/api/bookmarks", CheckAuth(dbLogins)(BookmarkGetHandler(dbBookmarks)))
	r.GET("/api/bookmarks", bookmarks.GetBookmarksHandler(dbBookmarks))
	r.OPTIONS("/api/bookmarks", bookmarks.GetBookmarksHandler(dbBookmarks))
	r.GET("/api/bookmarks/{id:[a-zA-Z0-9]*}", bookmarks.GetBookmarkHandler(dbBookmarks))
	r.POST("/api/bookmarks", bookmarks.PostBookmarkHandler(dbBookmarks))
	r.PUT("/api/bookmarks", bookmarks.PutBookmarkHandler(dbBookmarks))
	r.DELETE("/api/bookmarks/{id:[a-zA-Z0-9]*}", bookmarks.DeleteBookmarkHandler(dbBookmarks))

	r.Run()
}
