package service

import "Airplane-Divar/models"

type Logging interface {
	GetAdsActivity(ID int) ([]models.ActivityLog, error)
	ReportActivity(causerType string, causerID int, subjectType string, subjectID int, logTitle string) error
}
