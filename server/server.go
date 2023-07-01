package server

import (
	database "Airplane-Divar/database"
	"Airplane-Divar/datastore/account"
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
	account_datastore := account.New(db)
	user_datastore := user.New(db)
	accountHandler := handlers.NewAccountHandler(user_datastore, account_datastore)
	accountRoutes(e, accountHandler)

	// Payment
	paymentRoutes(e)

	log.Fatal(e.Start(":8080"))
}
