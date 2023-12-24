package repositories

import (
	"context"
	"time"

	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/config"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/constants"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/errors"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/libs/logger"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssigmentRepository struct {
	Config  config.Configuration
	MongoDB *mongo.Database
}

func (repo AssigmentRepository) GetAssigment(AssigmentID string) (models.AssigmentData, error) {
	t := time.Now()
	var assigment models.AssigmentData
	objectId, oErr := primitive.ObjectIDFromHex(AssigmentID)
	if oErr != nil {
		logger.Error("repositories", "GetAssigment", oErr.Error())
		return models.AssigmentData{}, errors.HandleServiceError(oErr)
	}
	err := repo.MongoDB.Collection(constants.CollectionName).FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
	).Decode(&assigment)

	if err != nil {
		logger.Error("repositories", "GetAssigmentData", err.Error())
		return models.AssigmentData{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetAssigment", t)
	return assigment, nil
}

func (repo AssigmentRepository) CreateAssigment(data models.AssigmentData) error {
	t := time.Now()
	_, err := repo.MongoDB.Collection(constants.CollectionName).InsertOne(context.TODO(), data)
	if err != nil {
		logger.Error("repositories", "CreateAssigment", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CreateAssigment", t)

	return nil
}

func (repo AssigmentRepository) GetStatistics() (models.StatisticsData, error) {
	t := time.Now()
	statitics := models.StatisticsData{}

	totalDocs, err := repo.MongoDB.Collection(constants.CollectionName).CountDocuments(
		context.TODO(), bson.D{})

	if err != nil {
		logger.Error("repositories", "GetStatistics", err.Error())
		return statitics, errors.HandleServiceError(err)
	}

	totalSuccess, err := repo.MongoDB.Collection(constants.CollectionName).CountDocuments(
		context.TODO(),
		bson.D{{Key: "noncombinated", Value: false}},
	)

	if err != nil {
		logger.Error("repositories", "GetStatistics", err.Error())
		return statitics, errors.HandleServiceError(err)
	}

	statitics.TotalAssignations = float32(totalDocs)
	statitics.SuccessfulAssignations = float32(totalSuccess)

	logger.Performance("repository", "GetStatistics", t)
	return statitics, nil
}
