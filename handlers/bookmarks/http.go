package bookmarks

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"

	"github.com/labstack/echo/v4"
)

type BookmarkssHandler struct {
	datastore datastore.Bookmark
}

func New(bookmarks datastore.Bookmark) *BookmarkssHandler {
	return &BookmarkssHandler{datastore: bookmarks}
}

func (b BookmarkssHandler) ListBookmarks(c echo.Context) error {
	user := c.Get("user")
	user = user.(models.User)
	id := uint(user.(models.User).ID)
	ad.UserID = id
}

func (b BookmarkssHandler) AddBookmark(c echo.Context) error {
}
