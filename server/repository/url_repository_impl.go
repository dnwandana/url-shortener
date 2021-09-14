package repository

import (
	"context"
	"errors"

	"github.com/dnwandana/url-shortener/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type urlRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewURLRepository(database *mongo.Database) URLRepository {
	return &urlRepositoryImpl{
		Collection: database.Collection("urls"),
	}
}

func (r *urlRepositoryImpl) Insert(url *entity.URL) error {
	_, err := r.Collection.InsertOne(context.Background(), url)
	if err != nil {
		return err
	}

	return nil
}

func (r *urlRepositoryImpl) FindByID(id string) (*entity.URL, error) {
	var url *entity.URL
	err := r.Collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&url)
	if err != nil {
		return nil, err
	}

	if url.ID == "" {
		return nil, errors.New("No URL found")
	}

	return url, nil
}

func (r *urlRepositoryImpl) Delete(id string) error {
	result, err := r.Collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("No URL deleted")
	}

	return nil
}
