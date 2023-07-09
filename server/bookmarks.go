package server

import (
	"Airplane-Divar/handlers/bookmarks"

	"github.com/labstack/echo/v4"
)

func bookmarksRoutes(e *echo.Echo, handler *bookmarks.BookmarkssHandler) {
	e.GET("/bookmarks/list", handler.ListBookmarks)
	e.PUT("/bookmarks/add", handler.AddBookmark)
}
