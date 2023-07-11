package expert

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"context"
	"errors"
	"log"

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

	if user.Role == consts.ROLE_EXPERT { // is_expert
		query.Where("expert_ads.expert_id = ? OR expert_ads.expert_id IS NULL", user.ID)
	} else if user.Role == consts.ROLE_AIRLINE { // is advertiser
		query.Where(&models.Ad{UserID: user.ID})
	}
	result := query.First(&expertAd)

	return expertAd, result.Error
}

func (e ExpertStorer) Get(
	ctx context.Context,
	requestID int,
	user models.User,
) (models.ExpertAds, error) {
	var expertAd models.ExpertAds
	query := e.db.WithContext(ctx).Joins("Ads").Where("expert_ads.id = ?", requestID)

	if user.Role == consts.ROLE_EXPERT { // is_expert
		query.Where("expert_ads.expert_id = ? OR expert_ads.expert_id IS NULL", user.ID)
	} else if user.Role == consts.ROLE_AIRLINE { // is advertiser
		query.Where(&models.Ad{UserID: user.ID})
	}
	result := query.First(&expertAd)

	return expertAd, result.Error
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
		if expertAd.Status != consts.EXPERT_PENDING_STATUS {
			return errors.New("you had been requested for expert check")
		}
		return nil
	}

	if err := e.db.WithContext(ctx).Model(&models.ExpertAds{}).
		Create(map[string]interface{}{
			"AdsID":  ad.ID,
			"UserID": user.ID,
			"Status": consts.WAIT_FOR_PAYMENT_STATUS,
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
		log.Println(filterOrCondition, "\n\n")
		for _, filter := range filterOrCondition {
			if len(filter.Exprs) > 0 {
				query.Where(filter)
			}
		}
	}
	if len(filterNotCondtion.Exprs) > 0 {
		query.Where(filterNotCondtion)
	}

	result := query.Find(&expertRequests)
	return expertRequests, result.Error
}

func (e ExpertStorer) Update(
	ctx context.Context, expertAdID int, user models.User, body models.UpdateExpertCheckRequest,
) (models.ExpertAds, error) {
	tmpExpertAd := models.ExpertAds{}
	updatedMap := make(map[string]interface{})

	if body.Status != "" {
		updatedMap["status"] = body.Status
		if body.Status == consts.EXPERT_PENDING_STATUS {
			updatedMap["expert_id"] = nil
		} else if body.Status == consts.WAIT_FOR_PAYMENT_STATUS {
			return tmpExpertAd, errors.New("not allowed")
		} else {
			updatedMap["expert_id"] = user.ID
		}
	}
	if body.Report != "" {
		updatedMap["report"] = body.Report
		updatedMap["status"] = consts.DONE_STATUS
		updatedMap["expert_id"] = user.ID
	}

	result := e.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Model(&tmpExpertAd).
		Where(
			"id = ? AND (expert_id = ? OR expert_id IS NULL) AND status != ?",
			expertAdID, user.ID, consts.WAIT_FOR_PAYMENT_STATUS,
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
			user.ID, adID, consts.WAIT_FOR_PAYMENT_STATUS,
		).
		Delete(&models.ExpertAds{})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("expert request not found")
	}
	return nil

}
