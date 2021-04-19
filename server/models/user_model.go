package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Fullname  string             `bson:"fullname"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type UserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserSignUp struct {
	Fullname             string `json:"fullname" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6"`
	ConfirmationPassword string `json:"confirmationPassword" validate:"required,min=6,eqfield=Password"`
}

type UserSignUpResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Fullname string             `json:"fullname"`
}
