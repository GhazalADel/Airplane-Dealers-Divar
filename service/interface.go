package service

import "Airplane-Divar/models"

type Logging interface {
	GetAdsActivity(ID int) ([]models.ActivityLog, error)
	ReportActivity(models.ActivityLog) error
}
