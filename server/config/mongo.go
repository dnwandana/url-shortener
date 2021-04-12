package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: using env value
	clientOption := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetMinPoolSize(10).
		SetMaxPoolSize(100).
		SetMaxConnIdleTime(60 * time.Second)

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: using env value
	database := client.Database("urlShortener")

	return database, nil
}
