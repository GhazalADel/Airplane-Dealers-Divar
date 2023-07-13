package bookmarks

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	logging_service "Airplane-Divar/service/logging"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookmarkssHandler struct {
	datastore datastore.Bookmark
}

func New(bookmarks datastore.Bookmark) *BookmarkssHandler {
	return &BookmarkssHandler{datastore: bookmarks}
}

type ErrorAddAd struct {
	ResponseCode int    `json:"responsecode"`
	Message      string `json:"message"`
}

// Get list of all bookmarks.
// @Summary bookmarks list
// @Description Retrieves all bookmarks of this user
// @Tags bookmarks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Success 200 {object} []models.AdResponse
// @Failure 403 {object} ErrorAddAd
// @Failure 500 {object} ErrorAddAd
// @Router /bookmarks/list [get]
func (b BookmarkssHandler) ListBookmarks(c echo.Context) error {
	user := c.Get("user").(models.User)
	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(http.StatusForbidden, models.Response{ResponseCode: 403, Message: "Airlines Can see bookmarks!"})
	}
	ads, err := b.datastore.GetAdsByUserID(int(user.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: err.Error()})
	}
	if len(ads) == 0 {
		return c.JSON(http.StatusOK, models.Response{ResponseCode: 200, Message: "You don't have any bookmark"})
	}
	return c.JSON(http.StatusOK, ads)
}

// add a new bookmark.
// @Summary add bookmark
// @Description add bookmark using given ad id
// @Tags bookmarks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Param id path int true "Ad ID"
// @Success 200 {object} models.BookmarksResponse
// @Failure 403 {object} ErrorAddAd
// @Failure 422 {object} ErrorAddAd
// @Failure 500 {object} ErrorAddAd
// @Router /bookmarks/add/{id} [put]
func (b BookmarkssHandler) AddBookmark(c echo.Context) error {
	user := c.Get("user").(models.User)
	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(http.StatusForbidden, models.Response{ResponseCode: 403, Message: "Airlines Can add bookmark!"})
	}
	ad_id := c.Param("id")
	ad_id_int, err := strconv.Atoi(ad_id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Id should be integer!"})
	}
	bookmark, err := b.datastore.AddBookmark(int(user.ID), ad_id_int)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: err.Error()})
	}

	// ------ Report Log ------
	logsrv := logging_service.GetInstance()
	err = logsrv.ReportActivity(user.Role, user.ID, "Ads", uint(ad_id_int), consts.LOG_BOOKMARK, "")
	if err != nil {
		_ = fmt.Errorf("cannot log activity %v", consts.LOG_BOOKMARK)
	}
	// ------ Report Log ------

	return c.JSON(http.StatusOK, bookmark)
}

// delete bookmark
// @Summary delete existing bookmark
// @Description delete existing bookmark using id
// @Tags bookmarks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Param id path int true "Ad ID"
// @Success 200 {string} string "Bookmark Deleted Successfully"
// @Failure 403 {object} ErrorAddAd
// @Failure 422 {object} ErrorAddAd
// @Failure 500 {object} ErrorAddAd
// @Router /bookmarks/delete/{id} [delete]
func (b BookmarkssHandler) DeleteBookmark(c echo.Context) error {
	user := c.Get("user").(models.User)
	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(http.StatusForbidden, models.Response{ResponseCode: 403, Message: "Airlines Can delete bookmark!"})
	}
	ad_id := c.Param("id")
	ad_id_int, err := strconv.Atoi(ad_id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Id should be integer!"})
	}
	err_delete := b.datastore.DeleteBookmark(int(user.ID), ad_id_int)
	if err_delete != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: err_delete.Error()})
	}

	// ------ Report Log ------
	logsrv := logging_service.GetInstance()
	err = logsrv.ReportActivity(user.Role, user.ID, "Ads", uint(ad_id_int), consts.LOG_BOOKMARK_REMOVE, "")
	if err != nil {
		_ = fmt.Errorf("cannot log activity %v", consts.LOG_BOOKMARK_REMOVE)
	}
	// ------ Report Log ------

	return c.JSON(http.StatusOK, "Bookmark Deleted Successfully")
}
