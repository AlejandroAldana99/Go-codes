package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4/middleware"
)

// Configuration :
type Configuration struct {
	Port                   string
	CORSConfig             middleware.CORSConfig
	MongoHost              string
	MongoPort              string
	MongoUser              string
	MongoPassword          string
	DatabaseName           string
	MongoConnectionTimeout int
	MongoMaxConnections    int64
}

var config Configuration

func init() {
	configureEnvironment()
	configureCors()
}

// GetConfig :
func GetConfig() Configuration {
	return config
}

func configureEnvironment() {
	if os.Getenv("GO_ENV") == "" {
		godotenv.Load()
	}
	config.Port = os.Getenv("PORT")
	config.MongoHost = os.Getenv("MONGODB_HOSTS")
	config.MongoPort = os.Getenv("MONGODB_PORT")
	config.MongoUser = os.Getenv("MONGODB_USER")
	config.MongoPassword = os.Getenv("MONGODB_PASSWORD")
	config.DatabaseName = os.Getenv("MONGODB_DATABASE")
	config.MongoConnectionTimeout, _ = strconv.Atoi(os.Getenv("MONGODB_CONNECTION_TIMEOUT"))
	config.MongoMaxConnections, _ = strconv.ParseInt(os.Getenv("MONGODB_MAX_CONNECTIONS"), 10, 64)

}

func configureCors() {

	corsAllowedHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")
	corsAllowedMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
	corsAllowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	config.CORSConfig = middleware.CORSConfig{
		AllowHeaders: corsAllowedHeaders,
		AllowMethods: corsAllowedMethods,
		AllowOrigins: corsAllowedOrigins,
		MaxAge:       7200,
	}
}
