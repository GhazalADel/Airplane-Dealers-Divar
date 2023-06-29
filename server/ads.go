package server

import (
	"Airplane-Divar/handlers"

	"github.com/labstack/echo/v4"
)

func adsRoutes(e *echo.Echo, handler *handlers.AdsHandler) {
	e.POST("/ads/add", handler.AddAdHandler)
}
