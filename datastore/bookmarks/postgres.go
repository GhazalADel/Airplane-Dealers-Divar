package bookmarks

import (
	"Airplane-Divar/models"
	"fmt"

	"gorm.io/gorm"
)

type BookmarkDatastorer struct {
	db *gorm.DB
}

func New(db *gorm.DB) BookmarkDatastorer {
	return BookmarkDatastorer{db: db}
}

func (b BookmarkDatastorer) GetAdsByUserID(id int) ([]models.Ad, error) {
	var user models.User
	b.db.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return []models.Ad{}, fmt.Errorf("Invalid UserID")
	}

	var ads []models.Ad

	res := b.db.Where("user_id = ?", id).Find(&ads)
	if res.Error != nil {
		return []models.Ad{}, fmt.Errorf("Database Failed")
	}
	return ads, nil
}

func (b BookmarkDatastorer) AddBookmark(userID, adID int) (models.Bookmarks, error) {
	var user models.User
	b.db.Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		return models.Bookmarks{}, fmt.Errorf("Invalid UserID")
	}

	var ad models.Ad
	b.db.Where("id = ?", adID).First(&ad)
	if ad.ID == 0 {
		return models.Bookmarks{}, fmt.Errorf("Invalid AdID")
	}

	bookmark := models.Bookmarks{
		UserID: uint(userID),
		AdsID:  uint(adID),
	}

	res := b.db.Create(&bookmark)
	if res.Error != nil {
		return models.Bookmarks{}, fmt.Errorf("Database Failed")
	}

	return bookmark, nil
}
