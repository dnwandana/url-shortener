package config

import (
	"context"
	"os"
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
	minPoolSize, err := strconv.Atoi(os.Getenv("MONGO_MIN_POOL"))
	if err != nil {
		panic(err)
	}

	// get maximum number of connections allowed
	maxPoolSize, err := strconv.Atoi(os.Getenv("MONGO_MAX_POOL"))
	if err != nil {
		panic(err)
	}

	// get maximum time that connections will remain idle in second
	maxConnIdle, err := strconv.Atoi(os.Getenv("MONGO_MAX_CONN_IDLE"))
	if err != nil {
		panic(err)
	}

	// setting client options
	clientOption := options.Client().
		ApplyURI(os.Getenv("MONGO_URI")).
		SetMinPoolSize(uint64(minPoolSize)).
		SetMaxPoolSize(uint64(maxPoolSize)).
		SetMaxConnIdleTime(time.Duration(maxConnIdle) * time.Second)

	// try create new client
	client, err := mongo.NewClient(clientOption)
	if err != nil {
		panic(err)
	}

	// create client connection
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// return database connection
	database := client.Database(os.Getenv("MONGO_DATABASE"))
	return database, nil
}
