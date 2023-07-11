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

func New(loggingDataStorer datastore.Logging) service.Logging {
	return &Logging{
		loggingDatastore: loggingDataStorer,
	}
}

func excludeLogIds(logID uint) bool {
	excludeLogId := []uint{9, 10, 11, 12, 13}

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
				LogName:     v.Log.Title,
				Description: v.Description,
			}

			resp = append(resp, logResp)

		}
	}

	return json.Marshal(resp)
}

func (loggingService *Logging) ReportActivity(causerType string, causerID int, subjectType string, subjectID int, logTitle string) error {

	logName := loggingService.loggingDatastore.FindLogByTitle(logTitle)
	if logName.Title == "" {
		return fmt.Errorf("logname %v not found", logTitle)
	}

	alog := models.ActivityLog{
		CauserType:  causerType,
		CauserID:    uint(causerID),
		SubjectType: subjectType,
		SubjecrID:   uint(subjectID),
		Log:         logName,
		LogID:       logName.ID,
	}

	return loggingService.loggingDatastore.AddActivity(alog)
}
