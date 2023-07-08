package db

import (
	"Airplane-Divar/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateTestDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Ad{}, &models.ExpertAds{},
		&models.Bookmarks{}, &models.Transaction{}, &models.AdminAds{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseTestDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
