package models

import "Airplane-Divar/utils"

type UpdateExpertCheckRequest struct {
	Status utils.ExpertStatus `json:"status"`
	Report string             `json:"report"`
}
