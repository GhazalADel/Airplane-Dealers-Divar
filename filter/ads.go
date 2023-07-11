package filter

import (
	"Airplane-Divar/utils"
	"net/url"
)

type AdsFilter struct {
	Base Filter

	PlaneAge      uint
	Price         int64
	FlyTime       uint64
	CategoryID    uint
	AirplaneModel string
	Status        string
}

func NewAdsFilter(v url.Values) *AdsFilter {
	f := New(v)

	return &AdsFilter{
		Base:          *f,
		PlaneAge:      utils.Uint(v.Get("plane_age")),
		Price:         utils.Int64(v.Get("price")),
		FlyTime:       utils.Uint64(v.Get("fly_time")),
		CategoryID:    utils.Uint(v.Get("category_id")),
		AirplaneModel: v.Get("airplen_model"),
	}
}
