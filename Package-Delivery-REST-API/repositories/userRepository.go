package repositories

import (
	"context"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Config  config.Configuration
	MongoDB *mongo.Database
}

func (repo UserRepository) GetUser(userID string) (models.UserData, error) {
	t := time.Now()
	var user models.UserData
	objectId, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		logger.Error("repositories", "GetUserData", oErr.Error())
		return models.UserData{}, errors.HandleServiceError(oErr)
	}
	err := repo.MongoDB.Collection("users").FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
	).Decode(&user)
	if err != nil {
		logger.Error("repositories", "GetUserData", err.Error())
		return models.UserData{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetUser", t)

	return user, nil
}

func (repo UserRepository) CreateUser(data models.UserData) error {

	t := time.Now()
	_, err := repo.MongoDB.Collection("users").InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("repositories", "CraeteUser", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CreateUser", t)

	return nil
}

func (repo UserRepository) GetUserByEmail(email string) (models.UserData, error) {
	t := time.Now()
	var user models.UserData
	err := repo.MongoDB.Collection("users").FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).
		Decode(&user)

	if err != nil {
		logger.Error("repositories", "GetUserData", err.Error())
		return models.UserData{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetUser", t)

	return user, nil
}
