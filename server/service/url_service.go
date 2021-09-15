package service

import (
	"github.com/dnwandana/url-shortener/model"
)

type URLService interface {
	Create(request *model.URLCreateRequest) (*model.URLResponse, error)

	FindOne(id string) (string, error)

	Delete(id, secret_key string) error
}
