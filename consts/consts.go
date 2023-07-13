package consts

import "database/sql/driver"

type Status string

const (
	WAIT_FOR_PAYMENT_STATUS Status = "Wait for payment status"
	EXPERT_PENDING_STATUS   Status = "Pending for expert"
	MATIN_PENDING_STATUS    Status = "Pending for matin"
	IN_PROGRESS_STATUS      Status = "In progress"
	DONE_STATUS             Status = "Done"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

// paginator
const PAGE_SIZE int = 10

// User roles
const (
	ROLE_MATIN   string = "Matin"
	ROLE_EXPERT         = "Expert"
	ROLE_ADMIN          = "Admin"
	ROLE_AIRLINE        = "Airline"
)

// Logs
const (
	LOG_CREATE_AD       string = "create_ads"
	LOG_ADMIN_WAIT      string = "send_to_admin"
	LOG_ADMIN_APPROVE   string = "admin_approved"
	LOG_ADMIN_REJECT    string = "admin_reject"
	LOG_REPAIR_REQUEST  string = "repair_request"
	LOG_REPAIR_RESULT   string = "repair_result"
	LOG_EXPERT_REQUEST  string = "expert_request"
	LOG_EXPERT_RESULT   string = "expert_result"
	LOG_PAYMENT         string = "payment"
	LOG_PAYMENT_SUCCESS string = "payment_success"
	LOG_PAYMENT_FAILED  string = "payment_failed"
	LOG_BOOKMARK        string = "bookmark"
	LOG_BOOKMARK_REMOVE string = "bookmark_remove"
)
