package server

import (
	database "Airplane-Divar/database"
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

	// Expert
	expertRoutes(e, db)

	// Repair
	repairRoutes(e, db)
	// User
	userDatastore := user.New(db)
	userHandler := handlers.NewUserHandler(userDatastore)
	userRoutes(e, userHandler)

	// Payment
	paymentDatastore := payment.New(db)
	paymentHandler := handlers.NewPaymentHandler(paymentDatastore)
	paymentRoutes(e, paymentHandler)

	log.Fatal(e.Start(":8080"))
}
