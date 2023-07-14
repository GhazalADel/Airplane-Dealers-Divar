package payment

import (
	"Airplane-Divar/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) PaymentStore {
	return PaymentStore{db: db}
}

func (p PaymentStore) Create(
	userID uint,
	fee int64,
	authority string,
	transactionType string,
	objectID uint,
) (string, error) {
	var transaction models.Transaction
	transaction.UserID = userID
	transaction.TransactionType = transactionType
	transaction.ObjectID = objectID
	transaction.Amount = fee
	transaction.Status = "Wait"
	transaction.Authority = authority
	transaction.CreatedAt = time.Now()

	// Insert Transaction Object Into Database
	createdTransaction := p.db.Create(&transaction)
	if createdTransaction.Error != nil {
		return "Transaction Cration Failed", createdTransaction.Error
	}
	return "okay", nil
}

func (p PaymentStore) Update(status string, idList []uint) error {
	err := p.db.Table("transactions").
		Where("id IN ?", idList).
		Updates(
			map[string]interface{}{"status": status},
		).Error

	return err
}

func (p PaymentStore) FindByAuthority(authority string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := p.db.Where(&models.Transaction{Authority: authority}).Find(&transactions)

	return transactions, result.Error
}

func (p PaymentStore) GetPriceByServices(ctx context.Context, services []string) (map[string]float64, error) {
	var prices []models.Configuration
	result := make(map[string]float64)
	err := p.db.WithContext(ctx).Where("name IN ?", services).Find(&prices).Error
	if err != nil {
		return result, err
	}

	if len(prices) > 0 {
		for _, p := range prices {
			result[p.Name] = p.Value
		}
	}

	return result, nil
}

func (p PaymentStore) GetTotalPriceByServices(prices map[string]float64) float64 {
	var sum float64 = 0
	for _, val := range prices {
		sum += val
	}

	return sum
}
