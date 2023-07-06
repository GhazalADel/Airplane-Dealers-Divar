package server

import (
	"Airplane-Divar/datastore/expert"
	"Airplane-Divar/datastore/user"
	handlers "Airplane-Divar/handlers/expert"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func expertRoutes(e *echo.Echo, db *gorm.DB) {
	expertDS := expert.NewExpertStorer(db)
	userDS := user.NewUserStorer(db)
	expertHandler := handlers.NewExpertHandler(expertDS, userDS)

	e.POST("/expert/ads/:adID/check-request", expertHandler.RequestToExpertCheck)
	e.GET("/expert/check-requests", expertHandler.GetAllExpertRequest)
	e.GET("/expert/ads/:adID", expertHandler.GetExpertRequest)
	e.PUT("/expert/check-request/:expertRequestID", expertHandler.UpdateCheckExpert)
	e.DELETE("/expert/ads/:adID", expertHandler.DeleteExpertRequest)
}
