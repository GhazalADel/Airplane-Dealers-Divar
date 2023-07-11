package server

import (
	"Airplane-Divar/datastore/repair"
	"Airplane-Divar/datastore/user"
	handlers "Airplane-Divar/handlers/repair"
	"Airplane-Divar/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func repairRoutes(e *echo.Echo, db *gorm.DB) {
	repairDS := repair.NewRepairStorer(db)
	userDS := user.New(db)
	repairHandler := handlers.NewRepairHandler(repairDS, userDS)

	e.POST("/repair/ads/:adID/request", repairHandler.RequestToRepairCheck, middlewares.IsLoggedIn)
	e.GET("/repair/ads/:adID", repairHandler.GetRepairRequestByAd, middlewares.IsLoggedIn)
	e.GET("/repair/requests", repairHandler.GetAllRepairRequest, middlewares.IsLoggedIn)
	e.PUT("/repair/request/:repairRequestID", repairHandler.UpdateRepairRequest, middlewares.IsLoggedIn)
	e.GET("/repair/request/:requestID", repairHandler.GetRepairRequest, middlewares.IsLoggedIn)
	e.DELETE("/repair/ads/:adID", repairHandler.DeleteRepairRequest, middlewares.IsLoggedIn)
}
