package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// UserSignIn struct represents form data when trying to sign-in.
type UserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserSignUp struct represents form data when trying to sign-up.
type UserSignUp struct {
	Fullname             string `json:"fullname" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=6"`
	ConfirmationPassword string `json:"confirmationPassword" validate:"required,min=6,eqfield=Password"`
}

// UserSignUpResponse struct represents the JSON response that the user will see after sign-up.
type UserSignUpResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Fullname string             `json:"fullname"`
}
