package ads

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAdsHandler_Get(t *testing.T) {
	testcases := []struct {
		id            string
		response      []models.Ad
		expectedCode  int
		expectedError string
	}{
		{
			id:            "2",
			expectedCode:  http.StatusInternalServerError,
			expectedError: "could not retrieve ads",
		},
		{
			id:            "1a",
			expectedCode:  http.StatusBadRequest,
			expectedError: "invalid parameter id",
		},
		{
			id:           "0",
			response:     mockAdData,
			expectedCode: http.StatusOK,
		},
		{
			id:           "1",
			response:     mockAdData[:1],
			expectedCode: http.StatusOK,
		},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/ads/"+v.id, nil)
		w := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, w)
		c.Set("user", mockUserData[0])

		c.SetParamNames("id")
		c.SetParamValues(v.id)

		a := New(mockDatastore{})
		a.Get(c)

		if v.expectedCode == http.StatusOK {
			var adRes []models.Ad
			err := json.Unmarshal(w.Body.Bytes(), &adRes)
			assert.NoError(t, err)

			if !reflect.DeepEqual(adRes, v.response) {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, adRes, v.response)
			}
		} else {
			var errorRes string
			err := json.Unmarshal(w.Body.Bytes(), &errorRes)
			assert.NoError(t, err)

			if !reflect.DeepEqual(errorRes, v.expectedError) {
				t.Errorf("[http Get() TEST%d]Failed. Got %v\tExpected %v\n", i+1, errorRes, v.expectedError)
			}
		}
	}
}

func TestAdsHandler_ListFilter(t *testing.T) {
	testcases := []struct {
		query         string
		response      []models.Ad
		expectedCode  int
		expectedError string
	}{
		{
			query:        "plane_age=7",
			response:     mockAdData[:1],
			expectedCode: http.StatusOK,
		},
		{
			query:        "category_id=1",
			response:     mockAdData[1:],
			expectedCode: http.StatusOK,
		},
		{
			query:        "",
			response:     mockAdData,
			expectedCode: http.StatusOK,
		},
		{
			query:        "category_id=2&price=1000",
			response:     mockAdData[1:],
			expectedCode: http.StatusOK,
		},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/ads?"+v.query, nil)
		w := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, w)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		a.List(c)

		if v.expectedCode == http.StatusOK {
			var adRes []models.Ad
			err := json.Unmarshal(w.Body.Bytes(), &adRes)
			assert.NoError(t, err)

			if !reflect.DeepEqual(adRes, v.response) {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, adRes, v.response)
			}
		} else {
			var errorRes string
			err := json.Unmarshal(w.Body.Bytes(), &errorRes)
			assert.NoError(t, err)

			if !reflect.DeepEqual(errorRes, v.expectedError) {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, errorRes, v.expectedError)
			}
		}
	}
}

func TestAdsHandler_ListFilterSort(t *testing.T) {
	testcases := []struct {
		query         string
		response      []models.Ad
		expectedCode  int
		expectedError string
	}{
		{
			query:        "sort=price,desc",
			response:     mockAdData[:1],
			expectedCode: http.StatusOK,
		},
		{
			query:        "sort=price,asc&sort=category_id,desc",
			response:     mockAdData,
			expectedCode: http.StatusOK,
		},
		{
			query:        "sort=price",
			response:     mockAdData[1:],
			expectedCode: http.StatusOK,
		},
		{
			query:        "",
			response:     mockAdData,
			expectedCode: http.StatusOK,
		},
		{
			query:         "sort=plane_age,asc&sort=favourite_colour,desc",
			expectedError: "could not retrieve ads",
			expectedCode:  http.StatusInternalServerError,
		},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/ads?"+v.query, nil)
		w := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, w)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		a.List(c)

		if v.expectedCode == http.StatusOK {
			var adRes []models.Ad
			err := json.Unmarshal(w.Body.Bytes(), &adRes)
			assert.NoError(t, err)

			if !reflect.DeepEqual(adRes, v.response) {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, adRes, v.response)
			}
		} else {
			var errorRes string
			err := json.Unmarshal(w.Body.Bytes(), &errorRes)
			assert.NoError(t, err)

			if !reflect.DeepEqual(errorRes, v.expectedError) {
				t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, errorRes, v.expectedError)
			}
		}
	}
}

