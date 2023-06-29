package ads

import (
	"Airplane-Divar/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

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
			CategoryID:    1,
			Status:        "Active",
			FlyTime:       timeStr(),
			AirplaneModel: "XYZ123",
			RepairCheck:   true,
			ExpertCheck:   false,
			PlaneAge:      5,
		},
		{
			ID:            2,
			UserID:        1,
			Image:         "example2.jpg",
			Description:   "This is example ad 2.",
			Subject:       "Example Ad 2",
			Price:         2000,
			CategoryID:    2,
			Status:        "Active",
			FlyTime:       timeStr(),
			AirplaneModel: "ABC456",
			RepairCheck:   true,
			ExpertCheck:   true,
			PlaneAge:      3,
		},
	}
)

func timeStr() time.Time {
	layout := "2006-01-02 15:04:05"  // Define the layout for parsing the time string
	timeStr := "2023-06-29 12:30:00" // Specify the specific time as a string

	// Parse the time string using the specified layout
	flyTime, _ := time.Parse(layout, timeStr)
	return flyTime
}

type mockDatastore struct{}

func (m mockDatastore) Get(id int) ([]models.Ad, error) {
	if id == 1 {
		return mockData[:1], nil
	} else if id == 2 {
		return nil, errors.New("db error")
	}

	return mockData, nil
}
