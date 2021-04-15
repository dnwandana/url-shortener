package repository

import (
	"context"
	"errors"
	"time"

	"github.com/dnwandana/url-shortener/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlRepository interface {
	FetchUrls() (*[]models.Url, error)
	InsertUrl(url *models.Url) (*models.Url, error)
	FetchUrl(id string) (*models.Url, error)
	UpdateUrl(id string, url *models.Url) (*models.Url, error)
	DeleteUrl(id string) error
}

type urlRepository struct {
	Collection *mongo.Collection
}

func NewUrlRepository(collection *mongo.Collection) UrlRepository {
	return &urlRepository{
		Collection: collection,
	}
}

func (r *urlRepository) FetchUrls() (*[]models.Url, error) {
	var urls []models.Url
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var url models.Url
		err = cursor.Decode(&url)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return &urls, nil
}

func (r *urlRepository) InsertUrl(url *models.Url) (*models.Url, error) {
	url.ObjectID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepository) FetchUrl(id string) (*models.Url, error) {
	var url *models.Url
	err := r.Collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepository) UpdateUrl(id string, url *models.Url) (*models.Url, error) {
	filter := bson.M{"id": id}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "id", Value: url.ID},
				{Key: "title", Value: url.Title},
				{Key: "url", Value: url.URL},
				{Key: "updatedAt", Value: time.Now()},
			},
		},
	}
	result, err := r.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("no url updated")
	}
	return url, nil
}

func (r *urlRepository) DeleteUrl(id string) error {
	result, err := r.Collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no url deleted")
	}
	return nil
}
