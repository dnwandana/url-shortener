package services

import (
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/repository"
)

type UrlService interface {
	ListAllShortUrl(userID string) (*[]models.UrlResponse, error)
	CreateShortUrl(url *models.Url) (*models.Url, error)
	GetShortUrl(id string) (*models.Url, error)
	UpdateShortUrl(id string, url *models.Url) (*models.Url, error)
	DeleteShortUrl(id string) error
}

type urlService struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(r repository.UrlRepository) UrlService {
	return &urlService{
		urlRepository: r,
	}
}

func (s *urlService) ListAllShortUrl(userID string) (*[]models.UrlResponse, error) {
	return s.urlRepository.FetchUrls(userID)
}

func (s *urlService) CreateShortUrl(url *models.Url) (*models.Url, error) {
	return s.urlRepository.InsertUrl(url)
}

func (s *urlService) GetShortUrl(id string) (*models.Url, error) {
	return s.urlRepository.FetchUrl(id)
}

func (s *urlService) UpdateShortUrl(id string, url *models.Url) (*models.Url, error) {
	return s.urlRepository.UpdateUrl(id, url)
}

func (s *urlService) DeleteShortUrl(id string) error {
	return s.urlRepository.DeleteUrl(id)
}
