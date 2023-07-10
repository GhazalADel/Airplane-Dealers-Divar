package server

import (
	"log"

	database "Airplane-Divar/database"
	adsDatastore "Airplane-Divar/datastore/ads"
	adsHandler "Airplane-Divar/handlers/ads"

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

	// Ads
	datastore := adsDatastore.New(db)
	adsHandler := adsHandler.New(datastore)
	adsRoutes(e, adsHandler)

	log.Fatal(e.Start(":8080"))
}
