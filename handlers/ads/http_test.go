package ads

import (
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
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
			response:     mockData,
			expectedCode: http.StatusOK,
		},
		{
			id:           "1",
			response:     mockData[:1],
			expectedCode: http.StatusOK,
		},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/ads/"+v.id, nil)
		w := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, w)

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
			response:     mockData[:1],
			expectedCode: http.StatusOK,
		},
		{
			query:        "category_id=1",
			response:     mockData[1:],
			expectedCode: http.StatusOK,
		},
		{
			query:        "",
			response:     mockData,
			expectedCode: http.StatusOK,
		},
		{
			query:        "category_id=2&price=1000",
			response:     mockData[1:],
			expectedCode: http.StatusOK,
		},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/ads?"+v.query, nil)
		w := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, w)

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
			response:     mockData[:1],
			expectedCode: http.StatusOK,
		},
		{
			query:        "sort=price,asc&sort=category_id,desc",
			response:     mockData,
			expectedCode: http.StatusOK,
		},
		{
			query:        "sort=price",
			response:     mockData[1:],
			expectedCode: http.StatusOK,
		},
		{
			query:        "",
			response:     mockData,
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

var (
	mockData = []models.Ad{
		{
			ID:            1,
			UserID:        1,
			Image:         "example1.jpg",
			Description:   "This is example ad 1.",
			Subject:       "Example Ad 1",
			Price:         1000,
			CategoryID:    2,
			Status:        "Active",
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
			Status:        "Active",
			FlyTime:       1000,
			AirplaneModel: "ABC456",
			RepairCheck:   true,
			ExpertCheck:   true,
			PlaneAge:      3,
		},
	}
)

type mockDatastore struct{}

func (m mockDatastore) Get(id int) ([]models.Ad, error) {
	if id == 1 {
		return mockData[:1], nil
	} else if id == 2 {
		return nil, errors.New("db error")
	}

	return mockData, nil
}

func (m mockDatastore) ListFilterByColumn(f *filter.AdsFilter) ([]models.Ad, error) {
	if f.PlaneAge == 7 {
		return mockData[:1], nil
	}
	if f.CategoryID == 1 {
		return mockData[1:], nil
	}
	if f.CategoryID == 2 && f.Price == 1000 {
		return mockData[1:], nil
	}

	return mockData, nil
}

func (m mockDatastore) ListFilterSort(f *filter.Filter) ([]models.Ad, error) {
	var orderClause []string
	for col, order := range f.Sort {
		orderClause = append(orderClause, fmt.Sprintf("%s %s", col, order))
	}
	order := strings.Join(orderClause, ",")

	if order == "price DESC" {
		return mockData[:1], nil
	}

	if order == "price ASC" {
		return mockData[1:], nil
	}

	if order == "price ASC,category_id DESC" {
		return mockData, nil
	}
	if order == "" {
		return mockData, nil
	}

	return nil, fmt.Errorf("no such column: age")
}
