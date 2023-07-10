package models

import "Airplane-Divar/utils"

type UpdateExpertCheckRequest struct {
	Status utils.Status `json:"status"`
	Report string       `json:"report"`
}

type UpdateRepairRequest struct {
	Status utils.Status `json:"status"`
}
