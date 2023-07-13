package service

type (
	Logging interface {
		GetAdsActivity(ID int) ([]byte, error)
		ReportActivity(
			causerType string,
			causerID int,
			subjectType string,
			subjectID int,
			logTitle string,
		) error
	}
)
