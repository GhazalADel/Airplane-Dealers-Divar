package utils

import (
	"Airplane-Divar/models"
	"errors"
	"strings"
)

// this function used to check user properties validation
func ValidateUser(jsonBody map[string]interface{}) (string, models.User, error) {
	msg := "OK"
	// Create User Object
	var user models.User
	user.FirstName = jsonBody["firstname"].(string)
	user.LastName = jsonBody["lastname"].(string)
	user.Email = jsonBody["email"].(string)
	user.Phone = jsonBody["phone"].(string)
	user.NationalID = jsonBody["nationalid"].(string)

	// Check FirstName Validation
	if len(strings.TrimSpace(user.FirstName)) == 0 {
		msg = "First Name can't be empty"
		return msg, models.User{}, errors.New("")
	}

	// Check LastName Validation
	if len(strings.TrimSpace(user.LastName)) == 0 {
		msg = "Last Name can't be empty"
		return msg, models.User{}, errors.New("")
	}

	// Check Phone Number Validation
	if !ValidatePhone(user.Phone) {
		msg = "Invalid Phone Number"
		return msg, models.User{}, errors.New("")
	}

	// Check Email Validation
	if !ValidateEmail(user.Email) {
		msg = "Invalid Email Address"
		return msg, models.User{}, errors.New("")
	}

	// Check NationalID Validation
	if !ValidateNationalID(user.NationalID) {
		msg = "Invalid National ID"
		return msg, models.User{}, errors.New("")
	}
	return msg, user, nil
}

// func Login(username, password string, is_admin bool, db *gorm.DB) (string, models.Account, error) {
// 	msg := "OK"
// 	// Find account based on input username
// 	var account models.Account

// 	db.Where("username = ?", username).First(&account)

// 	// Account Not Found
// 	if account.ID == 0 {
// 		msg = "Invalid Username"
// 		return msg, models.Account{}, errors.New("")
// 	}

// 	// Incorrect Password
// 	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
// 	if err != nil {
// 		msg = "Wrong Password"
// 		return msg, models.Account{}, errors.New("")
// 	}

// 	if is_admin {
// 		if !account.IsAdmin {
// 			msg = "You are not admin!"
// 			return msg, models.Account{}, errors.New("")
// 		}
// 	}
// 	//Account isn't active
// 	if !account.IsActive {
// 		msg = "Your Account Isn't Active"
// 		return msg, models.Account{}, errors.New("")
// 	}

// 	var token *jwt.Token
// 	if is_admin {
// 		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 			"id":    account.ID,
// 			"exp":   time.Now().Add(time.Hour).Unix(),
// 			"admin": true,
// 		})
// 	} else {
// 		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 			"id":    account.ID,
// 			"exp":   time.Now().Add(time.Hour).Unix(),
// 			"admin": false,
// 		})
// 	}

// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
// 	if err != nil {
// 		msg = "Failed To Create Token"
// 		return msg, models.Account{}, errors.New("")
// 	}

// 	// Update Account's Token In Database
// 	account.Token = tokenString
// 	db.Save(&account)
// 	return msg, account, nil
// }
