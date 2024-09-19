package services

import (
	"context"
	"time"

	"github.com/KCKT0112/LiteBlog/app/db"
	"github.com/KCKT0112/LiteBlog/app/models"
	"github.com/KCKT0112/LiteBlog/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(user models.User) (*mongo.InsertOneResult, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUID(id string) (*models.User, error)
}

type userService struct {
	collection *mongo.Collection
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		utils.Logger.Error("Error inserting user", zap.Error(err))
		return nil, err
	}
	return result, nil
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	filter := bson.D{{Key: "email", Value: email}}
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		utils.Logger.Error("Error finding user", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// GetUserByID implements UserService.
func (u *userService) GetUserByUID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	filter := bson.D{{Key: "uid", Value: id}}
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		utils.Logger.Error("Error finding user", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// NewUserService implements UserService.
func NewUserService() UserService {
	return &userService{
		collection: db.DB.Collection("users"),
	}
}
