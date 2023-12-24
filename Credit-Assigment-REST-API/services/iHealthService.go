package services

import "github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"

type IHealthService interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}
