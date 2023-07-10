package user

import (
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserStore {
	return UserStore{db: db}
}

func (a UserStore) Get(id int) ([]models.User, error) {
	return []models.User{}, nil
}

// this function used to create an account and insert it into database
func (a UserStore) Create(username string, password string) (string, models.User, error) {
	msg := "OK"
	// Instantiating Account Object
	var user models.User
	user.Username = username
	user.Role = utils.ROLE_AIRLINE
	user.Token = ""

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		msg = "Failed to Hashing Password"
		return msg, models.User{}, errors.New("")
	}
	user.Password = string(hash)

	//insert account into database
	createdUser := a.db.Create(&user)
	if createdUser.Error != nil {
		msg = "Failed to Create Account"
		return msg, models.User{}, errors.New("")
	}

	//generate token
	var token *jwt.Token
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"admin": false,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		msg = "Failed To Create Token"
		return msg, models.User{}, errors.New("")
	}
	user.Token = tokenString

	//update account
	a.db.Save(&user)

	return msg, user, nil

}

func (a UserStore) CheckUnique(username string) (string, error) {
	msg := "OK"
	// Is Input Username Unique or Not
	var existingUser models.User
	a.db.Where("username = ?", username).First(&existingUser)
	if existingUser.ID != 0 {
		msg = "Inupt Username has already been registered"
		return msg, errors.New("")
	}
	return msg, nil
}

func (a UserStore) Login(username, password string) (string, models.User, error) {
	msg := "OK"
	// Find account based on input username
	var user models.User

	a.db.Where("username = ?", username).First(&user)

	// Account Not Found
	if user.ID == 0 {
		msg = "Invalid Username"
		return msg, models.User{}, errors.New("")
	}

	// Incorrect Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		msg = "Wrong Password"
		return msg, models.User{}, errors.New("")
	}

	//Account isn't active
	if !user.IsActive {
		msg = "Your Account Isn't Active"
		return msg, models.User{}, errors.New("")
	}

	var token *jwt.Token

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"admin": false,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		msg = "Failed To Create Token"
		return msg, models.User{}, errors.New("")
	}

	// Update Account's Token In Database
	user.Token = tokenString
	a.db.Save(&user)
	return msg, user, nil
}
