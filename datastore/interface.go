package datastore

import (
	"Airplane-Divar/models"
	"context"

	"gorm.io/gorm/clause"
)

type (
	Expert interface {
		RequestToExpertCheck(ctx context.Context, adID int, userID int) error
		GetAllExpertRequests(
			ctx context.Context,
			filterAndCondition clause.AndConditions,
			filterOrCondition []clause.OrConditions,
			filterNotCondtion clause.NotConditions,
			page int,
		) ([]models.ExpertAds, error)
		Get(ctx context.Context, expertRequestID int) (models.ExpertAds, error)
		GetByAd(
			ctx context.Context,
			adID int,
			user models.User,
		) (models.ExpertAds, error)
		Update(
			ctx context.Context, expertAdID int,
			user models.User, body models.UpdateExpertCheckRequest,
		) (models.ExpertAds, error)
		Delete(
			ctx context.Context,
			adID int,
			user models.User,
		) error
	}

	Repair interface {
		RequestToRepairCheck(ctx context.Context, adID int, userID int) error
		GetByAd(
			ctx context.Context,
			adID int,
			user models.User,
		) (models.RepairRequest, error)
		// Get(ctx context.Context, repairRequestID int) (models.RepairRequest, error)
		// GetAllRepairRequests(
		// 	ctx context.Context,
		// 	filterAndCondition clause.AndConditions,
		// 	filterOrCondition []clause.OrConditions,
		// 	filterNotCondtion clause.NotConditions,
		// 	page int,
		// ) ([]models.RepairRequest, error)
		// Update(
		// 	ctx context.Context, repairRequestID int,
		// 	user models.User, body models.UpdateRepairRequest,
		// ) (models.RepairRequest, error)
		// Delete(
		// 	ctx context.Context,
		// 	adID int,
		// 	user models.User,
		// ) error
	}

	User interface {
		Get(ctx context.Context, id int) (models.User, error)
	}
)
