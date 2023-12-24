package repositories

import "github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"

type IAssigmentRepository interface {
	GetAssigment(assigmentID string) (models.AssigmentData, error)
	CreateAssigment(data models.AssigmentData) error
	GetStatistics() (models.StatisticsData, error)
}
