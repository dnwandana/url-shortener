package service

import (
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/util"
	"time"
)

type urlServiceImpl struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(r *repository.UrlRepository) UrlService {
	return &urlServiceImpl{
		urlRepository: *r,
	}
}

func (s *urlServiceImpl) List(userID string) *[]model.UrlResponse {
	// execute the request
	response, txErr := s.urlRepository.FindAll(userID)
	// check if there is an error
	util.ReturnErrorIfNeeded(txErr)
	// send response
	return response
}

func (s *urlServiceImpl) Create(userID string, url *model.UrlData) *model.UrlResponse {
	// validate the request body
	validationErr := util.Validate(url)
	// check if there is an error
	util.ReturnErrorIfNeeded(validationErr)

	// generate nanoid
	nanoid, nanoidErr := util.GenerateNanoID()
	// check if there is an error
	util.ReturnErrorIfNeeded(nanoidErr)

	// set request body to Url struct
	request := entity.Url{
		UserID:    userID,
		ID:        nanoid,
		Title:     url.Title,
		URL:       url.URL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// execute the request
	result, txErr := s.urlRepository.Insert(&request)
	// check if there is an error
	util.ReturnErrorIfNeeded(txErr)

	// return another struct for response
	response := model.UrlResponse{
		ID:        result.ID,
		Title:     result.Title,
		URL:       result.URL,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return &response
}

func (s *urlServiceImpl) Get(id string) (*entity.Url, error) {
	return s.urlRepository.FindById(id)
}

func (s *urlServiceImpl) Update(id string, url *model.UrlData) *model.UrlResponse {
	// validate the request body
	validationErr := util.Validate(url)
	// check if there is an error
	util.ReturnErrorIfNeeded(validationErr)

	// set request body to Url struct
	request := entity.Url{
		ID:        url.ID,
		Title:     url.Title,
		URL:       url.URL,
		UpdatedAt: time.Now(),
	}

	// execute the request
	update, updateErr := s.urlRepository.Update(id, &request)
	// check if there is an error
	util.ReturnErrorIfNeeded(updateErr)

	// execute the request
	result, resultErr := s.Get(update.ID)
	// check if there is an error
	util.ReturnErrorIfNeeded(resultErr)

	// return another struct for response
	response := model.UrlResponse{
		ID:        result.ID,
		Title:     result.Title,
		URL:       result.URL,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return &response
}

func (s *urlServiceImpl) Remove(id string) {
	result := s.urlRepository.Delete(id)
	util.ReturnErrorIfNeeded(result)
}
