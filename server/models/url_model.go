package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ObjectID  primitive.ObjectID `json:"-" bson:"_id"`
	UserID    string             `json:"userId" bson:"userId"`
	ID        string             `json:"id" bson:"id"`
	Title     string             `json:"title" bson:"title"`
	URL       string             `json:"url" bson:"url"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type UrlResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
