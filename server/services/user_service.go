package services

import (
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/repository"
)

type UserService interface {
	Create(user *models.User) (*models.User, error)
	Find(field, value string) (*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

// Create method is used to create a new users.
func (s *userService) Create(user *models.User) (*models.User, error) {
	return s.userRepository.Insert(user)
}

// Find method is used to search users from the database with the given field and value, based on the User struct.
func (s *userService) Find(field, value string) (*models.User, error) {
	return s.userRepository.Search(field, value)
}
