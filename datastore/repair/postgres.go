package repair

import (
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RepairStorer struct {
	db *gorm.DB
}

func NewRepairStorer(db *gorm.DB) RepairStorer {
	return RepairStorer{db: db}
}

func (e RepairStorer) RequestToRepairCheck(
	ctx context.Context, adID int, userID int,
) error {
	// get user
	var user models.User
	if err := e.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return err
	}

	// get ad
	var ad models.Ad
	if err := e.db.WithContext(ctx).First(&ad, adID).Error; err != nil {
		return err
	}
	if ad.RepairCheck {
		return errors.New("this had been already repaired by Matin")
	}

	// get or create repair request
	var repairRequest models.RepairRequest
	err := e.db.WithContext(ctx).
		Where("user_id = ? AND ads_id = ?", userID, adID).
		First(&repairRequest).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return err
	}

	if repairRequest.ID != 0 {
		if repairRequest.Status != utils.MATIN_PENDING_STATUS {
			return errors.New("you had been requested for repairing")
		}
		return nil
	}

	if err := e.db.WithContext(ctx).Model(&models.RepairRequest{}).
		Create(map[string]interface{}{
			"AdsID":  ad.ID,
			"UserID": user.ID,
			"Status": utils.WAIT_FOR_PAYMENT_STATUS,
		}).Error; err != nil {
		return err
	}

	return nil
}
