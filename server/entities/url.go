package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Url struct represents the url entity.
type Url struct {
	ObjectID  primitive.ObjectID `bson:"_id"`
	UserID    string             `bson:"userId"`
	ID        string             `bson:"id"`
	Title     string             `bson:"title"`
	URL       string             `bson:"url"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
