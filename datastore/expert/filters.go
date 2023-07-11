package expert

import (
	"Airplane-Divar/consts"
	"time"

	"gorm.io/gorm/clause"
)

type FilterAndConditionExpertRequest struct {
	Status   consts.Status
	FromDate string
	ExpertID int
	UserID   int
	AdsID    int
}

func (q FilterAndConditionExpertRequest) ToQueryModel() (clause.AndConditions, error) {
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

	if q.ExpertID > 0 {
		queryModel.Exprs = append(
			queryModel.Exprs, clause.Eq{Column: "expert_id", Value: q.ExpertID},
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

type FilterOrConditionExpertRequest struct {
	ExpertIDList []interface{}
	StatusList   []consts.Status
}

func (q FilterOrConditionExpertRequest) ToQueryModel() clause.OrConditions {
	queryModel := clause.OrConditions{}

	if len(q.ExpertIDList) > 0 {
		for _, val := range q.ExpertIDList {

			queryModel.Exprs = append(
				queryModel.Exprs, clause.Eq{Column: "expert_id", Value: val},
			)
		}
	}

	if len(q.StatusList) > 0 {
		for _, val := range q.StatusList {

			queryModel.Exprs = append(
				queryModel.Exprs, clause.Eq{Column: "status", Value: val},
			)
		}
	}

	return queryModel
}

type FilterNotConditionExpertRequest struct {
	Status consts.Status
}

func (q FilterNotConditionExpertRequest) ToQueryModel() clause.NotConditions {
	queryModel := clause.NotConditions{}

	if q.Status != "" {
		queryModel.Exprs = append(
			queryModel.Exprs, clause.Eq{Column: "status", Value: q.Status},
		)
	}

	return queryModel
}
