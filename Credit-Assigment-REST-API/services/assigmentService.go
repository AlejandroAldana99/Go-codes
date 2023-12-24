package services

import (
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/errors"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/libs/logger"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/repositories"
)

const millisecondsEq = 1000000.0

type AssigmentService struct {
	Repository repositories.IAssigmentRepository
}

func (service AssigmentService) GetAssigment(AssigmentID string) (models.AssigmentData, error) {
	Assigment, err := service.Repository.GetAssigment(AssigmentID)
	if err != nil {
		logger.Error("services", "GetAssigment", err.Error())
		return Assigment, errors.HandleServiceError(err)
	}

	return Assigment, nil
}

func (service AssigmentService) CreateAssigment(data models.AssigmentData) (models.ResponseData, error) {
	response := models.ResponseData{}

	data.Combinations = calculateLoans(data.OriginalInvestment)
	if data.Combinations[0].CreditType300 <= 0 && data.Combinations[0].CreditType500 <= 0 && data.Combinations[0].CreditType700 <= 0 {
		response.Message = "Amount impossible to invest"
		data.NonCombinated = true
	}

	err := service.Repository.CreateAssigment(data)
	if err != nil {
		logger.Error("services", "CreateAssigment", err.Error())
		return response, errors.HandleServiceError(err)
	}

	response.CreditResponse = data.Combinations
	return response, nil
}

func (service AssigmentService) GetStatistics() (models.StatisticsData, error) {
	statistics, err := service.Repository.GetStatistics()
	if err != nil {
		logger.Error("services", "GetStatistics", err.Error())
		return statistics, errors.HandleServiceError(err)
	}

	statistics.FailedAssignations = (statistics.TotalAssignations - statistics.SuccessfulAssignations)

	avgSuccess := (statistics.SuccessfulAssignations / statistics.TotalAssignations)
	avgFailed := (statistics.FailedAssignations / statistics.TotalAssignations)

	statistics.AvgSuccessfulAssignations = float32(avgSuccess)
	statistics.AvgFailedAssignations = float32(avgFailed)

	return statistics, nil
}
