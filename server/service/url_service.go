package service

import (
	"github.com/dnwandana/url-shortener/model"
)

type URLService interface {
	Create(url *model.URLCreateRequest) (*model.ResponseData, error)

	FindOne(id string) (*model.ResponseMessage, error)

	Delete(id string) (*model.ResponseMessage, error)
}
