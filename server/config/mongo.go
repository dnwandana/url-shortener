package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseConnection will return MongoDB Database.
func DatabaseConnection() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get minimum number of connections allowed
	minPoolSize, minPoolErr := strconv.Atoi(Env("MONGO_MIN_POOL"))
	if minPoolErr != nil {
		log.Fatal("=> minPoolSize error:", minPoolSize)
	}

	// get maximum number of connections allowed
	maxPoolSize, maxPoolErr := strconv.Atoi(Env("MONGO_MAX_POOL"))
	if maxPoolErr != nil {
		log.Fatal("=> maxPoolSize error:", maxPoolErr)
	}

	// get maximum time that connections will remain idle in second
	maxConnIdle, maxConnErr := strconv.Atoi(Env("MONGO_MAX_CONN_IDLE"))
	if maxConnErr != nil {
		log.Fatal("=> maxConnIdle error:", maxConnErr)
	}

	// setting client options
	clientOption := options.Client().
		ApplyURI(Env("MONGO_URI")).
		SetMinPoolSize(uint64(minPoolSize)).
		SetMaxPoolSize(uint64(maxPoolSize)).
		SetMaxConnIdleTime(time.Duration(maxConnIdle) * time.Second)

	// try create new client
	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, err
	}

	// create client connection
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// return database connection
	database := client.Database(Env("MONGO_DATABASE"))

	return database, nil
}
