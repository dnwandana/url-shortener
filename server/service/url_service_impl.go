package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/exception"
	"github.com/dnwandana/url-shortener/model"
	"github.com/dnwandana/url-shortener/repository"
	"github.com/dnwandana/url-shortener/util"
)

type urlServiceImpl struct {
	URLRepository repository.URLRepository
}

func NewURLService(urlRepository *repository.URLRepository) URLService {
	return &urlServiceImpl{
		URLRepository: *urlRepository,
	}
}

func (service *urlServiceImpl) Create(request *model.URLCreateRequest) *model.URLResponse {
	// declare id and expireAt variable
	var id string
	var expireAt time.Time

	// validate the requestBody and check if there is an error
	err := request.Validate()
	if err != nil {
		panic(exception.BadRequestError{
			Message: err.Error(),
		})
	}

	// check if Custom ID is specified in the requestBody
	if request.ID == "" {
		//  if no app will generate random 7 digit alphabet
		id = util.GenerateNanoID(7)
	} else {
		// if Custom ID is specified, app will check if the Custom ID is already used or not
		id = request.ID
		urlExist, _ := service.URLRepository.FindByID(id)
		if urlExist != nil {
			exception.PanicIfNeeded(exception.BadRequestError{
				Message: "Custom back-half already used",
			})
		}
	}

	// check if TTL is specified in the requestBody
	var hour = time.Duration(1) * time.Hour
	var day = time.Duration(24) * time.Hour
	var week = time.Duration(168) * time.Hour
	var month = time.Duration(720) * time.Hour
	var year = time.Duration(8760) * time.Hour

	switch request.TTL {
	case "hour":
		expireAt = time.Now().Add(hour)
	case "day":
		expireAt = time.Now().Add(day)
	case "week":
		expireAt = time.Now().Add(week)
	case "month":
		expireAt = time.Now().Add(month)
	case "year":
		expireAt = time.Now().Add(year)
	default:
		expireAt = time.Now().Add(year)
	}

	// assign short url into variable domain.tld/go/id
	shortUrl := fmt.Sprintf("%s/go/%s", os.Getenv("DOMAIN"), id)
	// generate random secret_key
	secret_key := util.GenerateNanoID(7)
	// assign requestBody into entity.URL
	url := entity.URL{
		ID:        id,
		URL:       request.URL,
		SecretKey: secret_key,
		ExpireAt:  expireAt,
	}

	// execute the request to insert data to database
	err = service.URLRepository.Insert(&url)
	// and check if there is an error
	exception.PanicIfNeeded(err)

	// return response struct
	response := model.URLResponse{
		ID:        url.ID,
		LongURL:   request.URL,
		ShortURL:  shortUrl,
		SecretKey: secret_key,
		ExpireAt:  expireAt,
	}

	return &response
}

func (service *urlServiceImpl) FindOne(id string) string {
	// execute the request to get url
	data, err := service.URLRepository.FindByID(id)
	// if no url found with the given id
	// function will return an empty string
	if err != nil {
		return ""
	}

	// if url found, return the long url
	return data.URL
}

func (service *urlServiceImpl) Delete(id, secret_key string) {
	// check if the requested url is in the database
	url, _ := service.URLRepository.FindByID(id)
	if url == nil {
		// if there are no matched url with the given id
		exception.PanicIfNeeded(exception.BadRequestError{
			Message: "No URL deleted",
		})
	} else {
		// if the given secret_key is not same as the secret_key from database
		if url.SecretKey != secret_key {
			exception.PanicIfNeeded(exception.BadRequestError{
				Message: "Wrong secret_key",
			})
		}

		// execute the request to delete url and check if there is an error
		err := service.URLRepository.Delete(id)
		exception.PanicIfNeeded(err)
	}
}
