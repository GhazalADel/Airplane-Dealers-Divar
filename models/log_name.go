package models

type LogName struct {
	ID    uint   `gorm:"type:int;primaryKey"`
	Title string `gorm:"type:varchar(100)"`
}

/*
log list:
	1. 	create_ads
	2. 	send_to_admin
	3. 	admin_approved
	4. 	admin_reject
	5. 	repair_request
	6. 	repair_result
	7. 	expert_request
	8. 	expert_result
	9. 	payment
	10. payment_success
	11. payment_failed
	12. bookmark
	13. bookmark_remove
*/

func (LogName) TableName() string {
	return "log_name"
}

func LogsList() []LogName {
	logs := []LogName{
		{ID: 1, Title: "create_ads"},
		{ID: 2, Title: "send_to_admin"},
		{ID: 3, Title: "admin_approved"},
		{ID: 4, Title: "admin_reject"},
		{ID: 5, Title: "repair_request"},
		{ID: 6, Title: "repair_result"},
		{ID: 7, Title: "expert_request"},
		{ID: 8, Title: "expert_result"},
		{ID: 9, Title: "payment"},
		{ID: 10, Title: "payment_success"},
		{ID: 11, Title: "payment_failed"},
		{ID: 12, Title: "bookmark"},
		{ID: 13, Title: "bookmark_remove"},
	}
	return logs
}
