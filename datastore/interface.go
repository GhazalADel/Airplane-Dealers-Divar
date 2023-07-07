package datastore

import (
	"Airplane-Divar/models"
)

type LoggingDataLayer interface {
	AddNewLogName(id uint, title string) error
	ReportActivity(al models.ActivityLog) error
	GetAdsActivity(id int) ([]models.ActivityLog, error)
}
