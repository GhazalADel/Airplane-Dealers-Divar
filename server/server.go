package server

import (
	database "Airplane-Divar/database"
	"Airplane-Divar/datastore/account"
	"Airplane-Divar/datastore/payment"
	"Airplane-Divar/datastore/user"
	"Airplane-Divar/handlers"
	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func StartServer() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Account
	accountDatastore := account.New(db)
	userDatastore := user.New(db)
	accountHandler := handlers.NewAccountHandler(userDatastore, accountDatastore)
	accountRoutes(e, accountHandler)

	// Payment
	paymentDatastore := payment.New(db)
	paymentHandler := handlers.NewPaymentHandler(paymentDatastore)
	paymentRoutes(e, paymentHandler)

	log.Fatal(e.Start(":8080"))
}
