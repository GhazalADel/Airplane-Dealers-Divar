package models

import "time"

type Transaction struct {
	ID              uint      `gorm:"primary_key"`
	UserID          uint      `gorm:"not null"`
	TransactionType string    `gorm:"type:varchar(255);not null"`
	ObjectID        uint      `gorm:"type:bigint;not null"`
	Amount          int64     `gorm:"type:bigint"`
	Status          string    `gorm:"type:varchar(255)"`
	Authority       string    `gorm:"type:varchar(255)"`
	CreatedAt       time.Time `gorm:"type:datetime"`
	User            User
}

func (Transaction) TableName() string {
	return "transactions"
}
