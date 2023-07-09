package server

import (
	"Airplane-Divar/datastore/repair"
	"Airplane-Divar/datastore/user"
	handlers "Airplane-Divar/handlers/repair"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func repairRoutes(e *echo.Echo, db *gorm.DB) {
	repairDS := repair.NewRepairStorer(db)
	userDS := user.NewUserStorer(db)
	repairHandler := handlers.NewRepairHandler(repairDS, userDS)

	e.POST("/repair/ads/:adID/check-request", repairHandler.RequestToRepairCheck)
	// e.GET("/expert/check-requests", repairHandler.GetAllExpertRequest)
	// e.GET("/expert/ads/:adID", repairHandler.GetExpertRequest)
	// e.PUT("/expert/check-request/:expertRequestID", repairHandler.UpdateCheckExpert)
	// e.DELETE("/expert/ads/:adID", repairHandler.DeleteExpertRequest)
}
