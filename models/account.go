package models

type Account struct {
	ID        uint        `gorm:"primary_key"`
	UserID    uint        `gorm:"not null"`
	Username  string      `gorm:"type:varchar(255);unique;not null"`
	Password  string      `gorm:"type:varchar(255);not null"`
	Token     string      `gorm:"not null"`
	Role      int         `gorm:"type:int;not null"`
	IsActive  bool        `gorm:"default:true"`
	IsAdmin   bool        `gorm:"default:false"`
	Bookmarks []Bookmarks `gorm:"many2many:bookmarks"`
}

func (Account) TableName() string {
	return "accouns"
}
