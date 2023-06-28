package models

import "time"

type Ad struct {
	ID            uint      `gorm:"primary_key"`
	UserID        uint      `gorm:"not null"`
	Image         string    `gorm:"type:varchar(255)"`
	Description   string    `gorm:"type:text"`
	Subject       string    `gorm:"type:varchar(255);not null"`
	Price         int64     `gorm:"type:bigint;not null"`
	CategoryID    uint      `gorm:"not null"`
	Status        string    `gorm:"type:varchar(255)"`
	FlyTime       time.Time `gorm:"type:datetime"`
	AirplaneModel string    `gorm:"type:varchar(255)"`
	RepairCheck   bool      `gorm:"type:boolean"`
	ExpertCheck   bool      `gorm:"type:boolean"`
	PlaneAge      int       `gorm:"type:int"`
}

func (Ad) TableName() string {
	return "ads"
}
