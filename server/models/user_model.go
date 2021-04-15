package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Fullname  string             `json:"fullname" bson:"fullname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"-" bson:"createdAt"`
	UpdatedAt time.Time          `json:"-" bson:"updatedAt"`
}

type UserSignUp struct {
	Fullname             string `json:"fullname"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmationPassword"`
}

type UserSignUpResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Fullname string             `json:"fullname"`
}
