package logins

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Created  time.Time          `json:"created" bson:"created"`
	Updated  time.Time          `json:"updated" bson:"updated"`
}
