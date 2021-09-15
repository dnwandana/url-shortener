package service

import (
	"errors"
	"os"
	"time"

	"github.com/dnwandana/url-shortener/entity"
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

func (service *urlServiceImpl) Create(request *model.URLCreateRequest) (*model.URLResponse, error) {
	var id string
	var expireAt time.Time

	if request.ID == "" {
		id, _ = util.GenerateNanoID(7)
	} else {
		id = request.ID
		urlExist, _ := service.URLRepository.FindByID(id)

		if urlExist != nil {
			return nil, errors.New("custom id already used")
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

	shortUrl := os.Getenv("DOMAIN") + id
	secret_key, _ := util.GenerateNanoID(7)
	url := entity.URL{
		ID:        id,
		URL:       request.URL,
		SecretKey: secret_key,
		ExpireAt:  expireAt,
	}

	err := service.URLRepository.Insert(&url)
	if err != nil {
		return nil, err
	}

	response := model.URLResponse{
		ID:        url.ID,
		LongURL:   request.URL,
		ShortURL:  shortUrl,
		SecretKey: secret_key,
		ExpireAt:  expireAt,
	}

	return &response, nil
}

func (service *urlServiceImpl) FindOne(id string) (string, error) {
	data, err := service.URLRepository.FindByID(id)
	if err != nil {
		return "", err
	}

	return data.URL, nil
}

func (service *urlServiceImpl) Delete(id, secret_key string) error {
	url, _ := service.URLRepository.FindByID(id)
	if url == nil {
		return errors.New("no url deleted")
	}

	if url.SecretKey != secret_key {
		return errors.New("wrong secret_key")
	}

	err := service.URLRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
