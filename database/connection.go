package db

import (
	"Airplane-Divar/config"
	"Airplane-Divar/models"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConn *gorm.DB

func Connect() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	// -------env----------
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.PG.HOST, cfg.PG.USER, cfg.PG.PASSWORD, cfg.PG.DB, cfg.PG.PORT, cfg.PG.SSLMODE, cfg.PG.TIMEZONE)
	// -------env----------

	// If not connect - use "db" instead of "localhost"
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	// 	"localhost", "postgres", "root", "divar_airplane", "5432", "disable", "Asia/Tehran")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	dbConn = db
	//initializeDatabase()
	return nil
}

func GetConnection() (*gorm.DB, error) {
	if dbConn == nil {
		err := Connect()
		if err != nil {
			return nil, errors.New("database connection is not initialized")
		}
	}
	return dbConn, nil
}

func CloseDatabase(db *gorm.DB) error {
	postDB, err := db.DB()
	if err != nil {
		return err
	}
	err = postDB.Close()
	if err != nil {
		return err
	}

	return nil
}

func initializeDatabase() {
	//check categories table for tuples
	var categories []models.Category
	dbConn.Select("id").Find(&categories)
	if len(categories) == 0 {
		//insert categories
		categories := []string{"small-passenger", "big-passenger"}
		for _, v := range categories {
			var c models.Category
			c.Name = v
			dbConn.Save(&c)
		}

	}

}
