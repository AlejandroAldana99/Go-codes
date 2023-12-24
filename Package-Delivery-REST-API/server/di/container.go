package di

import (
	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/controllers"
	"github.com/AlejandroAldana99/mvp_api/database"
	"github.com/AlejandroAldana99/mvp_api/repositories"
	"github.com/AlejandroAldana99/mvp_api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

// Repositories
func newOrderRepository(client *mongo.Database) *repositories.OrderRepository {
	return &repositories.OrderRepository{
		Config:  config.GetConfig(),
		MongoDB: client,
	}
}

func newUserRepository(client *mongo.Database) *repositories.UserRepository {
	return &repositories.UserRepository{
		Config:  config.GetConfig(),
		MongoDB: client,
	}
}

// Services
func newOrderService(repository *repositories.OrderRepository) *services.OrderService {
	return &services.OrderService{
		Repository: repository,
	}
}

func newUserService(repository *repositories.UserRepository) *services.UserService {
	return &services.UserService{
		Repository: repository,
	}
}

// Controllers
func newController(serviceOrder *services.OrderService, serviceUser *services.UserService) *controllers.ControllerData {
	return &controllers.ControllerData{
		ServiceOrder: serviceOrder,
		ServiceUser:  serviceUser,
	}
}

func newHealthController(client *mongo.Client, service *services.HealthService) *controllers.HealthController {
	return &controllers.HealthController{
		Configuration: config.GetConfig(),
		MongoClient:   client,
		ServiceHealth: service,
	}
}

func newHealthService() *services.HealthService {
	return &services.HealthService{}
}

// BuildContainer :
func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(database.NewMongoDBClient)
	_ = container.Provide(database.NewMongoCollection)
	_ = container.Provide(newOrderRepository)
	_ = container.Provide(newUserRepository)
	_ = container.Provide(newOrderService)
	_ = container.Provide(newUserService)
	_ = container.Provide(newController)
	_ = container.Provide(newHealthService)
	_ = container.Provide(newHealthController)

	return container
}
