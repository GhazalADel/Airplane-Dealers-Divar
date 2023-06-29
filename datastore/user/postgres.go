package user

import (
	"Airplane-Divar/models"
	"errors"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserStore {
	return UserStore{db: db}
}

func (u UserStore) Get(id int) ([]models.User, error) {
	return []models.User{}, nil
}

func (u UserStore) Create(user models.User) (models.User, error) {
	res := u.db.Create(&user)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	return user, nil
}

func (u UserStore) CheckUnique(user models.User) (string, error) {
	msg := "Ok"
	// Is Input Phone Number Unique or Not
	var existingUser models.User
	u.db.Where("phone = ?", user.Phone).First(&existingUser)
	if existingUser.ID != 0 {
		msg = "Inupt Phone Number has already been registered"
		return msg, errors.New("")
	}

	// Is Input Email Address Unique or Not
	u.db.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		msg = "Inupt Email Address has already been registered"
		return msg, errors.New("")
	}

	// Is Input National ID Unique or Not
	u.db.Where("national_id = ?", user.NationalID).First(&existingUser)
	if existingUser.ID != 0 {
		msg = "Inupt National ID has already been registered"
		return msg, errors.New("")
	}
	return msg, nil
}
