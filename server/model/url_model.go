package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// URLCreateRequest struct represents form data when shortening URL.
type URLCreateRequest struct {
	ID  string `json:"id"`
	URL string `json:"url" validate:"required,url"`
	TTL string `json:"ttl"`
}

func (request URLCreateRequest) Validate() error {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ID, validation.When(request.ID != "", validation.Length(3, 0))),
		validation.Field(&request.URL, validation.Required, is.URL),
	)

	if err != nil {
		return err
	}

	return nil
}

// URLResponse struct represents the JSON response after shortening the URL.
type URLResponse struct {
	ID        string    `json:"id"`
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	SecretKey string    `json:"secret_key"`
	ExpireAt  time.Time `json:"expire_at"`
}
