package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	minPoolSize, minPoolErr := strconv.Atoi(Env("MONGO_MIN_POOL"))
	if minPoolErr != nil {
		log.Fatal("=> minPoolSize error:", minPoolSize)
	}

	maxPoolSize, maxPoolErr := strconv.Atoi(Env("MONGO_MAX_POOL"))
	if maxPoolErr != nil {
		log.Fatal("=> maxPoolSize error:", maxPoolErr)
	}

	maxConnIdle, maxConnErr := strconv.Atoi(Env("MONGO_MAX_CONN_IDLE"))
	if maxConnErr != nil {
		log.Fatal("=> maxConnIdle error:", maxConnErr)
	}

	clientOption := options.Client().
		ApplyURI(Env("MONGO_URI")).
		SetMinPoolSize(uint64(minPoolSize)).
		SetMaxPoolSize(uint64(maxPoolSize)).
		SetMaxConnIdleTime(time.Duration(maxConnIdle) * time.Second)

	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	database := client.Database(Env("MONGO_DATABASE"))

	return database, nil
}
