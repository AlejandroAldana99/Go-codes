package services

import "github.com/AlejandroAldana99/mvp_api/models"

type IHealthService interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}
