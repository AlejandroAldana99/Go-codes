package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/models"
	"github.com/AlejandroAldana99/mvp_api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/labstack/echo/v4"
)

const millisecondsEq = 1000000.0

// HealthController :
type HealthController struct {
	Configuration config.Configuration
	MongoClient   *mongo.Client
	ServiceHealth services.IHealthService
}

// HealthCheck :
func (controller *HealthController) HealthCheck(c echo.Context) error {
	status := models.Pass.String()
	health := models.HealthStatus{
		Version:     "1.0",
		Status:      status,
		Description: "API Pod HealthCheck",
		Details:     []models.HealthComponentDetail{},
	}

	chanPodHealth := make(chan models.HealthComponentDetail)
	defer closeChannels(chanPodHealth)
	go controller.ServiceHealth.CheckPod(chanPodHealth)
	podHealth := <-chanPodHealth

	health.Details = append(health.Details, podHealth)

	return c.JSON(http.StatusOK, health)
}

// HealthCheckDependencies :
func (controller *HealthController) HealthCheckDependencies(c echo.Context) error {

	status := models.Pass.String()
	httpStatus := http.StatusOK

	mongoDetail := controller.checkMongoDBHealth()

	if mongoDetail.Status == models.Fail.String() {
		status = models.Fail.String()
		httpStatus = http.StatusServiceUnavailable
	}

	return c.JSON(httpStatus, models.HealthStatus{
		Status:      status,
		Description: "API health status",
		Version:     "1.0",
		Details: []models.HealthComponentDetail{
			mongoDetail,
		},
	})
}

func (controller *HealthController) checkMongoDBHealth() models.HealthComponentDetail {
	status := models.Pass.String()
	start := time.Now()
	ctx, _ := context.WithTimeout(context.Background(), (time.Duration(controller.Configuration.MongoConnectionTimeout) * time.Second))
	err := controller.MongoClient.Ping(ctx, readpref.PrimaryPreferred())

	if err != nil {
		logger.Error("controllers", "checkMongoDBHealth", err.Error())
		status = models.Fail.String()
	}

	finish := float32(time.Since(start).Nanoseconds()) / millisecondsEq
	return models.HealthComponentDetail{
		Name:        "MongoDB Connection",
		Type:        "datastore",
		ID:          "MONGODB-RSAUTH-CONN",
		MetricValue: finish,
		MetricUnit:  "ms",
		Time:        time.Now(),
		Status:      status,
	}
}

func closeChannels(channels ...chan models.HealthComponentDetail) {
	for _, item := range channels {
		close(item)
	}
}
