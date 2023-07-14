package datastore

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	"context"

	"gorm.io/gorm/clause"
)

type (
	Ad interface {
		UpdateStatus(id int, status consts.AdStatus) (models.Ad, error)
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
		Get(
			ctx context.Context,
			requestID int,
			user models.User,
		) (models.ExpertAds, error)
		GetByAd(
			ctx context.Context,
			adID int,
			user models.User,
		) (models.ExpertAds, error)
		Update(
			ctx context.Context, expertAdID int, upadtedColumn map[string]interface{},
		) error
		UpdateByExpert(
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
		Get(
			ctx context.Context,
			requestID int,
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
			ctx context.Context, repairRequestID int, upadtedColumn map[string]interface{},
		) error
		UpdateByUser(
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
		Create(username string, password string, role string) (string, models.User, error)
		Login(username, password string) (string, models.User, error)
		CheckUnique(username string) (string, error)
	}

	Payment interface {
		Create(
			userID uint,
			fee int64,
			authority string,
			transactionType string,
			objectID uint,
		) (string, error)
		Update(status string, idList []uint) error
		FindByAuthority(authority string) ([]models.Transaction, error)
		GetPriceByServices(ctx context.Context, services []string) (map[string]float64, error)
		GetTotalPriceByServices(prices map[string]float64) float64
	}
	Bookmark interface {
		GetAdsByUserID(id int) ([]models.AdResponse, error)
		AddBookmark(userID, adID int) (models.BookmarksResponse, error)
		DeleteBookmark(userID, adID int) error
	}

	Logging interface {
		AddNewLogName(id uint, title string) error
		FindLogByTitle(title string) models.LogName
		AddActivity(al models.ActivityLog) error
		GetAdsActivityByID(id int) ([]models.ActivityLog, error)
		GetLogNameByID(id uint) string
	}
)
