package ads

import (
	database "Airplane-Divar/database"
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"fmt"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestDatastore(t *testing.T) {
	db, err := database.CreateTestDatabase()
	defer database.CloseTestDatabase(db)
	if err != nil {
		t.Errorf("could not connect to sql, err: %v", err)
	}

	cleanup1 := createUser(t, db)
	defer cleanup1()

	cleanup2 := createAds(t, db)
	defer cleanup2()

	cleanup3 := createCategories(t, db)
	defer cleanup3()

	a := New(db)
	testAdStorer_Get(t, a)
	testAdStorer_ListFilterByColumn(t, a)
	testAdStorer_ListFilterSort(t, a)
	testAdStorer_GetCategoryByName(t, a)
	testAdStorer_CreateAd(t, a)
}

func testAdStorer_Get(t *testing.T, db AdDatastorer) {
	testcases := []struct {
		id   int
		resp []models.Ad
	}{
		{0, []models.Ad{
			{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
			{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
			// {3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
		}},
		{1, []models.Ad{{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}}}},
	}

	for i, v := range testcases {
		resp, _ := db.Get(v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[Get() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		} else {
			fmt.Println("[Get() TEST", i+1, "]Pass.")
		}
	}
}

func testAdStorer_ListFilterByColumn(t *testing.T, db AdDatastorer) {
	base := &filter.Filter{
		Offset: -10,
		Limit:  10,
	}

	testcases := []struct {
		filter filter.AdsFilter
		resp   []models.Ad
	}{
		{
			filter.AdsFilter{
				Base:     *base,
				PlaneAge: 7,
			},
			[]models.Ad{{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}}},
		},
		{
			filter.AdsFilter{
				Base:       *base,
				CategoryID: 1,
			},
			[]models.Ad{
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
				//{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
			},
		},
		{
			filter.AdsFilter{
				Base:       *base,
				CategoryID: 1,
				Price:      1000,
			},
			[]models.Ad{
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
				//{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				// {3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
			},
		},
		{
			filter.AdsFilter{
				Base:       *base,
				CategoryID: 1,
				// Price:      1000,
				FlyTime: 1000,
			},
			[]models.Ad{
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
				//{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
			},
		},
		{
			filter.AdsFilter{
				Base: *base,
				// CategoryID: 1,
				// Price:      1000,
				// FlyTime: 1000,
			},
			[]models.Ad{
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
				{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
			},
		},
	}
	for i, v := range testcases {
		resp, _ := db.ListFilterByColumn(&v.filter)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[ListFilterByColumn() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		} else {
			fmt.Println("[ListFilterByColumn() TEST", i+1, "]Pass.")
		}
	}
}

func testAdStorer_ListFilterSort(t *testing.T, db AdDatastorer) {

	testcases := []struct {
		filter filter.Filter
		resp   []models.Ad
	}{
		{
			filter.Filter{
				Offset: -10,
				Limit:  10,
				Sort: map[string]string{
					"price": "ASC",
				},
			},
			[]models.Ad{
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
				{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
			},
		},
		{
			filter.Filter{
				Offset: -10,
				Limit:  10,
				Sort: map[string]string{
					"price": "DESC",
				},
			},
			[]models.Ad{
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
				{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
			},
		},
		{
			filter.Filter{
				Offset: -10,
				Limit:  10,
				Sort: map[string]string{
					"price":     "DESC",
					"plane_age": "ASC",
				},
			},
			[]models.Ad{
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
				{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
			},
		},
		{
			filter.Filter{
				Offset: -10,
				Limit:  10,
				Sort: map[string]string{
					"age":  "DESC",
					"year": "ASC",
				},
			},
			nil,
		},
		{
			filter.Filter{
				Offset: -10,
				Limit:  10,
			},
			[]models.Ad{
				{1, 1, "example1.jpg", "This is example ad 1.", "Example Ad 1", 1000, 1, "Active", 1000, "XYZ123", true, false, 5, models.User{}, models.Category{}},
				{2, 1, "example2.jpg", "This is example ad 2.", "Example Ad 2", 2000, 2, "Active", 1000, "ABC456", true, true, 3, models.User{}, models.Category{}},
				{3, 1, "example3.jpg", "This is example ad 3.", "Example Ad 3", 3000, 1, "Inactive", 1000, "DEF789", false, false, 7, models.User{}, models.Category{}},
			},
		},
	}
	for i, v := range testcases {
		resp, _ := db.ListFilterSort(&v.filter)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[ListFilterSort() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		} else {
			fmt.Println("[ListFilterSort() TEST", i+1, "]Pass.")
		}
	}
}

func testAdStorer_GetCategoryByName(t *testing.T, db AdDatastorer) {
	testcases := []struct {
		name string
		res  models.Category
	}{
		{"small-passenger", models.Category{ID: 1, Name: "small-passenger"}},
		{"big-passenger", models.Category{ID: 2, Name: "big-passenger"}},
	}

	for i, v := range testcases {
		resp, _ := db.GetCategoryByName(v.name)

		if !reflect.DeepEqual(resp, v.res) {
			t.Errorf("[GetCategoryByName() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.res)
		} else {
			fmt.Println("[GetCategoryByName() TEST", i+1, "]Pass.")
		}
	}
}

func testAdStorer_CreateAd(t *testing.T, db AdDatastorer) {
	testcases := []struct {
		ad  models.Ad
		res models.Ad
	}{
		{models.Ad{
			UserID:        1,
			Image:         "",
			Description:   "Description2",
			Subject:       "Subject2",
			Price:         100000,
			CategoryID:    1,
			FlyTime:       50,
			AirplaneModel: "Good Model2",
			Status:        string(utils.INACTIVE),
			RepairCheck:   true,
			ExpertCheck:   false,
			PlaneAge:      3,
		},
			models.Ad{
				ID:            4,
				UserID:        1,
				Image:         "",
				Description:   "Description2",
				Subject:       "Subject2",
				Price:         100000,
				CategoryID:    1,
				FlyTime:       50,
				AirplaneModel: "Good Model2",
				Status:        string(utils.INACTIVE),
				RepairCheck:   true,
				ExpertCheck:   false,
				PlaneAge:      3,
			},
		},
	}

	for i, v := range testcases {
		resp, _ := db.CreateAd(&v.ad)

		if !reflect.DeepEqual(resp, v.res) {
			t.Errorf("[CreateAdminAd() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.res)
		} else {
			fmt.Println("[CreateAdminAd() TEST", i+1, "]Pass.")
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

	return func() {
		db.Exec("DELETE FROM users")
	}
}

func createAds(t *testing.T, db *gorm.DB) func() {
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
			FlyTime:       1000,
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
			Status:        "Inactive",
			FlyTime:       1000,
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
	}

	return func() {
		db.Exec("DELETE FROM ads")
	}
}

func createCategories(t *testing.T, db *gorm.DB) func() {
	categories := []models.Category{
		{
			ID:   1,
			Name: "small-passenger",
		},
		{
			ID:   2,
			Name: "big-passenger",
		},
	}

	for _, cat := range categories {
		if err := db.Create(&cat).Error; err != nil {
			t.Fatal(err)
		}
	}

	return func() {
		db.Exec("DELETE FROM categories")
	}
}
