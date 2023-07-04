package datastore

import (
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
)

type Ad interface {
	// Create(models.Ad) (models.Ad, error)
	Get(id int) ([]models.Ad, error)
	ListFilterByColumn(f *filter.AdsFilter) ([]models.Ad, error)
	ListFilterSort(f *filter.Filter) ([]models.Ad, error)
}
