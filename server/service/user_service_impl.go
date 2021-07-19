package service

import (
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r *repository.UserRepository) UserService {
	return &userService{
		userRepository: *r,
	}
}

func (s *userService) Create(request *model.UserSignUp) {
	// validate the request body
	validationErr := util.Validate(request)
	// check if there is an error
	util.ReturnErrorIfNeeded(validationErr)

	// check if the email is already registered
	isEmailExist := s.FindByEmail(request.Email)
	// if the email is already registered send a JSON error
	if isEmailExist != nil {
		util.ReturnErrorIfNeeded("Email Already Exist")
	}

	// password encryption
	hashedPassword, hashErr := util.HashPassword(request.Password)
	// check if there is an error
	util.ReturnErrorIfNeeded(hashErr)

	// set data to User struct
	data := entity.User{
		Fullname:  request.Fullname,
		Email:     request.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// execute the request
	txErr := s.userRepository.Insert(&data)
	// check if there is an error
	util.ReturnErrorIfNeeded(txErr)
}

func (s *userService) Login(request *model.UserSignIn) (*entity.User, string) {
	// validate the request body
	validationErr := util.Validate(request)
	// check if there is an error
	util.ReturnErrorIfNeeded(validationErr)

	// check if the user exist with the given email
	user := s.FindByEmail(request.Email)

	// compare if the given password is the same as the user password from the database
	isPasswordMatch := util.VerifyPassword(user.Password, request.Password)
	// send an error if the provided password are not the same
	if !isPasswordMatch {
		util.ReturnErrorIfNeeded("Invalid Password")
	}

	// generate JWT Token
	token, tokenErr := util.GenerateJWT(user)
	// check if there is an error
	util.ReturnErrorIfNeeded(tokenErr)
	// return token
	return user, token
}

func (s *userService) FindByEmail(email string) *entity.User {
	// execute the request
	user, notFound := s.userRepository.FindByEmail(email)
	// check if there is an error
	if notFound != nil {
		util.ReturnErrorIfNeeded("No User Exist")
	}

	// if there are no error, return user
	return user
}

func (s *userService) FetchData(userID string) *model.UserInformation {
	// convert from string into ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	// check if there is an error
	util.ReturnErrorIfNeeded(err)

	// execute the request
	user, notFound := s.userRepository.FindByID(objectID)
	// check if there is an error
	if notFound != nil {
		util.ReturnErrorIfNeeded("No User Exist")
	}

	// send another struct
	data := model.UserInformation{
		Fullname: user.Fullname,
		Email:    user.Email,
	}
	return &data
}
