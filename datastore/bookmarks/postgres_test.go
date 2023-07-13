package bookmarks

import (
	"Airplane-Divar/consts"
	database "Airplane-Divar/database"
	"Airplane-Divar/models"
	"errors"
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

	cleanup2 := createCategories(t, db)
	defer cleanup2()

	cleanup3 := createAds(t, db)
	defer cleanup3()

	cleanup4 := createBookmarks(t, db)
	defer cleanup4()

	a := New(db)
	//testBookmarkStorer_GetAdsByUserID(t, a)
	//testBookmarkStorer_AddBookmark(t, a)
	testBookmarkStorer_DeleteBookmark(t, a)
}

func testBookmarkStorer_GetAdsByUserID(t *testing.T, db BookmarkDatastorer) {
	testcases := []struct {
		user_id int
		res     []models.AdResponse
	}{
		{1, []models.AdResponse{
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
		}},
		{2, []models.AdResponse{
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
		}},
	}

	for i, v := range testcases {
		resp, _ := db.GetAdsByUserID(v.user_id)

		if !reflect.DeepEqual(resp, v.res) {
			t.Errorf("[GetAdsByUserID() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.res)
		} else {
			fmt.Println("[GetAdsByUserID() TEST", i+1, "]Pass.")
		}
	}
}

func testBookmarkStorer_AddBookmark(t *testing.T, db BookmarkDatastorer) {
	testcases := []struct {
		userID int
		adID   int
		res    models.BookmarksResponse
	}{
		{1,
			2,
			models.BookmarksResponse{
				UserID: 1,
				AdsID:  2,
			},
		},
		{2,
			3,
			models.BookmarksResponse{
				UserID: 2,
				AdsID:  3,
			},
		},
	}
	for i, v := range testcases {
		resp, _ := db.AddBookmark(v.userID, v.adID)

		if !reflect.DeepEqual(resp, v.res) {
			t.Errorf("[AddBookmark() TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.res)
		} else {
			fmt.Println("[AddBookmark() TEST", i+1, "]Pass.")
		}
	}
}

func testBookmarkStorer_DeleteBookmark(t *testing.T, db BookmarkDatastorer) {
	testcases := []struct {
		userID int
		adID   int
		res    error
	}{
		{1,
			1,
			nil,
		},
		{
			1,
			3,
			errors.New(""),
		},
	}
	for i, v := range testcases {
		err := db.DeleteBookmark(v.userID, v.adID)

		if (v.res == nil && err == nil) || (v.res != nil && err != nil) {
			fmt.Println("[DeleteBookmark() TEST", i+1, "]Pass.")
		} else {
			t.Errorf("[DeleteBookmark() TEST%d]Failed. Got %v\tExpected %v\n", i+1, err, v.res)
		}
	}
}

func createUser(t *testing.T, db *gorm.DB) func() {
	users := []models.User{
		{
			ID:       1,
			Username: "john_doe",
			Password: "password123",
			Role:     consts.ROLE_AIRLINE,
		},
		{
			ID:       2,
			Username: "rose_parker",
			Password: "pass",
			Role:     consts.ROLE_AIRLINE,
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			t.Fatal(err)
		}
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

func createBookmarks(t *testing.T, db *gorm.DB) func() {
	bookmarks := []models.Bookmarks{
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
	for _, book := range bookmarks {
		if err := db.Create(&book).Error; err != nil {
			t.Fatal(err)
		}
	}
	return func() {
		db.Exec("DELETE FROM bookmarks")
	}

}
