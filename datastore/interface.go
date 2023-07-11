package datastore

import (
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
)

type Ad interface {
	// Create(models.Ad) (models.Ad, error)
	Get(id int, userRole string) ([]models.Ad, error)
	ListFilterByColumn(f *filter.AdsFilter) ([]models.Ad, error)
	ListFilterSort(f *filter.Filter) ([]models.Ad, error)
	GetCategoryByName(name string) (models.Category, error)
	CreateAd(ad *models.Ad) (models.Ad, error)


type User interface {
	Get(id int) ([]models.User, error)
	Create(username string, password string) (string, models.User, error)
	Login(username, password string) (string, models.User, error)
	CheckUnique(username string) (string, error)
}

type Payment interface {
	Create(userID uint, fee int64, authority string) (string, error)
}
