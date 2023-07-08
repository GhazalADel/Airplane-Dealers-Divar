package datastore

import (
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
)

type Ad interface {
	Get(id int) ([]models.Ad, error)
	ListFilterByColumn(f *filter.AdsFilter) ([]models.Ad, error)
	ListFilterSort(f *filter.Filter) ([]models.Ad, error)
	GetCategoryByName(name string) (models.Category, error)
	CreateAdminAd(ad *models.AdminAds) (models.AdminAds, error)
}
