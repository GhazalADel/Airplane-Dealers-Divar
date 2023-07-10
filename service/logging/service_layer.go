package logging_service

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"Airplane-Divar/service"
)

type Logging struct {
	LoggingDatastore datastore.Logging
}

func (loggingSrv *Logging) New(loggingDataStorer datastore.Logging) service.Logging {
	return &Logging{
		LoggingDatastore: loggingDataStorer,
	}
}

func (loggingSrv *Logging) GetAdsActivity(ID int) ([]models.ActivityLog, error) {
	// TODO
	// Some Manipulation On Data !

	// no need for payment, bookmark and buy logs
	// excludeLogID := []uint{9, 10, 11, 12, 13}

	return loggingSrv.LoggingDatastore.GetAdsActivityByID(ID)
}

func (loggingSrv *Logging) ReportActivity(alog models.ActivityLog) error {
	// TODO
	return loggingSrv.LoggingDatastore.AddActivity(alog)
}
