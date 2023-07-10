package repair

import (
	"Airplane-Divar/utils"
	"time"

	"gorm.io/gorm/clause"
)

type FilterAndConditionRepairRequest struct {
	Status   utils.Status
	FromDate string
	UserID   int
	AdsID    int
}

func (q FilterAndConditionRepairRequest) ToQueryModel() (clause.AndConditions, error) {
	queryModel := clause.AndConditions{}

	if q.UserID > 0 {
		queryModel.Exprs = append(
			queryModel.Exprs, clause.Eq{Column: "user_id", Value: q.UserID},
		)
	}

	if q.Status != "" {
		queryModel.Exprs = append(
			queryModel.Exprs, clause.Eq{Column: "status", Value: q.Status},
		)
	}

	if q.AdsID > 0 {
		queryModel.Exprs = append(
			queryModel.Exprs, clause.Eq{Column: "ads_id", Value: q.AdsID},
		)
	}

	if q.FromDate != "" {
		t, err := time.Parse("2006-01-02", q.FromDate)
		if err != nil {
			return queryModel, err
		}

		queryModel.Exprs = append(queryModel.Exprs, clause.Gte{Column: "created_at", Value: t})
	}

	return queryModel, nil
}

type FilterNotConditionRepairRequest struct {
	Status utils.Status
}

func (q FilterNotConditionRepairRequest) ToQueryModel() clause.NotConditions {
	queryModel := clause.NotConditions{}

	if q.Status != "" {
		queryModel.Exprs = append(
			queryModel.Exprs, clause.Eq{Column: "status", Value: q.Status},
		)
	}

	return queryModel
}
