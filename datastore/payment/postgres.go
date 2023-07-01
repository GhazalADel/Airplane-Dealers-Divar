package payment

import (
	"Airplane-Divar/models"
	"time"

	"gorm.io/gorm"
)

type PaymentStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) PaymentStore {
	return PaymentStore{db: db}
}

func (u PaymentStore) Create(userID uint, fee int64, authority string) (string, error) {
	var transaction models.Transaction
	transaction.UserID = userID
	transaction.Amount = fee
	transaction.Status = "Wait"
	transaction.Authority = authority
	transaction.CreatedAt = time.Now()

	// Insert Transaction Object Into Database
	createdTransaction := u.db.Create(&transaction)
	if createdTransaction.Error != nil {
		return "Transaction Cration Failed", createdTransaction.Error
	}
	return "okay", nil
}