func TestAdHandler_AddAd(t *testing.T) {
	e := echo.New()

	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte("invalid_Json")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err := a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Invalid JSON", response.Message)
	})

	t.Run("JSON Without Price", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "XYZ123",
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Input Json doesn't include Price", response.Message)
	})

	t.Run("JSON Without Category", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "XYZ123",
			"Price":         500000,
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Input Json doesn't include Category", response.Message)
	})

	t.Run("non-string category", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "XYZ123",
			"Price":         500000,
			"Category":      54,
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Category should be string !", response.Message)
	})

	t.Run("invalid category name", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "XYZ123",
			"Price":         500000,
			"Category":      "Hello",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Invalid Category Name", response.Message)
	})

	t.Run("non-airline user", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "XYZ123",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[1])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Airlines Can Add an ad!", response.Message)
	})

	t.Run("non-string model", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": 24,
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Plane Model should be string !", response.Message)
	})

	t.Run("non-number price", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "something",
			"Price":         "548000",
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Price should be a number !", response.Message)
	})

	t.Run("non-integer price", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       1000,
			"AirplaneModel": "something",
			"Price":         54.5,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Price should be an integer !", response.Message)
	})

	t.Run("non-integer fly time", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       78.5,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "fly_time should be an integer !", response.Message)
	})

	t.Run("non-boolean repair_check", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   "hello",
			"ExpertCheck":   false,
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Repair Check should be boolean !", response.Message)
	})

	t.Run("non-boolean ExpertCheck", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   false,
			"ExpertCheck":   "bye",
			"PlaneAge":      7,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Expert Check should be boolean !", response.Message)
	})

	t.Run("non-integer age", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      7.25,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Age should be an integer !", response.Message)
	})
	t.Run("invalid age", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "example1.jpg",
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      123,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "The year of the invention of the airplane was 1903 !", response.Message)
	})

	t.Run("non-string image", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         8745,
			"Description":   "This is example ad 1.",
			"Subject":       "Example Ad 1",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      23,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Image should be an url !", response.Message)
	})

	t.Run("non-string subject", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "image",
			"Description":   "This is example ad 1.",
			"Subject":       7852,
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      23,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "subject should be string !", response.Message)
	})

	t.Run("non-string description", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "image",
			"Description":   55,
			"Subject":       "Subject",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      23,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

		var response models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "description should be string !", response.Message)
	})

	t.Run("valid request", func(t *testing.T) {
		addAdReqBody := map[string]interface{}{
			"Image":         "image",
			"Description":   "Desc",
			"Subject":       "Subject",
			"FlyTime":       78,
			"AirplaneModel": "something",
			"Price":         500000,
			"Category":      "small-passenger",
			"RepairCheck":   true,
			"ExpertCheck":   false,
			"PlaneAge":      23,
		}
		jsonData, err := json.Marshal(addAdReqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/ads/add", bytes.NewReader([]byte(jsonData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Set("user", mockUserData[0])

		a := New(mockDatastore{})
		err = a.AddAdHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.Ad
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		//check some fields
		assert.Equal(t, "image", response.Image)
		assert.Equal(t, "Desc", response.Description)
		assert.Equal(t, "Subject", response.Subject)
	})

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

	mockAdData = []models.Ad{
		{
			ID:            1,
			UserID:        1,
			Image:         "example1.jpg",
			Description:   "This is example ad 1.",
			Subject:       "Example Ad 1",
			Price:         1000,
			CategoryID:    2,
			Status:        string(consts.INACTIVE),
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
			Role:     consts.ROLE_ADMIN,
		},
	}
)

type mockDatastore struct {
	ads_data []models.Ad
}

func (m mockDatastore) Get(id int, userRole string) ([]models.Ad, error) {
	if id == 1 && userRole == "Airline" {
		return mockAdData[:1], nil
	} else if id == 2 {
		return nil, errors.New("db error")
	}

	return mockAdData, nil
}

func (m mockDatastore) ListFilterByColumn(f *filter.AdsFilter) ([]models.Ad, error) {
	if f.PlaneAge == 7 {
		return mockAdData[:1], nil
	}
	if f.CategoryID == 1 {
		return mockAdData[1:], nil
	}
	if f.CategoryID == 2 && f.Price == 1000 {
		return mockAdData[1:], nil
	}

	return mockAdData, nil
}

func (m mockDatastore) ListFilterSort(f *filter.Filter) ([]models.Ad, error) {
	var orderClause []string
	for col, order := range f.Sort {
		orderClause = append(orderClause, fmt.Sprintf("%s %s", col, order))
	}
	order := strings.Join(orderClause, ",")

	if order == "price DESC" {
		return mockAdData[:1], nil
	}

	if order == "price ASC" {
		return mockAdData[1:], nil
	}

	if order == "price ASC,category_id DESC" {
		return mockAdData, nil
	}
	if order == "" {
		return mockAdData, nil
	}

	return nil, fmt.Errorf("no such column: age")
}

func (m mockDatastore) GetCategoryByName(name string) (models.Category, error) {
	if name == "small-passenger" {
		return mockCategoryData[0], nil
	} else if name == "big-passenger" {
		return mockCategoryData[1], nil
	}
	return models.Category{}, errors.New("Database Error")
}

func (m mockDatastore) CreateAd(a *models.Ad) (models.Ad, error) {
	ex_size := len(m.ads_data)
	a.ID = uint(ex_size) + 1
	a.Status = string(consts.INACTIVE)
	m.ads_data = append(m.ads_data, *a)
	return *a, nil
}

func (m mockDatastore) UpdateStatus(id int, status consts.AdStatus) (models.Ad, error) {
	s := mockAdData[id].Status
	if s == string(consts.ACTIVE) {
		mockAdData[id].Status = "Active"
	} else {
		mockAdData[id].Status = "Inactive"
	}
	return mockAdData[id], nil
}
