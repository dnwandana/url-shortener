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

func (s *userService) Create(user *models.User) (*models.User, error) {
	return s.userRepository.Insert(user)
}

func (s *userService) Find(field, value string) (*models.User, error) {
	return s.userRepository.Search(field, value)
}
