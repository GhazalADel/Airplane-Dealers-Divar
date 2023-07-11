package models

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
