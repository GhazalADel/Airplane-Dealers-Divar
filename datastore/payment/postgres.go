package user

import (
	"Airplane-Divar/models"

	"gorm.io/gorm"
)

type PaymentStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) PaymentStore {
	return PaymentStore{db: db}
}

func (u PaymentStore) Get(id int) ([]models.Transaction, error) {
	return []models.Transaction{}, nil
}

func (u PaymentStore) Create(user models.Transaction) (models.Transaction, error) {
	return models.Transaction{}, nil
}
