package services

import (
	e "errors"
	"time"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/AlejandroAldana99/mvp_api/errors"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/repositories"
)

const millisecondsEq = 1000000.0

type OrderService struct {
	Repository repositories.IOrderRepository
}

func (service OrderService) GetOrder(orderID string, userID string, role string) (models.OrderData, error) {
	order, err := service.Repository.GetOrder(orderID)
	if err != nil {
		logger.Error("services", "GetOrder", err.Error())
		return order, errors.HandleServiceError(err)
	}

	if order.OwnerID != userID && role != constants.AdminRole {
		uErr := e.New("invalid order")
		logger.Error("services", "GetOrder", uErr.Error())
		return models.OrderData{}, errors.HandleServiceError(uErr)
	}

	return order, nil
}

func (service OrderService) CreateOrder(data models.OrderData) (models.ResponseData, error) {
	// Add time
	data.TimeRegistry = time.Now()
	response := models.ResponseData{
		Notification: "",
		Status:       "Faild",
	}
	// Valid status
	if !constants.StatusList[data.Status] {
		err := e.New("Invalid Status")
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return response, errors.HandleServiceError(err)
	}
	// Valid coordinates
	if !validCoordinates(data.Coordinates.Latitude, data.Coordinates.Longitude) {
		err := e.New("Invalid Coordinates")
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return response, errors.HandleServiceError(err)
	}
	// Read packages and bussines rules
	for i := range data.Packages {
		sp := false
		data.Packages[i], sp = completeSize(data.Packages[i])
		if sp {
			data.Description = "Special Package: Special deal is required."
			response.Notification = "Special Package: Special deal is required."
		}
	}

	err := service.Repository.CreateOrder(data)
	if err != nil {
		logger.Error("services", "CreateOrder", err.Error())
		return response, errors.HandleServiceError(err)
	}

	response.Status = "Success"
	return response, nil
}

func (service OrderService) UpdateOrderStatus(orderID string, status string, role string, userID string) error {
	if !constants.StatusList[status] {
		err := e.New("invalid status")
		logger.Error("services", "UpdateOrderStatus", err.Error())
		return errors.HandleServiceError(err)
	}

	// Cancelation statement
	if status == constants.CancelStatus {
		data, dErr := service.Repository.GetOrder(orderID)
		if dErr != nil {
			logger.Error("services", "CancelOrder", dErr.Error())
			return errors.HandleServiceError(dErr)
		}

		if data.OwnerID != userID {
			uErr := e.New("invalid order")
			logger.Error("services", "CancelOrder", uErr.Error())
			return errors.HandleServiceError(uErr)
		}

		if compareStatus(data.Status) {
			err := e.New("non-cancellable order")
			logger.Error("services", "CancelOrder", err.Error())
			return errors.HandleServiceError(err)
		}

		refund := false
		if compareTime(data.TimeRegistry, time.Now()) {
			refund = true
		}

		err := service.Repository.UpdateOrderStatus(orderID, status, refund)
		if err != nil {
			logger.Error("services", "CancelOrder", err.Error())
			return errors.HandleServiceError(err)
		}
		// Update statement
	} else {
		if role != constants.AdminRole {
			err := e.New("invalid role")
			logger.Error("services", "UpdateOrderStatus", err.Error())
			return errors.HandleServiceError(err)
		}
		err := service.Repository.UpdateOrderStatus(orderID, status, false)
		if err != nil {
			logger.Error("services", "UpdateOrderStatus", err.Error())
			return errors.HandleServiceError(err)
		}
	}

	return nil
}
