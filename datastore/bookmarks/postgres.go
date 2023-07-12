package bookmarks

import (
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookmarkDatastorer struct {
	db *gorm.DB
}

func New(db *gorm.DB) BookmarkDatastorer {
	return BookmarkDatastorer{db: db}
}

func (b BookmarkDatastorer) GetAdsByUserID(id int) ([]models.AdResponse, error) {
	var user models.User
	b.db.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return []models.AdResponse{}, fmt.Errorf("Invalid UserID")
	}

	var bookmarks []models.Bookmarks

	res := b.db.Where("user_id = ?", id).Find(&bookmarks)
	if res.Error != nil {
		return []models.AdResponse{}, fmt.Errorf("Database Failed")
	}

	var ads []models.AdResponse
	for _, book := range bookmarks {
		var ad models.Ad
		b.db.Where("id = ?", book.AdsID).First(&ad)
		adRes := models.AdResponse{
			ID:            ad.ID,
			UserID:        ad.UserID,
			Image:         ad.Image,
			Description:   ad.Description,
			Subject:       ad.Subject,
			Price:         ad.Price,
			CategoryID:    ad.CategoryID,
			Status:        ad.Status,
			FlyTime:       ad.FlyTime,
			AirplaneModel: ad.AirplaneModel,
			RepairCheck:   ad.RepairCheck,
			ExpertCheck:   ad.ExpertCheck,
			PlaneAge:      ad.PlaneAge,
		}
		// if ad.ID == 0 {
		// 	b.db.Delete(&book)
		// } else {
		ads = append(ads, adRes)
		//}
	}
	return ads, nil
}

func (b BookmarkDatastorer) AddBookmark(userID, adID int) (models.BookmarksResponse, error) {
	var user models.User
	b.db.Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		return models.BookmarksResponse{}, fmt.Errorf("Invalid UserID")
	}

	var ad models.Ad
	b.db.Where("id = ?", adID).First(&ad)
	if ad.ID == 0 {
		return models.BookmarksResponse{}, fmt.Errorf("Invalid AdID")
	}

	if ad.Status != string(utils.ACTIVE) {
		return models.BookmarksResponse{}, fmt.Errorf("You can't bookmark this ad")
	}

	bookmark := models.Bookmarks{
		UserID: uint(userID),
		AdsID:  uint(adID),
	}

	res := b.db.Create(&bookmark)
	if res.Error != nil {
		return models.BookmarksResponse{}, fmt.Errorf("You bookmarked this ad before")
	}

	bookRes := models.BookmarksResponse{
		UserID: bookmark.UserID,
		AdsID:  bookmark.AdsID,
	}

	return bookRes, nil
}

func (b BookmarkDatastorer) DeleteBookmark(userID, adID int) error {
	var user models.User
	b.db.Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		return fmt.Errorf("Invalid UserID")
	}

	var ad models.Ad
	b.db.Where("id = ?", adID).First(&ad)
	if ad.ID == 0 {
		return fmt.Errorf("Invalid AdID")
	}

	var book models.Bookmarks
	res := b.db.Where("user_id = ? AND ads_id = ?", userID, adID).First(&book)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("You didn't bookmark this ad before!")
		}
		return fmt.Errorf("Database failed")
	}

	res = b.db.Delete(&book)
	if res.Error != nil {
		return fmt.Errorf("Failed to Delete Bookmark")
	}
	return nil
}
