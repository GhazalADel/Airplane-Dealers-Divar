package models

import "time"

type ActionLog struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `gorm:"type:datetime"`
	SubjectType string    `gorm:"type:varchar(50);not null"`
	SubjecrID   uint      `gorm:"type:bigint;not null"`
	CauserType  string    `gorm:"type:varchar(50)"`
	CauserID    uint      `gorm:"type:bigint"`
	Log         LogName   `gorm:"foreignKey:LogID"`
	LogID       uint      `gorm:"type:bigint"`
	Description string    `gorm:"type:varchar(255)"`
}

func (ActionLog) TableName() string {
	return "activity_log"
}
