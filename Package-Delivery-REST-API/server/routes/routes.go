package routes

import (
	"net/http"

	"github.com/AlejandroAldana99/mvp_api/controllers"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/middleware"
	"github.com/AlejandroAldana99/mvp_api/server/di"
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
				Pattern:     "/mvp/order/:id",
				HandlerFunc: middleware.ParamsValidatorID(middleware.ValidateToken(controller.GetOrderData)),
				Name:        "GetOrder",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/mvp/order",
				HandlerFunc: middleware.ValidateOrderBody(controller.CreateOrderData),
				Name:        "CreateOrder",
			},
			Route{
				Method:      http.MethodPatch,
				Pattern:     "/mvp/order",
				HandlerFunc: middleware.ParamsOrderValidator(middleware.ValidateToken(controller.UpdateOrderStatus)),
				Name:        "UpdateStatus",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/mvp/user/:id",
				HandlerFunc: middleware.ValidateToken(middleware.ParamsValidatorID(controller.GetUserData)),
				Name:        "GetUser",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/mvp/user",
				HandlerFunc: middleware.ValidateUserBody(controller.CreateUserData),
				Name:        "CreateUser",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "/mvp/user/login",
				HandlerFunc: middleware.ParamsLoginValidator(controller.Login),
				Name:        "Login",
			},
		}
	})

	if err != nil {
		logger.Error("routes", "init", err.Error())
	}
}
