package models

type User struct {
	ID         uint   `gorm:"primary_key"`
	FirstName  string `gorm:"type:varchar(255);not null"`
	LastName   string `gorm:"type:varchar(255);not null"`
	Phone      string `gorm:"type:varchar(255);unique;not null"`
	Email      string `gorm:"type:varchar(255);unique;not null"`
	NationalID string `gorm:"type:varchar(255);unique;not null"`
}

func (User) TableName() string {
	return "users"
}
