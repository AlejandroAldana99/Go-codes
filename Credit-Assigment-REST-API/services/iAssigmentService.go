package services

import "github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"

type IAssigmentService interface {
	GetAssigment(AssigmentID string) (models.AssigmentData, error)
	CreateAssigment(data models.AssigmentData) (models.ResponseData, error)
	GetStatistics() (models.StatisticsData, error)
}
