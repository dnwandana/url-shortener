package service

import (
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/model"
)

type UrlService interface {
	// List method which returns all shortUrls belonging to that user.
	List(userID string) (*[]model.UrlResponse, error)

	// Create method is used to create a new shortUrl.
	Create(url *entity.Url) (*entity.Url, error)

	// Get method is used to get a specific shortUrl.
	Get(id string) (*entity.Url, error)

	// Update method is used to update the existing shortUrl from database.
	Update(id string, url *entity.Url) (*entity.Url, error)

	// Remove method is used to delete the existing shortUrl from database.
	Remove(id string) error
}
