package model

import (
	"time"
)

// UrlData struct represents form data when creating or updating from a database.
type UrlData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url" validate:"required,url"`
}

// UrlResponse struct represents the JSON response that the user will see after making the request.
type UrlResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
