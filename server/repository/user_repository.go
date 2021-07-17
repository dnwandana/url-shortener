package repository

import (
	"github.com/dnwandana/url-shortener/entities"
)

type UserRepository interface {
	// Insert method is used to create a new users.
	Insert(user *entities.User) (*entities.User, error)

	// FindByEmail method is used to get a specific user.
	FindByEmail(email string) (*entities.User, error)
}
