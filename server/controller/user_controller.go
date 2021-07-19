package controller

import (
	"github.com/dnwandana/url-shortener/middleware"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/service"
	"github.com/dnwandana/url-shortener/util"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

// SetupRoutes Setup endpoint, parameter, middleware, and handler.
func (controller *UserController) SetupRoutes(app *fiber.App) {
	app.Get("/account", middleware.CookieRequired(), middleware.JWTRequired(), controller.getData)
	app.Post("/account/sign-up", controller.signUp)
	app.Post("/account/sign-in", controller.signIn)
}

// getData handler that handles requests to get user information
func (controller *UserController) getData(ctx *fiber.Ctx) error {
	// getting userId from cookies
	userID := ctx.Cookies("userId")

	// execute the request
	userInfo := controller.UserService.FetchData(userID)

	// succeed getting user information
	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       userInfo,
	})
}

// signUp handle request for creating a new user.
func (controller *UserController) signUp(ctx *fiber.Ctx) error {
	// parse data from request body
	var request *model.UserSignUp
	parserErr := ctx.BodyParser(&request)
	// check if there is an error
	util.ReturnErrorIfNeeded(parserErr)

	// execute the request
	controller.UserService.Create(request)

	// user created
	return ctx.SendStatus(fiber.StatusCreated)
}

// signIn handle request for getting cookies and JWT token
func (controller *UserController) signIn(ctx *fiber.Ctx) error {
	// parse data from request body
	var request *model.UserSignIn
	parserErr := ctx.BodyParser(&request)
	// check if there is an error
	util.ReturnErrorIfNeeded(parserErr)

	// execute the request
	user, token := controller.UserService.Login(request)

	// set cookies
	userIdCookie := util.SetCookies("userId", user.ID.Hex())
	jwtCookie := util.SetCookies("token", token)
	ctx.Cookie(userIdCookie)
	ctx.Cookie(jwtCookie)
	return ctx.SendStatus(fiber.StatusOK)
}
