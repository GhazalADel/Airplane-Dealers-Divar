package server

import (
	"Airplane-Divar/handlers"
	"Airplane-Divar/middlewares"

	"github.com/labstack/echo/v4"
)

func paymentRoutes(e *echo.Echo, handlers *handlers.PaymentHandler) {
	e.POST("/accounts/payment/request", handlers.PaymentRequestHandler, middlewares.IsLoggedIn)
	e.GET("/accounts/payment/verify", handlers.PaymentVerifyHandler)
}
