package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

func (Category) TableName() string {
	return "categories"
}
