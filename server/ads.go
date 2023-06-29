package server

import (
	"Airplane-Divar/handlers/ads"

	"github.com/labstack/echo/v4"
)

func adsRoutes(e *echo.Echo, handler *ads.AdsHandler) {
	e.POST("/ads/add", handler.AddAdHandler)
	e.GET("/ads/:id", handler.Get)
	e.GET("/ads", handler.GetAll)
}
