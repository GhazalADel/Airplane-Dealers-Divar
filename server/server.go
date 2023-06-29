package server

import (
	"log"

	database "Airplane-Divar/database"
	"Airplane-Divar/handlers"

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
	adsHandler := handlers.NewAdsHandler(db)
	adsRoutes(e, adsHandler)

	log.Fatal(e.Start(":8080"))
}
