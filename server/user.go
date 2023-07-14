package server

import (
	userHandler "Airplane-Divar/handlers/user"

	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo, handler *userHandler.UserHandler) {
	e.POST("/users/login", handler.LoginHandler)
	e.POST("/users/register", handler.RegisterHandler)
}
