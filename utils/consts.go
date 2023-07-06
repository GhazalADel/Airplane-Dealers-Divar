package utils

import (
	"database/sql/driver"
)

type ExpertStatus string

const (
	EXPERT_WAIT_FOR_PAYMENT_STATUS ExpertStatus = "Wait for payment status"
	EXPERT_PENDING_STATUS          ExpertStatus = "Pending for expert"
	EXPERT_IN_PROGRESS_STATUS      ExpertStatus = "In progress"
	EXPERT_CONFIRMED_STATUS        ExpertStatus = "Confirmed"
)

func (ct *ExpertStatus) Scan(value interface{}) error {
	*ct = ExpertStatus(value.(string))
	return nil
}

func (ct ExpertStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

// paginator
const PAGE_SIZE int = 10
