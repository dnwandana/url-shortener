package services

import (
	"github.com/dnwandana/url-shortener/entity"
)

type UserService interface {
	// Create method is used to create a new users.
	Create(user *entity.User) (*entity.User, error)
	// FindByEmail method is used to get a specific user.
	FindByEmail(email string) (*entity.User, error)
}
