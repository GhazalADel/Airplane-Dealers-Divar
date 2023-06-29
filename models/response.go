package models

type Response struct {
	ResponseCode uint16 `json:"responsecode"`
	Message      string `json:"message"`
}
