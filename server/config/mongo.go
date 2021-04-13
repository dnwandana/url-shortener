package config

import (
	"context"
	"strconv"
	"time"

	"github.com/dnwandana/url-shortener/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseConnection() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	minPoolSize, err := strconv.Atoi(Env("MONGO_MIN_POOL"))
	utils.Log("=> minPoolSize error: ", err)

	maxPoolSize, err := strconv.Atoi(Env("MONGO_MAX_POOL"))
	utils.Log("=> maxPoolSize error: ", err)

	maxConnIdle, err := strconv.Atoi(Env("MONGO_MAX_CONN_IDLE"))
	utils.Log("=> maxConnIdle error: ", err)

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
