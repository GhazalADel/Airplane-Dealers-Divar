package models

import (
	"Airplane-Divar/utils"
	"time"
)

type RepairRequest struct {
	ID        uint         `gorm:"primary_key"`
	Status    utils.Status `gorm:"type:status_type"`
	CreatedAt time.Time    `gorm:"default:CURRENT_TIMESTAMP()"`
	AdsID     uint         `gorm:"type:bigint;not null"`
	UserID    uint         `gorm:"type:uint;not null"`
	User      User         `gorm:"foreignKey:UserID"`
	Ads       Ad
}

func (RepairRequest) TableName() string {
	return "repair_request"
}
