package ads

import (
	"Airplane-Divar/models"
	"fmt"

	"gorm.io/gorm"
)

type AdDatastorer struct {
	db *gorm.DB
}

func New(db *gorm.DB) AdDatastorer {
	return AdDatastorer{db: db}
}

func (a AdDatastorer) Get(id int) ([]models.Ad, error) {
	var ads []models.Ad
	var result *gorm.DB

	if id != 0 {
		result = a.db.Where("id = ?", id).Find(&ads)
	} else {
		result = a.db.Find(&ads)
	}

	if result.Error != nil {
		return []models.Ad{}, fmt.Errorf("database error: Get ads from database")
	}
	return ads, nil
}
