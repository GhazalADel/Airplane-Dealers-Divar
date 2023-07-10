package logging

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"fmt"

	"gorm.io/gorm"
)

type LoggingStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) datastore.Logging {
	return &LoggingStore{db: db}
}

func (logDL *LoggingStore) AddNewLogName(id uint, title string) error {
	var logname models.LogName
	var result *gorm.DB

	// Check id and title
	result = logDL.db.Where("ID = ? OR Title = ?", id, title).Find(&logname)

	if result.Error != nil {
		return fmt.Errorf("database error: get log_name from database")
	} else if result.RowsAffected != 0 {
		return fmt.Errorf("log_name with id = %v or title = %v is already exists", &id, &title)
	}

	// Add new logname
	result = logDL.db.Create(models.LogName{ID: id, Title: title})

	if result.Error != nil {
		return fmt.Errorf("database error: insert new log_name to database")
	}

	return nil
}

func (logDL *LoggingStore) AddActivity(al models.ActivityLog) error {
	var result *gorm.DB
	var logname models.LogName
	actvtlog := models.ActivityLog{
		SubjectType: al.SubjectType,
		SubjecrID:   al.SubjecrID,
		CauserType:  al.CauserType,
		CauserID:    al.CauserID,
		Log:         al.Log,
		LogID:       al.LogID,
		Description: al.Description,
	}

	// Check for Valid Log Title
	result = logDL.db.Where("id = ?", al.LogID).Find(&logname)
	if result.Error != nil {
		return fmt.Errorf("database error: get log_name from database")
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("invalid log title")
	}

	// Inserting Activity log
	result = logDL.db.Create(actvtlog)
	if result.Error != nil {
		return fmt.Errorf("database error: insert new activity log to database")
	}

	return nil
}

func (logDL *LoggingStore) GetAdsActivityByID(id int) ([]models.ActivityLog, error) {

	var dbResult *gorm.DB
	var activityResult []models.ActivityLog

	dbResult = logDL.db.
		Where("SubjectType = ? AND SubjectID = ? ", "Ads", id).
		Order("CreatedAt").
		Find(&activityResult)

	if dbResult.Error != nil {
		return []models.ActivityLog{}, fmt.Errorf("database error: select activity log from database")
	}

	return activityResult, nil
}
