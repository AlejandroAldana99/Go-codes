package di

import (
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/config"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/controllers"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/database"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/repositories"
	"github.com/AlejandroAldana99/Credit-Assigment-REST-API/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

// Repositories
func newAssigmentRepository(client *mongo.Database) *repositories.AssigmentRepository {
	return &repositories.AssigmentRepository{
		Config:  config.GetConfig(),
		MongoDB: client,
	}
}

// Services
func newAssigmentService(repository *repositories.AssigmentRepository) *services.AssigmentService {
	return &services.AssigmentService{
		Repository: repository,
	}
}

// Controllers
func newController(service *services.AssigmentService) *controllers.ControllerData {
	return &controllers.ControllerData{
		Service: service,
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
	_ = container.Provide(newAssigmentRepository)
	_ = container.Provide(newAssigmentService)
	_ = container.Provide(newController)
	_ = container.Provide(newHealthService)
	_ = container.Provide(newHealthController)

	return container
}
