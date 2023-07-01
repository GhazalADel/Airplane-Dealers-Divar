package server

import (
	"Airplane-Divar/handlers"

	"github.com/labstack/echo/v4"
)

func accountRoutes(e *echo.Echo, handler *handlers.AccountHandler) {
	e.POST("/accounts/login", handler.LoginHandler)
	e.POST("/accounts/register", handler.RegisterHandler)
}
