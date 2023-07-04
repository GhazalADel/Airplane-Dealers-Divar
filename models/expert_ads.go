package models

type ExpertAds struct {
	ID       uint   `gorm:"primaryKey"`
	Report   string `gorm:"type:varchar(255)"`
	Status   int    `gorm:"type:int"`
	ExpertID uint   `gorm:"type:uint"`
	AdsID    uint   `gorm:"type:uint;not null"`
	Ad       Ad     `gorm:"foreignKey:AdsID"`
	Expert   User   `gorm:"foreginKey:ExpertID"`
}

func (ExpertAds) TableName() string {
	return "expert_ads"
}
