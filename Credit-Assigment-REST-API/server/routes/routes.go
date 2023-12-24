package routes

import (
	"net/http"

	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/controllers"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/libs/logger"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/middleware"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/server/di"
	"github.com/labstack/echo/v4"
)

// Route represents the route structure for the service
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// Routes represents a route collection
type Routes []Route

// ServiceRoutes is a route collection for service handling
var ServiceRoutes Routes

func init() {
	controllersProvider := di.BuildContainer()
	err := controllersProvider.Invoke(func(
		healthController *controllers.HealthController,
		controller *controllers.ControllerData,
	) {
		ServiceRoutes = Routes{
			Route{
				Method:      http.MethodGet,
				Pattern:     "/health",
				HandlerFunc: healthController.HealthCheck,
				Name:        "HealthCheck",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/health/dependencies",
				HandlerFunc: healthController.HealthCheckDependencies,
				Name:        "HealthCheckDependencies",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/credit-assigment/:id",
				HandlerFunc: middleware.ValidatorParams(controller.GetAssigmentData),
				Name:        "GetAssigment",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/credit-assigment",
				HandlerFunc: middleware.ValidateBody(controller.CreateAssigmentData),
				Name:        "CreateAssigment",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/statistics",
				HandlerFunc: controller.GetStatisticsData,
				Name:        "Statistics",
			},
		}
	})

	if err != nil {
		logger.Error("routes", "init", err.Error())
	}
}
