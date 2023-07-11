package models

import (
	"time"
)

type Response struct {
	ResponseCode uint16 `json:"responsecode"`
	Message      string `json:"message"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Error string
}

type GetExpertRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	ExpertID  int       `json:"expertID"`
	AdSubject string    `json:"adSubject"`
	Status    string    `json:"status"`
	Report    string    `json:"report"`
	CreatedAt time.Time `json:"createdAt"`
}

type ExpertRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	AdID      int       `json:"adID"`
	ExpertID  int       `json:"expertID"`
	Status    string    `json:"status"`
	Report    string    `json:"report"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetRepairRequestResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
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

type ActivityLogResponse struct {
	ID          uint      `json:"ID"`
	CreatedAt   time.Time `json:"LoggedAt"`
	CauserType  string    `json:"CauserType"`
	CauserID    uint      `json:"CauserID"`
	SubjectType string    `json:"SubjectType"`
	SubjecrID   uint      `json:"SubjectId"`
	LogName     string    `json:"LogName"`
	Description string    `json:"Description"`
}
