package models

type Configuration struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Value float64
}

func (Configuration) TableName() string {
	return "configuration"
}
