package server

import (
	"Airplane-Divar/handlers/ads"
	"Airplane-Divar/middlewares"

	"github.com/labstack/echo/v4"
)

func adsRoutes(e *echo.Echo, handler *ads.AdsHandler) {
	e.POST("/ads/add", handler.AddAdHandler, middlewares.IsLoggedIn)
	e.GET("/ads/:id", handler.Get, middlewares.IsLoggedIn)
	e.GET("/ads", handler.List, middlewares.IsLoggedIn)
}
