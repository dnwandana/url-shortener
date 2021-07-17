package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User struct represents the user entity.
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Fullname  string             `bson:"fullname"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
