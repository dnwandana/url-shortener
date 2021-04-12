package main

import (
	"fmt"
	"log"

	"github.com/dnwandana/url-shortener/config"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/routes"
	"github.com/dnwandana/url-shortener/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := config.DatabaseConnection()
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}
	fmt.Println("Connected to the database.")

	// TODO: using env value
	urlCollection := db.Collection("urls")
	urlRepo := repository.NewUrlRepository(urlCollection)
	urlService := services.NewUrlService(urlRepo)

	app := fiber.New()
	app.Use(cors.New())

	apiRoute := app.Group("/go")
	routes.UrlRoutes(apiRoute, urlService)

	// TODO: using env value
	log.Fatal(app.Listen(":5000"))
}
