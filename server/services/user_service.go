package services

import (
	"github.com/dnwandana/url-shortener/entities"
)

type UserService interface {
	// Create method is used to create a new users.
	Create(user *entities.User) (*entities.User, error)
	// FindByEmail method is used to get a specific user.
	FindByEmail(email string) (*entities.User, error)
}
