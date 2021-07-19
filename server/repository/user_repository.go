package repository

import (
	"github.com/dnwandana/url-shortener/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	// Insert method is used to create a new users.
	Insert(user *entity.User) error

	// FindByEmail method is used to get a specific user.
	FindByEmail(email string) (*entity.User, error)

	// FindByID method is used to get a specific user information.
	FindByID(userID primitive.ObjectID) (*entity.User, error)
}
