package services

import (
	"github.com/dnwandana/url-shortener/entities"
	"github.com/dnwandana/url-shortener/models"
)

type UrlService interface {
	// List method which returns all shortUrls belonging to that user.
	List(userID string) (*[]models.UrlResponse, error)

	// Create method is used to create a new shortUrl.
	Create(url *entities.Url) (*entities.Url, error)

	// Get method is used to get a specific shortUrl.
	Get(id string) (*entities.Url, error)

	// Update method is used to update the existing shortUrl from database.
	Update(id string, url *entities.Url) (*entities.Url, error)

	// Remove method is used to delete the existing shortUrl from database.
	Remove(id string) error
}
