package repositories

import "github.com/AlejandroAldana99/Package-Delivery-REST-API/models"

type IOrderRepository interface {
	GetOrder(userID string) (models.OrderData, error)
	CreateOrder(data models.OrderData) error
	UpdateOrderStatus(orderID string, status string, refund bool) error
}
