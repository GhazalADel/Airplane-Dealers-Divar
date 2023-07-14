package server

import (
	"Airplane-Divar/datastore/expert"
	"Airplane-Divar/datastore/payment"
	"Airplane-Divar/datastore/repair"
	"Airplane-Divar/datastore/user"
	handlers "Airplane-Divar/handlers/payment"
	"Airplane-Divar/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func paymentRoutes(e *echo.Echo, db *gorm.DB) {
	paymentDS := payment.New(db)
	repairDS := repair.NewRepairStorer(db)
	expertDS := expert.NewExpertStorer(db)
	userDS := user.New(db)
	paymentHandler := handlers.NewPaymentHandler(userDS, expertDS, repairDS, paymentDS)

	e.POST("/users/payment/request", paymentHandler.PaymentRequestHandler, middlewares.IsLoggedIn)
	e.GET("/users/payment/verify", paymentHandler.PaymentVerifyHandler)
}
