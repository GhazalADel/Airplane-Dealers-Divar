package models

import (
	"Airplane-Divar/consts"
	"time"
)

type ExpertAds struct {
	ID        uint          `gorm:"primary_key"`
	Report    string        `gorm:"type:text"`
	Status    consts.Status `gorm:"type:status_type"`
	CreatedAt time.Time     `gorm:"default:current_timestamp"`
	Expert    User          `gorm:"foreignKey:ExpertID"`
	ExpertID  uint          `gorm:"type:bigint"`
	AdsID     uint          `gorm:"type:bigint;not null"`
	UserID    uint          `gorm:"type:uint;not null"`
	User      User          `gorm:"foreignKey:UserID"`
	Ads       Ad
}

func (ExpertAds) TableName() string {
	return "expert_ads"
}
