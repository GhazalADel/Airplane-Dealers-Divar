package server

import (
	database "Airplane-Divar/database"
	adsDatastore "Airplane-Divar/datastore/ads"
	"Airplane-Divar/datastore/user"
	"Airplane-Divar/handlers"
	adsHandler "Airplane-Divar/handlers/ads"
	userHandler "Airplane-Divar/handlers/user"
	"log"

	bookmarkDatastore "Airplane-Divar/datastore/bookmarks"
	bookmarksHanlder "Airplane-Divar/handlers/bookmarks"

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
	// Ads
	datastore := adsDatastore.New(db)
	adsHandler := adsHandler.New(datastore)
	adsRoutes(e, adsHandler)

	// User
	userDatastore := user.New(db)
	userHandler := userHandler.NewUserHandler(userDatastore)
	userRoutes(e, userHandler)

	// Payment
	paymentRoutes(e, db)

	// Bookmarks
	bmDatastore := bookmarkDatastore.New(db)
	bmHandlers := bookmarksHanlder.New(bmDatastore)
	bookmarksRoutes(e, bmHandlers)

	log.Fatal(e.Start(":8080"))
}
