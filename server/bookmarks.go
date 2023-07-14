package server

import (
	"Airplane-Divar/handlers/bookmarks"
	"Airplane-Divar/middlewares"

	"github.com/labstack/echo/v4"
)

func bookmarksRoutes(e *echo.Echo, handler *bookmarks.BookmarkssHandler) {
	e.DELETE("/bookmarks/delete/:id", handler.DeleteBookmark, middlewares.IsLoggedIn)
	e.GET("/bookmarks/list", handler.ListBookmarks, middlewares.IsLoggedIn)
	e.PUT("/bookmarks/add/:id", handler.AddBookmark, middlewares.IsLoggedIn)
}
