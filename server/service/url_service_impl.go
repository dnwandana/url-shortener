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
	var id string
	var expireAt time.Time

	err := request.Validate()
	exception.PanicIfNeeded(exception.BadRequestError{
		Message: err.Error(),
	})

	if request.ID == "" {
		id = util.GenerateNanoID(7)
	} else {
		id = request.ID
		urlExist, _ := service.URLRepository.FindByID(id)
		if urlExist != nil {
			exception.PanicIfNeeded(exception.BadRequestError{
				Message: "custom id already used",
			})
		}
	}

	if request.TTL == "" {
		expireAt = time.Now().Add(time.Duration(720) * time.Hour)
	} else {
		var hour = time.Duration(1) * time.Hour
		var week = time.Duration(168) * time.Hour
		var month = time.Duration(720) * time.Hour

		switch request.TTL {
		case "hour":
			expireAt = time.Now().Add(hour)
		case "week":
			expireAt = time.Now().Add(week)
		case "month":
			expireAt = time.Now().Add(month)
		default:
			expireAt = time.Now().Add(month)
		}
	}

	shortUrl := fmt.Sprintf("%s/go/%s", os.Getenv("DOMAIN"), id)
	secret_key := util.GenerateNanoID(7)
	url := entity.URL{
		ID:        id,
		URL:       request.URL,
		SecretKey: secret_key,
		ExpireAt:  expireAt,
	}

	err = service.URLRepository.Insert(&url)
	exception.PanicIfNeeded(err)

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
	data, err := service.URLRepository.FindByID(id)
	if err != nil {
		return ""
	}

	return data.URL
}

func (service *urlServiceImpl) Delete(id, secret_key string) {
	url, _ := service.URLRepository.FindByID(id)
	if url == nil {
		exception.PanicIfNeeded(exception.BadRequestError{
			Message: "no url deleted",
		})
	} else {
		if url.SecretKey != secret_key {
			exception.PanicIfNeeded(exception.BadRequestError{
				Message: "wrong secret_key",
			})
		}

		err := service.URLRepository.Delete(id)
		exception.PanicIfNeeded(err)
	}
}
