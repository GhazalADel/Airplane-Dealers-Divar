package models

import "time"

type ActivityLog struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time `gorm:"type:datetime;autoCreateTime"`
	SubjectType string    `gorm:"type:varchar(50);not null"`
	SubjecrID   uint      `gorm:"type:bigint;not null"`
	CauserType  string    `gorm:"type:varchar(50)"`
	CauserID    uint      `gorm:"type:bigint"`
	Log         LogName   `gorm:"foreignKey:LogID"`
	LogID       uint      `gorm:"type:int"`
	Description string    `gorm:"type:varchar(255)"`
}

func (ActivityLog) TableName() string {
	return "activity_log"
}
