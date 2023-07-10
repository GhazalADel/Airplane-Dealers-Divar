package server

import (
	database "Airplane-Divar/database"
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

	log.Fatal(e.Start(":8080"))
}
