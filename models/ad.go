package models

type Ad struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null"`
	Image         string `gorm:"type:varchar(255)"`
	Description   string `gorm:"type:text"`
	Subject       string `gorm:"type:varchar(255);not null"`
	Price         uint64 `gorm:"type:uint;not null"`
	CategoryID    uint   `gorm:"not null"`
	Status        string `gorm:"type:varchar(255)"`
	FlyTime       uint   `gorm:"type:uint"`
	AirplaneModel string `gorm:"type:varchar(255)"`
	RepairCheck   bool   `gorm:"type:boolean"`
	ExpertCheck   bool   `gorm:"type:boolean"`
	PlaneAge      uint   `gorm:"type:uint"`
	Category      Category
}

func (Ad) TableName() string {
	return "ads"
}
