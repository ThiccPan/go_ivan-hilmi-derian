package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Logger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	  }))
}