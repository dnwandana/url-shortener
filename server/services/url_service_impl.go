package services

import (
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/models"
	"github.com/dnwandana/url-shortener/repository"
)

type urlServiceImpl struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(r *repository.UrlRepository) UrlService {
	return &urlServiceImpl{
		urlRepository: *r,
	}
}

func (s *urlServiceImpl) List(userID string) (*[]models.UrlResponse, error) {
	return s.urlRepository.FindAll(userID)
}

func (s *urlServiceImpl) Create(url *entity.Url) (*entity.Url, error) {
	return s.urlRepository.Insert(url)
}

func (s *urlServiceImpl) Get(id string) (*entity.Url, error) {
	return s.urlRepository.FindById(id)
}

func (s *urlServiceImpl) Update(id string, url *entity.Url) (*entity.Url, error) {
	return s.urlRepository.Update(id, url)
}

func (s *urlServiceImpl) Remove(id string) error {
	return s.urlRepository.Delete(id)
}
