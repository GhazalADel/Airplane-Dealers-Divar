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
	e.GET("/repair/ads/:adID", repairHandler.GetRepairRequest)
	e.GET("/repair/check-requests", repairHandler.GetAllRepairRequest)
	e.PUT("/repair/check-request/:repairRequestID", repairHandler.UpdateRepairRequest)
	// e.DELETE("/expert/ads/:adID", repairHandler.DeleteExpertRequest)
}
