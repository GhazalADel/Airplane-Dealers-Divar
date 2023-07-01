package models

type LogName struct {
	ID    uint   `gorm:"type:bigint;primary_key"`
	Title string `gorm:"type:varchar(100)"`
}

func (LogName) TableName() string {
	return "log_name"
}
