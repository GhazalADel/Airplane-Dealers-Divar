package models

type User struct {
	ID        uint        `gorm:"primaryKey"`
	Username  string      `gorm:"type:varchar(255);unique;not null"`
	Password  string      `gorm:"type:varchar(255);not null"`
	Role      int         `gorm:"type:int;not null"`
	Bookmarks []Bookmarks `gorm:"many2many:bookmarks"`
}

func (User) TableName() string {
	return "users"
}
