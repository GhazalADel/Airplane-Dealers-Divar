package models

type AdminAds struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null"`
	Image         string `gorm:"type:varchar(255)"`
	Description   string `gorm:"type:text"`
	Subject       string `gorm:"type:varchar(255);not null"`
	Price         uint64 `gorm:"type:uint;not null"`
	CategoryID    uint   `gorm:"not null"`
	FlyTime       uint   `gorm:"type:uint"`
	AirplaneModel string `gorm:"type:varchar(255)"`
	RepairCheck   bool   `gorm:"type:boolean"`
	ExpertCheck   bool   `gorm:"type:boolean"`
	PlaneAge      uint   `gorm:"type:uint"`
	User          User
	Category      Category
}

func (AdminAds) TableName() string {
	return "admin_ads"
}
