package repository

import (
	"github.com/dnwandana/url-shortener/entities"
	"github.com/dnwandana/url-shortener/models"
)

type UrlRepository interface {
	// FindAll method which returns all shortUrls belonging to that user.
	FindAll(userID string) (*[]models.UrlResponse, error)

	// Insert method is used to create a new shortUrl.
	Insert(url *entities.Url) (*entities.Url, error)

	// FindById method is used to get a specific shortUrl.
	FindById(id string) (*entities.Url, error)

	// Update method is used to update the existing shortUrl from database.
	Update(id string, url *entities.Url) (*entities.Url, error)

	// Delete method is used to delete the existing shortUrl from database.
	Delete(id string) error
}
