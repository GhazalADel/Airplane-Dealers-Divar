package db

import (
	"Airplane-Divar/config"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	dbConn = db
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
