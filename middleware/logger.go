package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\x1b[34m[${time_custom}]\x1b[0m " + // Blue formatted time
			"\x1b[32m${method}\x1b[0m " + // Green HTTP method
			"\x1b[36m${uri}\x1b[0m " + // Cyan URI
			"\x1b[33mStatus:\x1b[0m \x1b[31m${status}\x1b[0m\n", // Yellow label + Red status
		CustomTimeFormat: "02-01-2006 15:04",
	})
}
