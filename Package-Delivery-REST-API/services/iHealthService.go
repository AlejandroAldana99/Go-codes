package services

import "github.com/AlejandroAldana99/Package-Delivery-REST-API/models"

type IHealthService interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}
