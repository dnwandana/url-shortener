package repository

import (
	"context"
	"github.com/dnwandana/url-shortener/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRepositoryImpl{
		Collection: collection,
	}
}

func (r *UserRepositoryImpl) Insert(user *entities.User) (*entities.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	var user *entities.User
	filter := bson.M{
		"email": email,
	}
	err := r.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
