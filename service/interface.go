package service

type (
	Logging interface {
		GetAdsActivity(ID int) ([]byte, error)
		ReportActivity(
			causerType string,
			causerID uint,
			subjectType string,
			subjectID uint,
			logTitle string,
			description string,
		) error
	}
)
