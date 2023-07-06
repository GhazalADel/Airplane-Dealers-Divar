package user

import (
	"Airplane-Divar/models"
	"context"

	"gorm.io/gorm"
)

type UserStorer struct {
	db *gorm.DB
}

func NewUserStorer(db *gorm.DB) UserStorer {
	return UserStorer{db: db}
}

func (u UserStorer) Get(ctx context.Context, id int) (models.User, error) {
	var user models.User
	if err := u.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil

}
