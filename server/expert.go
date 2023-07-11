package server

import (
	"Airplane-Divar/datastore/expert"
	"Airplane-Divar/datastore/user"
	handlers "Airplane-Divar/handlers/expert"
	"Airplane-Divar/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func expertRoutes(e *echo.Echo, db *gorm.DB) {
	expertDS := expert.NewExpertStorer(db)
	userDS := user.New(db)
	expertHandler := handlers.NewExpertHandler(expertDS, userDS)

	e.POST("/expert/ads/:adID/check-request", expertHandler.RequestToExpertCheck, middlewares.IsLoggedIn)
	e.GET("/expert/check-requests", expertHandler.GetAllExpertRequest, middlewares.IsLoggedIn)
	e.GET("/expert/ads/:adID", expertHandler.GetExpertRequestByAd, middlewares.IsLoggedIn)
	e.GET("/expert/check-request/:requestID", expertHandler.GetExpertRequest, middlewares.IsLoggedIn)
	e.PUT("/expert/check-request/:expertRequestID", expertHandler.UpdateCheckExpert, middlewares.IsLoggedIn)
	e.DELETE("/expert/ads/:adID", expertHandler.DeleteExpertRequest, middlewares.IsLoggedIn)
}
