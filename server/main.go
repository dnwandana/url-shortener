package main

import (
	"log"

	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/routes"
	"github.com/dnwandana/url-shortener/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// get database connection
	db, dbErr := config.DatabaseConnection()
	if dbErr != nil {
		log.Fatal("=> database error:", dbErr.Error())
	}

	// setup url collection, instantiate repo, and services
	urlCollection := db.Collection(config.Env("URL_COLLECTION"))
	urlRepo := repository.NewUrlRepository(urlCollection)
	urlService := services.NewUrlService(urlRepo)

	// setup user collection, instantiate repo, and services
	userCollection := db.Collection(config.Env("USER_COLLECTION"))
	userRepo := repository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepo)

	// instantitate fiber application
	app := fiber.New()

	// register cors middleware and allow browser expose credentials
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// setup userService and urlService into `/go` endpoint
	apiRoute := app.Group("/go")
	routes.UserRoutes(apiRoute, userService)
	routes.UrlRoutes(apiRoute, urlService)

	// listen to port `:5000` and log any errors
	log.Fatal(app.Listen(":5000"))
}
