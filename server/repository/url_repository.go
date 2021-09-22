package repository

import "github.com/dnwandana/url-shortener/entity"

type URLRepository interface {
	Insert(url *entity.URL) error

	FindByID(id string) (*entity.URL, error)

	Delete(id string) error
}
