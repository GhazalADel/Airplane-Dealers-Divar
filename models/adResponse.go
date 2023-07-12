package models

type AdResponse struct {
	ID            uint
	UserID        uint
	Image         string
	Description   string
	Subject       string
	Price         uint64
	CategoryID    uint
	Status        string
	FlyTime       uint
	AirplaneModel string
	RepairCheck   bool
	ExpertCheck   bool
	PlaneAge      uint
}
