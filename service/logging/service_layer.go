package logging_service

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"Airplane-Divar/service"
	"encoding/json"
	"fmt"
)

type Logging struct {
	loggingDatastore datastore.Logging
}

var logService *Logging

func Initialize(loggingDataStorer datastore.Logging) {
	if logService == nil {
		logService = &Logging{
			loggingDatastore: loggingDataStorer,
		}
	}
}

func GetInstance() service.Logging {

	return logService
}

func excludeLogIds(logID uint) bool {
	// no need for payment logs
	excludeLogId := []uint{9, 10, 11}

	for _, lid := range excludeLogId {
		if logID == lid {
			return true
		}
	}
	return false
}

func (loggingService *Logging) GetAdsActivity(ID int) ([]byte, error) {

	adsLogs, err := loggingService.loggingDatastore.GetAdsActivityByID(ID)
	if err != nil {
		return nil, err
	}

	resp := []models.ActivityLogResponse{}

	for _, v := range adsLogs {
		if !excludeLogIds(v.LogID) {
			logResp := models.ActivityLogResponse{
				ID:          v.ID,
				CreatedAt:   v.CreatedAt,
				CauserType:  v.CauserType,
				CauserID:    v.CauserID,
				SubjectType: v.SubjectType,
				SubjectID:   v.SubjectID,
				LogName:     loggingService.loggingDatastore.GetLogNameByID(v.LogID),
				Description: v.Description,
			}

			resp = append(resp, logResp)

		}
	}

	return json.MarshalIndent(resp, "", "	")
}

func (loggingService *Logging) ReportActivity(causerType string, causerID uint, subjectType string, subjectID uint, logTitle string, description string) error {

	logName := loggingService.loggingDatastore.FindLogByTitle(logTitle)
	if logName.Title == "" {
		return fmt.Errorf("logname %v not found", logTitle)
	}

	alog := models.ActivityLog{
		CauserType:  causerType,
		CauserID:    uint(causerID),
		SubjectType: subjectType,
		SubjectID:   uint(subjectID),
		Log:         logName,
		LogID:       logName.ID,
		Description: description,
	}

	return loggingService.loggingDatastore.AddActivity(alog)
}
