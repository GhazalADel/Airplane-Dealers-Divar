package bookmarks

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkHandler_ListBookmarks(t *testing.T) {
	e := echo.New()

	t.Run("non-airline user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bookmarks/list", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[2])

		a := New(mockDatastore{})
		err := a.ListBookmarks(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Airlines Can see bookmarks!", response.Message)
	})

	t.Run("zero bookmark", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bookmarks/list", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[3])

		a := New(mockDatastore{})
		err := a.ListBookmarks(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "You don't have any bookmark", response.Message)
	})

	t.Run("valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bookmarks/list", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err := a.ListBookmarks(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response []models.AdResponse
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, 1, len(response))
	})

}
func TestBookmarkHandler_AddBookmark(t *testing.T) {
	e := echo.New()
	t.Run("non-airline user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/bookmarks/add/1", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[2])

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Airlines Can add bookmark!", response.Message)
	})

	t.Run("non-integer id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/bookmarks/add/salam", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("salam")

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Id should be integer!", response.Message)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/bookmarks/add/123", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("123")

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
	})

	t.Run("inactive ad", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/bookmarks/add/3", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("3")

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/bookmarks/add/2", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("2")

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
	})
}

func TestBookmarkHandler_DeleteBookmark(t *testing.T) {
	e := echo.New()

	t.Run("non-airline user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/bookmarks/delete/1", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[2])
		c.SetParamNames("id")
		c.SetParamValues("1")

		a := New(mockDatastore{})
		err := a.DeleteBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Airlines Can delete bookmark!", response.Message)
	})

	t.Run("non-integer id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/bookmarks/delete/salam", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("salam")

		a := New(mockDatastore{})
		err := a.DeleteBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Id should be integer!", response.Message)
	})

	t.Run("invalid id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/bookmarks/delete/100", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("100")

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/bookmarks/delete/1", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])
		c.SetParamNames("id")
		c.SetParamValues("1")

		a := New(mockDatastore{})
		err := a.AddBookmark(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
	})

}

type mockDatastore struct {
	bookmarks_data []models.BookmarksResponse
}

var (
	mockCategoryData = []models.Category{
		{
			ID:   1,
			Name: "small-passenger",
		},
		{
			ID:   2,
			Name: "big-passenger",
		},
	}

	mockAdData = []models.AdResponse{
		{
			ID:            1,
			UserID:        1,
			Image:         "example1.jpg",
			Description:   "This is example ad 1.",
			Subject:       "Example Ad 1",
			Price:         1000,
			CategoryID:    2,
			Status:        string(consts.ACTIVE),
			FlyTime:       1000,
			AirplaneModel: "XYZ123",
			RepairCheck:   true,
			ExpertCheck:   false,
			PlaneAge:      7,
		},
		{
			ID:            2,
			UserID:        1,
			Image:         "example2.jpg",
			Description:   "This is example ad 2.",
			Subject:       "Example Ad 2",
			Price:         2000,
			CategoryID:    1,
			Status:        string(consts.ACTIVE),
			FlyTime:       1000,
			AirplaneModel: "ABC456",
			RepairCheck:   true,
			ExpertCheck:   true,
			PlaneAge:      3,
		},
		{
			ID:            3,
			UserID:        1,
			Image:         "example3.jpg",
			Description:   "This is example ad 3.",
			Subject:       "Example Ad 3",
			Price:         3000,
			CategoryID:    1,
			Status:        string(consts.INACTIVE),
			FlyTime:       1000,
			AirplaneModel: "ABC456",
			RepairCheck:   true,
			ExpertCheck:   true,
			PlaneAge:      3,
		},
	}
	mockUserData = []models.User{
		{
			ID:       1,
			Username: "John",
			Password: "John123",
			Role:     consts.ROLE_AIRLINE,
		},
		{
			ID:       2,
			Username: "Rose",
			Password: "Rose123",
			Role:     consts.ROLE_AIRLINE,
		},
		{
			ID:       3,
			Username: "Sophia",
			Password: "Sophia123",
			Role:     consts.ROLE_ADMIN,
		},
		{
			ID:       4,
			Username: "Charlie",
			Password: "Charlie123",
			Role:     consts.ROLE_AIRLINE,
		},
	}
	mockBookmarkData = []models.BookmarksResponse{
		{
			UserID: 1,
			AdsID:  1,
		},
		{
			UserID: 2,
			AdsID:  2,
		},
		{
			UserID: 2,
			AdsID:  1,
		},
	}
)

func (m mockDatastore) GetAdsByUserID(id int) ([]models.AdResponse, error) {
	if id == 1 {
		return []models.AdResponse{mockAdData[0]}, nil
	} else if id == 2 {
		return mockAdData[0:2], nil
	} else if id == 4 {
		return []models.AdResponse{}, nil
	}
	return mockAdData, nil
}

func (m mockDatastore) AddBookmark(userID, adID int) (models.BookmarksResponse, error) {
	if adID < 0 || adID > 2 {
		return models.BookmarksResponse{}, errors.New("")
	}
	found := false
	for _, v := range m.bookmarks_data {
		if v.AdsID == uint(adID) && v.UserID == uint(userID) {
			found = true
			break
		}
	}
	if found {
		return models.BookmarksResponse{}, errors.New("")
	}
	isActive := true
	for _, v := range mockAdData {
		if v.ID == uint(adID) {
			if v.Status != string(consts.ACTIVE) {
				isActive = false
				break
			}
		}
	}
	if !isActive {
		return models.BookmarksResponse{}, errors.New("")
	}
	b := models.BookmarksResponse{
		UserID: uint(userID),
		AdsID:  uint(adID),
	}
	m.bookmarks_data = append(m.bookmarks_data, b)
	return b, nil
}

func (m mockDatastore) DeleteBookmark(userID, adID int) error {
	ex_size := len(m.bookmarks_data)
	ind := -1
	for i, v := range m.bookmarks_data {
		if v.AdsID == uint(adID) && v.UserID == uint(userID) {
			ind = i
			break
		}
	}
	m.bookmarks_data = append(m.bookmarks_data[:ind], m.bookmarks_data[ind+1:]...)
	if len(m.bookmarks_data)+1 == ex_size {
		return nil
	}
	return errors.New("")
}
