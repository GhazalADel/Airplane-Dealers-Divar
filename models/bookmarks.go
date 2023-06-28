package models

type Bookmarks struct {
	UserID uint `gorm:"primaryKey"`
	AdsID  uint `gorm:"primaryKey"`
	User   User
	Ads    Ads
}

func (Bookmarks) TableName() string {
	return "bookmarks"
}
