package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// URL struct represents the URL entity.
type URL struct {
	ObjectID  primitive.ObjectID `bson:"_id"`
	ID        string             `bson:"id"`
	URL       string             `bson:"url"`
	SecretKey string             `bson:"secret_key"`
	ExpireAt  time.Time          `bson:"expireAt"`
}
