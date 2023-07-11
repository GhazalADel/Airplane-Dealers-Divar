package datastore

import (
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	"context"

	"gorm.io/gorm/clause"
)

type (
	Ad interface {
		// Create(models.Ad) (models.Ad, error)
		Get(id int, userRole string) ([]models.Ad, error)
		ListFilterByColumn(f *filter.AdsFilter) ([]models.Ad, error)
		ListFilterSort(f *filter.Filter) ([]models.Ad, error)
		GetCategoryByName(name string) (models.Category, error)
		CreateAd(ad *models.Ad) (models.Ad, error)
	}

	Expert interface {
		RequestToExpertCheck(ctx context.Context, adID int, user models.User) error
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
		RequestToRepairCheck(ctx context.Context, adID int, user models.User) error
		GetByAd(
			ctx context.Context,
			adID int,
			user models.User,
		) (models.RepairRequest, error)
		GetAllRepairRequests(
			ctx context.Context,
			filterAndCondition clause.AndConditions,
			filterOrCondition []clause.OrConditions,
			filterNotCondtion clause.NotConditions,
			page int,
		) ([]models.RepairRequest, error)
		Update(
			ctx context.Context, repairRequestID int,
			user models.User, body models.UpdateRepairRequest,
		) (models.RepairRequest, error)
		Delete(
			ctx context.Context,
			adID int,
			user models.User,
		) error
	}

	User interface {
		Get(id int) ([]models.User, error)
		Create(username string, password string) (string, models.User, error)
		Login(username, password string) (string, models.User, error)
		CheckUnique(username string) (string, error)
	}

	Payment interface {
		Create(userID uint, fee int64, authority string) (string, error)
	}
)
