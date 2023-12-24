package services

import "github.com/AlejandroAldana99/Package-Delivery-REST-API/models"

type IOrderService interface {
	GetOrder(orderID string, userID string, role string) (models.OrderData, error)
	CreateOrder(data models.OrderData) (models.ResponseData, error)
	UpdateOrderStatus(orderID string, status string, role string, userID string) error
}
