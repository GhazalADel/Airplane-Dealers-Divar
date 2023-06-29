package account

import (
	"Airplane-Divar/models"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) AccountStore {
	return AccountStore{db: db}
}

func (a AccountStore) Get(id int) ([]models.Account, error) {
	return []models.Account{}, nil
}

// this function used to create an account and insert it into database
func (a AccountStore) Create(user_id int, username string, is_admin bool, password string) (string, models.Account, error) {
	msg := "OK"
	// Instantiating Account Object
	var account models.Account
	account.UserID = uint(user_id)
	account.Username = username
	account.IsAdmin = is_admin
	account.Token = ""

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		msg = "Failed to Hashing Password"
		return msg, models.Account{}, errors.New("")
	}
	account.Password = string(hash)

	//insert account into database
	createdAccount := a.db.Create(&account)
	if createdAccount.Error != nil {
		msg = "Failed to Create Account"
		return msg, models.Account{}, errors.New("")
	}

	//generate token
	var token *jwt.Token
	if is_admin {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    account.ID,
			"exp":   time.Now().Add(time.Hour).Unix(),
			"admin": true,
		})
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    account.ID,
			"exp":   time.Now().Add(time.Hour).Unix(),
			"admin": false,
		})
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		msg = "Failed To Create Token"
		return msg, models.Account{}, errors.New("")
	}
	account.Token = tokenString

	//update account
	a.db.Save(&account)

	return msg, account, nil

}

func (a AccountStore) CheckUnique(username string) (string, error) {
	msg := "OK"
	// Is Input Username Unique or Not
	var existingAccount models.Account
	a.db.Where("username = ?", username).First(&existingAccount)
	if existingAccount.ID != 0 {
		msg = "Inupt Username has already been registered"
		return msg, errors.New("")
	}
	return msg, nil
}
