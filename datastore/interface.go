package datastore

import "Airplane-Divar/models"

type Ad interface {
	// Create(models.Ad) (models.Ad, error)
	// Update(id int) (models.Ad, error)
	// Delete(id int) error
	Get(id int) ([]models.Ad, error)
}
