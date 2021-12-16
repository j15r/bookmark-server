package bookmarks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j15r/bookmark-server/internal/db"
	"github.com/j15r/bookmark-server/internal/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Bookmark struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url     string             `json:"url" bson:"url"`
	Desc    string             `json:"desc" bson:"desc"`
	Path    string             `json:"path" bson:"path"`
	Created time.Time          `json:"created" bson:"created"`
	Updated time.Time          `json:"updated" bson:"updated"`
	UserId  primitive.ObjectID `json:"userId" bson:"userId"`
}

func GetBookmarksHandler(bookmarkCollection *db.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginId := r.Context().Value("loginId")
		if loginId != nil {
			bsonObjectID := bson.ObjectIdHex(loginId.(string))
			var bookmarks []Bookmark
			err := bookmarkCollection.collection.Find(bson.M{"userId": bsonObjectID}).All(&bookmarks)
			if err != nil {
				helper.SendError(w, err.Error(), http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				body, _ := json.Marshal(bookmarks)
				helper.SendResponse(w, body, http.StatusOK)
			}
		} else {
			helper.SendError(w, "Context has no loginId", http.StatusInternalServerError)
		}
	})
}

func GetBookmarkHandler(bookmarkCollection *db.DB) gin.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginId := r.Context().Value("loginId")
		if loginId != nil {
			bsonObjectID := bson.ObjectIdHex(loginId.(string))
			var bookmark Bookmark
			bookmark, err := bookmarkCollection.getSingleBookmark(r, bsonObjectID)

			if err != nil {
				SendError(w, err.Error(), http.StatusInternalServerError)
			} else {
				body, _ := json.Marshal(bookmark)
				SendResponse(w, body, http.StatusOK)
			}
		} else {
			SendError(w, "Context has no loginId", http.StatusInternalServerError)
		}
	})
}

func PostBookmarkHandler(bookmarkCollection *db.DB) gin.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginId := r.Context().Value("loginId")
		if loginId != nil {
			bsonObjectID := bson.ObjectIdHex(loginId.(string))
			var bookmark Bookmark
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bookmark)

			bookmark.ID = bson.NewObjectId()
			bookmark.Created = time.Now()
			bookmark.Updated = time.Now()
			bookmark.UserId = bsonObjectID

			err := bookmarkCollection.collection.Insert(bookmark)
			if err != nil {
				SendError(w, err.Error(), http.StatusInternalServerError)
			} else {
				body, _ := json.Marshal(bookmark)
				SendResponse(w, body, http.StatusOK)
			}
		} else {
			SendError(w, "Context has no loginId", http.StatusInternalServerError)
		}
	})
}

func DeleteBookmarkHandler(db *db.DB) gin.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})
		if err != nil {
			SendError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})
}

func PutBookmarkHandler(bookmarkCollection *db.DB) gin.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginId := r.Context().Value("loginId")
		if loginId != nil {
			bsonObjectID := bson.ObjectIdHex(loginId.(string))
			vars := mux.Vars(r)
			var bookmark Bookmark
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bookmark)
			bookmark.Updated = time.Now()
			bookmark.UserId = bsonObjectID

			err := bookmarkCollection.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &bookmark})
			if err != nil {
				SendError(w, err.Error(), http.StatusInternalServerError)
			} else {
				body, _ := json.Marshal(bookmark)
				SendResponse(w, body, http.StatusOK)
			}
		} else {
			SendError(w, "Context has no loginId", http.StatusInternalServerError)
		}
	})
}
