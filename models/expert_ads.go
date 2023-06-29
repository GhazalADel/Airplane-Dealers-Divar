package models

type ExpertAds struct {
	ID       uint   `gorm:"primary_key"`
	Report   string `gorm:"type:varchar()"`
	Status   int    `gorm:"type:int"`
	Expert   User   `gorm:"foreignKey:ExpertID"`
	ExpertID uint   `gorm:"type:uint"`
	AdsID    uint   `gorm:"type:uint;not null"`
	Ad       Ad
}

func (ExpertAds) TableName() string {
	return "expert_ads"
}
