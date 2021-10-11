package service

import (
	"github.com/dnwandana/url-shortener/model"
)

type URLService interface {
	Create(request *model.URLCreateRequest) *model.URLResponse

	FindOne(id string) string

	Delete(id, secret_key string)
}
