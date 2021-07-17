package controller

import (
	"github.com/dnwandana/url-shortener/entities"
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/services"
	"github.com/dnwandana/url-shortener/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

// SetupRoutes Setup endpoint, parameter, middleware, and handler.
func (controller *UserController) SetupRoutes(app *fiber.App) {
	app.Post("/go/sign-up", controller.signUp())
	app.Post("/go/sign-in", controller.signIn())
}

// signUp handler which handle request for creating a new user.
func (controller *UserController) signUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse data from request body
		var data *models.UserSignUp
		parserErr := ctx.BodyParser(&data)
		// check if there is an error
		if parserErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		// validate the request body
		validationErr := utils.Validate(data)
		// check if there is an error
		if validationErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		// check if the email is already registered
		isEmailExist, _ := controller.UserService.FindByEmail(data.Email)
		// if the email is already registered send a JSON error
		if isEmailExist != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "email already exist",
			})
		}
		// password encryption
		hashedPassword, hashErr := utils.HashPassword(data.Password)
		// check if there is an error
		if hashErr != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      hashErr.Error(),
			})
		}
		// set data to User struct
		user := entities.User{
			Fullname:  data.Fullname,
			Email:     data.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// execute the request
		result, dbErr := controller.UserService.Create(&user)
		// check if there is an error
		if dbErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      dbErr.Error(),
			})
		}
		// send another struct for JSON response
		response := models.UserSignUpResponse{
			ID:       result.ID,
			Fullname: result.Fullname,
		}
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"statusCode": fiber.StatusCreated,
			"user":       response,
		})
	}
}

// signIn handler which handle request for getting cookies and JWT token
func (controller *UserController) signIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse data from request body
		var data *models.UserSignIn
		parserErr := ctx.BodyParser(&data)
		// check if there is an error
		if parserErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      parserErr.Error(),
			})
		}
		// validate the request body
		validationErr := utils.Validate(data)
		// check if there is an error
		if validationErr != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      validationErr,
			})
		}
		// check if the user exist with the given email
		user, _ := controller.UserService.FindByEmail(data.Email)
		// send an error if the user does not exist
		if user == nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "no user exist",
			})
		}
		// compare if the given password is the same as the user password from the database
		isPasswordMatch := utils.VerifyPassword(user.Password, data.Password)
		// send an error if the provided password are not the same
		if !isPasswordMatch {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      "invalid password",
			})
		}
		// generate JWT Token
		token, tokenErr := utils.GenerateJWT(user)
		// check if there is an error
		if tokenErr != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"error":      tokenErr.Error(),
			})
		}
		// set cookies
		userIdCookie := utils.SetCookies("userId", user.ID.Hex())
		jwtCookie := utils.SetCookies("token", token)
		ctx.Cookie(userIdCookie)
		ctx.Cookie(jwtCookie)
		return ctx.SendStatus(fiber.StatusOK)
	}
}
