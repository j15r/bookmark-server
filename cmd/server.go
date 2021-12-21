package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j15r/bookmark-server/cmd/bookmarks"
)

func main() {
	println("Bookmark server is starting...")
	r := gin.Default()

	// r.GET("/api/bookmarks", CheckAuth(dbLogins)(BookmarkGetHandler(dbBookmarks)))
	r.GET("/api/bookmarks", bookmarks.GetBookmarks)
	// r.GET("/api/bookmarks/{id:[a-zA-Z0-9]*}", bookmarks.GetBookmarkHandler(dbBookmarks))
	// r.POST("/api/bookmarks", bookmarks.PostBookmarkHandler(dbBookmarks))
	// r.PUT("/api/bookmarks", bookmarks.PutBookmarkHandler(dbBookmarks))
	// r.DELETE("/api/bookmarks/{id:[a-zA-Z0-9]*}", bookmarks.DeleteBookmarkHandler(dbBookmarks))

	r.Run()
}
