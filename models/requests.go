package models

import "Airplane-Divar/consts"

type UpdateExpertCheckRequest struct {
	Status consts.Status `json:"status"`
	Report string        `json:"report"`
}

type UpdateRepairRequest struct {
	Status consts.Status `json:"status"`
}
