package repository

import (
	"github.com/dnwandana/url-shortener/entity"
)

type UserRepository interface {
	// Insert method is used to create a new users.
	Insert(user *entity.User) (*entity.User, error)

	// FindByEmail method is used to get a specific user.
	FindByEmail(email string) (*entity.User, error)
}
