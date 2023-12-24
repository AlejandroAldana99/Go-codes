package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlejandroAldana99/mvp_api/config"
	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/AlejandroAldana99/mvp_api/middleware"
	"github.com/AlejandroAldana99/mvp_api/server/routes"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

var server *echo.Echo

func init() {
	server = echo.New()
	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(server)
}

func setupRoutes() {
	for _, r := range routes.ServiceRoutes {
		server.Add(
			r.Method,
			r.Pattern,
			r.HandlerFunc,
		).Name = r.Name
	}
}

func setupMiddleware() {
	server.Use(
		em.CORSWithConfig(
			config.GetConfig().CORSConfig,
		),
		middleware.Logger,
	)
}

func setupErrorHandler() {
	server.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		logger.Request(
			c.Request().Method,
			code,
			c.Request().RequestURI,
			time.Now(),
		)
		// Call the default handler to return the HTTP response
		server.DefaultHTTPErrorHandler(err, c)
	}
}

func startServer() {
	server.
		Logger.
		Fatal(
			server.Start(
				fmt.Sprintf(":%s", config.GetConfig().Port)),
		)
}

// InitServer :
func InitServer() {
	server.HideBanner = true
	setupRoutes()
	setupErrorHandler()
	setupMiddleware()
	startServer()
}
