package datastore

import "Airplane-Divar/models"

type User interface {
	Get(id int) ([]models.User, error)
	Create(username string, password string) (string, models.User, error)
	Login(username, password string) (string, models.User, error)
	CheckUnique(username string) (string, error)
}

type Payment interface {
	Create(userID uint, fee int64, authority string) (string, error)
}

type Logging interface {
	AddNewLogName(id uint, title string) error
	AddActivity(al models.ActivityLog) error
	GetAdsActivityByID(id int) ([]models.ActivityLog, error)
}
