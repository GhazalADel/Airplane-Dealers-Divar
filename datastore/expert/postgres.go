package expert

import (
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExpertStorer struct {
	db *gorm.DB
}

func NewExpertStorer(db *gorm.DB) ExpertStorer {
	return ExpertStorer{db: db}
}

func (e ExpertStorer) GetByAd(
	ctx context.Context,
	adID int,
	user models.User,
) (models.ExpertAds, error) {
	var expertAd models.ExpertAds
	query := e.db.WithContext(ctx).Joins("Ads").Where("expert_ads.ads_id = ?", adID)

	if user.Role == utils.ROLE_EXPERT { // is_expert
		query.Where("expert_ads.expert_id = ? OR expert_ads.expert_id IS NULL", user.ID)
	} else if user.Role == utils.ROLE_AIRLINE { // is advertiser
		query.Where(&models.Ad{UserID: user.ID})
	}
	result := query.First(&expertAd)

	return expertAd, result.Error
}

func (e ExpertStorer) Get(ctx context.Context, expertRequestID int) (models.ExpertAds, error) {
	var expertAd models.ExpertAds
	if err := e.db.WithContext(ctx).First(&expertAd, expertRequestID).Error; err != nil {
		return expertAd, err
	}

	return expertAd, nil
}

func (e ExpertStorer) RequestToExpertCheck(
	ctx context.Context, adID int, user models.User,
) error {
	// get ad
	var ad models.Ad
	if err := e.db.WithContext(ctx).First(&ad, adID).Error; err != nil {
		return err
	}
	if ad.ExpertCheck {
		return errors.New("this had been already checked by experts")
	}

	// get or create expert_ad
	var expertAd models.ExpertAds
	err := e.db.WithContext(ctx).
		Where("user_id = ? AND ads_id = ?", user.ID, adID).
		First(&expertAd).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}

	if expertAd.ID != 0 {
		if expertAd.Status != utils.EXPERT_PENDING_STATUS {
			return errors.New("you had been requested for expert check")
		}
		return nil
	}

	if err := e.db.WithContext(ctx).Model(&models.ExpertAds{}).
		Create(map[string]interface{}{
			"AdsID":  ad.ID,
			"UserID": user.ID,
			"Status": utils.WAIT_FOR_PAYMENT_STATUS,
		}).Error; err != nil {
		return err
	}

	return nil
}

func (e ExpertStorer) GetAllExpertRequests(
	ctx context.Context,
	filterAndCondition clause.AndConditions,
	filterOrCondition []clause.OrConditions,
	filterNotCondtion clause.NotConditions,
	page int,
) ([]models.ExpertAds, error) {
	expertRequests := []models.ExpertAds{}

	query := e.db.WithContext(ctx).Scopes(utils.Paginate(page))
	if len(filterAndCondition.Exprs) > 0 {
		query.Where(filterAndCondition)
	}
	if len(filterOrCondition) > 0 {
		for _, filter := range filterOrCondition {
			query.Where(filter)
		}
	}
	if len(filterNotCondtion.Exprs) > 0 {
		query.Where(filterNotCondtion)
	}

	result := query.Find(&expertRequests)
	return expertRequests, result.Error
}

type FilterAndConditionExpertRequest struct {
	Status   utils.Status
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
	StatusList   []utils.Status
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
	Status utils.Status
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

func (e ExpertStorer) Update(
	ctx context.Context, expertAdID int, user models.User, body models.UpdateExpertCheckRequest,
) (models.ExpertAds, error) {
	tmpExpertAd := models.ExpertAds{}
	updatedMap := make(map[string]interface{})

	if body.Status != "" {
		updatedMap["status"] = body.Status
		if body.Status == utils.EXPERT_PENDING_STATUS {
			updatedMap["expert_id"] = nil
		} else if body.Status == utils.WAIT_FOR_PAYMENT_STATUS {
			return tmpExpertAd, errors.New("not allowed")
		} else {
			updatedMap["expert_id"] = user.ID
		}
	}
	if body.Report != "" {
		updatedMap["report"] = body.Report
		updatedMap["status"] = utils.DONE_STATUS
	}

	result := e.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Model(&tmpExpertAd).
		Where(
			"id = ? AND (expert_id = ? OR expert_id IS NULL)",
			expertAdID, user.ID,
		).
		Updates(updatedMap)

	return tmpExpertAd, result.Error
}

func (e ExpertStorer) Delete(
	ctx context.Context,
	adID int,
	user models.User,
) error {
	result := e.db.WithContext(ctx).
		Where(
			"user_id = ? AND ads_id = ? AND status = ?",
			user.ID, adID, utils.WAIT_FOR_PAYMENT_STATUS,
		).
		Delete(&models.ExpertAds{})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("expert request not found")
	}
	return nil

}
