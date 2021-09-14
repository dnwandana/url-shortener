package service

import (
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/repository"
)

type urlServiceImpl struct {
	URLRepository repository.URLRepository
}

func NewURLService(urlRepository repository.URLRepository) URLService {
	return &urlServiceImpl{
		URLRepository: urlRepository,
	}
}

func (service *urlServiceImpl) Create(url *model.URLCreateRequest) (*model.ResponseData, error) {
	panic("implement me")
}

func (service *urlServiceImpl) FindOne(id string) (*model.ResponseMessage, error) {
	panic("implement me")
}

func (service *urlServiceImpl) Delete(id string) (*model.ResponseMessage, error) {
	panic("implement me")
}
