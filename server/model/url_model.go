package model

import (
	"time"
)

// URLCreateRequest struct represents form data when shortening URL.
type URLCreateRequest struct {
	ID  string `json:"id"`
	URL string `json:"url" validate:"required,url"`
	TTL string `json:"ttl"`
}

// URLResponse struct represents the JSON response after shortening the URL.
type URLResponse struct {
	ID        string    `json:"id"`
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	SecretKey string    `json:"secret_key"`
	ExpireAt  time.Time `json:"expire_at"`
}
