package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// UrlForm struct represents form data when creating or updating from a database.
type UrlForm struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url" validate:"required,url"`
}

// UrlResponse struct represents the JSON response that the user will see after making the request.
type UrlResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
