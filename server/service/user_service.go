package service

import (
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/model"
)

type UserService interface {
	// Create method is used to create a new account.
	Create(request *model.UserSignUp)

	// Login method is used to get JWT Token
	Login(request *model.UserSignIn) (*entity.User, string)

	// FetchData method is used to get user information
	FetchData(userID string) *model.UserInformation
}
