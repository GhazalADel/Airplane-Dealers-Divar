package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Token    string `gorm:"not null"`
	Role     string `gorm:"type:varchar(255);not null"`
	IsActive bool   `gorm:"default:true"`
}

func (User) TableName() string {
	return "users"
}
