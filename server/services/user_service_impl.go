package services

import (
	"github.com/dnwandana/url-shortener/entities"
	"github.com/dnwandana/url-shortener/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r *repository.UserRepository) UserService {
	return &userService{
		userRepository: *r,
	}
}

func (s *userService) Create(user *entities.User) (*entities.User, error) {
	return s.userRepository.Insert(user)
}

func (s *userService) FindByEmail(email string) (*entities.User, error) {
	return s.userRepository.FindByEmail(email)
}
