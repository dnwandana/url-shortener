package entity

import (
	"time"
)

// URL struct represents the URL entity.
type URL struct {
	ID        string    `bson:"id"`
	URL       string    `bson:"url"`
	SecretKey string    `bson:"secret_key"`
	ExpireAt  time.Time `bson:"expireAt"`
}
