package models

import (
	"time"
)
type Response struct {
	ResponseCode uint16 `json:"responsecode"`
	Message      string `json:"message"`
}

type MessageResponse struct {
	Message string
}

type SuccessResponse struct {
	Success bool
}

type ErrorResponse struct {
	Error string
}

type GetExpertRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userName"`
	ExpertID  int       `json:"expertID"`
	AdSubject string    `json:"adSubject"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ExpertRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	AdID      int       `json:"adID"`
	ExpertID  int       `json:"expertID"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetRepairRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userName"`
	AdSubject string    `json:"adSubject"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type RepairRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	AdID      int       `json:"adID"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
