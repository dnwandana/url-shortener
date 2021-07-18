package repository

import (
	"context"
	"errors"
	"github.com/dnwandana/url-shortener/entity"
	"github.com/dnwandana/url-shortener/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type urlRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewUrlRepository(collection *mongo.Collection) UrlRepository {
	return &urlRepositoryImpl{
		Collection: collection,
	}
}

func (r *urlRepositoryImpl) FindAll(userID string) (*[]model.UrlResponse, error) {
	var urls []model.UrlResponse
	filter := bson.M{
		"userId": userID,
	}
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var url model.UrlResponse
		err = cursor.Decode(&url)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return &urls, nil
}

func (r *urlRepositoryImpl) Insert(url *entity.Url) (*entity.Url, error) {
	url.ObjectID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepositoryImpl) FindById(id string) (*entity.Url, error) {
	var url *entity.Url
	err := r.Collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *urlRepositoryImpl) Update(id string, url *entity.Url) (*entity.Url, error) {
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

func (r *urlRepositoryImpl) Delete(id string) error {
	result, err := r.Collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no url deleted")
	}
	return nil
}
