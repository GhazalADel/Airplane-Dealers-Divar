package ads

import (
	database "Airplane-Divar/database"
	"Airplane-Divar/models"
	"fmt"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

var (
	fly_time time.Time
)

func TestDatastore(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("could not connect to sql, err: %v", err)
	}

	cleanup1 := createUser(t, db)
	defer cleanup1()

	cleanup2 := createAds(t, db)
	defer cleanup2()

	a := New(db)
	testAdStorer_Get(t, a)

}

func testAdStorer_Get(t *testing.T, db AdDatastorer) {
	testcases := []struct {
		id   int
		resp []models.Ad
	}{
		{0, []models.Ad{
			{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", fly_time, "XYZ123", true, false, 5, models.User{}, models.Category{}},
			{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", fly_time, "ABC456", true, true, 3, models.User{}, models.Category{}},
			{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", fly_time, "DEF789", false, false, 7, models.User{}, models.Category{}},
		}},
		{1, []models.Ad{{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", fly_time, "XYZ123", true, false, 5, models.User{}, models.Category{}}}},
	}

	for i, v := range testcases {
		resp, _ := db.Get(v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		}
	}
}

func createUser(t *testing.T, db *gorm.DB) func() {
	user := models.User{
		ID:       1,
		Username: "john_doe",
		Password: "password123",
		Role:     1,
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}

	fmt.Println("User created successfully:", user.ID)

	return func() {
		db.Exec("DELETE FROM users")
	}
}

func createAds(t *testing.T, db *gorm.DB) func() {
	layout := "2006-01-02 15:04:05"  // Define the layout for parsing the time string
	timeStr := "2023-06-29 12:30:00" // Specify the specific time as a string

	// Parse the time string using the specified layout
	flyTime, err := time.Parse(layout, timeStr)
	if err != nil {
		t.Fatal(err)
	}
	fly_time = flyTime

	ads := []models.Ad{
		{
			ID:            1,
			UserID:        1,
			Image:         "example1.jpg",
			Description:   "This is example ad 1.",
			Subject:       "Example Ad 1",
			Price:         1000,
			CategoryID:    1,
			Status:        "Active",
			FlyTime:       flyTime,
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
			FlyTime:       flyTime,
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
			Status:        "Inactive",
			FlyTime:       flyTime,
			AirplaneModel: "DEF789",
			RepairCheck:   false,
			ExpertCheck:   false,
			PlaneAge:      7,
		},
	}

	for _, ad := range ads {
		if err := db.Create(&ad).Error; err != nil {
			t.Fatal(err)
		}
		fmt.Println("Ad created successfully:", ad.ID)
	}

	return func() {
		db.Exec("DELETE FROM ads")
	}
}
