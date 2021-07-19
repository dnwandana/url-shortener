package service

import (
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/model"
)

type UrlService interface {
	// List method which returns all shortUrls belonging to that user.
	List(userID string) *[]model.UrlResponse

	// Create method is used to create a new shortUrl.
	Create(userID string, url *model.UrlData) *model.UrlResponse

	// Get method is used to get a specific shortUrl.
	Get(id string) (*entity.Url, error)

	// Update method is used to update the existing shortUrl from database.
	Update(id string, url *model.UrlData) *model.UrlResponse

	// Remove method is used to delete the existing shortUrl from database.
	Remove(id string)
}
