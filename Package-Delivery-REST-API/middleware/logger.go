package middleware

import (
	"time"

	"github.com/AlejandroAldana99/mvp_api/libs/logger"
	"github.com/labstack/echo/v4"
)

// Logger :
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (e error) {
		start := time.Now()
		e = next(context)
		if e != nil {
			return e
		}
		logger.Request(
			context.Request().Method,
			context.Response().Status,
			context.Request().RequestURI,
			start,
		)
		return
	}

}
