package services

import (
	"net/http"
	"time"

	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/models"
)

type HealthService struct{}

func updateHealthDetails(details *models.HealthComponentDetail, now time.Time, statusCode int) {
	details.Status = "fail"
	if statusCode == http.StatusOK {
		details.Status = "pass"
	}
	details.MetricValue = float32(time.Since(now).Nanoseconds()) / millisecondsEq
	details.MetricUnit = "ms"
}

// CheckPod :
func (service *HealthService) CheckPod(chanHealth chan models.HealthComponentDetail) {
	now := time.Now()
	details := models.HealthComponentDetail{
		Name: "instance",
		Type: "pod",
		Time: now,
	}

	statusCode := 200
	updateHealthDetails(&details, now, statusCode)
	chanHealth <- details
}
