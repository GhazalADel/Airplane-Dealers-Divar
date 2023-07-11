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
