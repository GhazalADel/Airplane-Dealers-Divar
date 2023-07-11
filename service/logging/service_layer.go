package logging_service

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"Airplane-Divar/service"
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

func (loggingService *Logging) GetAdsActivity(ID int) ([]models.ActivityLog, error) {
	// TODO
	// Some Manipulation On Data !
	// no need for payment, bookmark and buy logs
	// excludeLogID := []uint{9, 10, 11, 12, 13}

	return loggingService.loggingDatastore.GetAdsActivityByID(ID)
}

func (loggingService *Logging) ReportActivity(causerType string, causerID int, subjectType string, subjectID int, logTitle string) error {

	logName := loggingService.loggingDatastore.FindLogByTitle(logTitle)
	if logName.Title == "" {
		return fmt.Errorf("invalid logname %v", logTitle)
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
