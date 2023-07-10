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
	CreateAd(ad *models.Ad) (models.Ad, error)
}
