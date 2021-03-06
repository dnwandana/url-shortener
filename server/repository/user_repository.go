package repository

import (
	"context"

	"github.com/dnwandana/url-shortener/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Insert(user *models.User) (*models.User, error)
	Search(field, value string) (*models.User, error)
}

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		Collection: collection,
	}
}

// Insert method is used to create a new users.
func (r *userRepository) Insert(user *models.User) (*models.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Search method is used to search users from the database with the given field and value, based on the User struct.
func (r *userRepository) Search(field, value string) (*models.User, error) {
	var user *models.User
	filter := bson.M{
		field: value,
	}
	err := r.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
