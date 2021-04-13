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
	db, _ := config.DatabaseConnection()

	// TODO: using env value
	urlCollection := db.Collection("urls")
	urlRepo := repository.NewUrlRepository(urlCollection)
	urlService := services.NewUrlService(urlRepo)

	// TODO: using env value
	userCollection := db.Collection("users")
	userRepo := repository.NewUserRepository(userCollection)
	userService := services.NewUserService(userRepo)

	app := fiber.New()
	app.Use(cors.New())

	apiRoute := app.Group("/go")
	routes.UserRoutes(apiRoute, userService)
	routes.UrlRoutes(apiRoute, urlService)

	// TODO: using env value
	log.Fatal(app.Listen(":5000"))
}
