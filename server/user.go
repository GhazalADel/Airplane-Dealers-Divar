package server

import (
	"Airplane-Divar/handlers"

	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo, handler *handlers.UserHandler) {
	e.POST("/users/login", handler.LoginHandler)
	e.POST("/users/register", handler.RegisterHandler)
}
