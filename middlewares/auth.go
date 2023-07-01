package middlewares

import (
	database "Airplane-Divar/database"
	"Airplane-Divar/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		tokenString := req.Header.Get("Authorization")

		//Account Doesn't have Token
		if tokenString == "" {
			return echo.ErrConflict
		}

		//Parse Token
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			//Wrong Algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected sigining method %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil

		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			//Check Expiration Time
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return echo.ErrUnauthorized
			}

			//Connect To Database
			db, err := database.GetConnection()
			if err != nil {
				return echo.ErrInternalServerError
			}

			//Find Account
			var account models.Account
			db.First(&account, claims["id"])

			//Token And Id are not For Same Accounts
			if account.ID == 0 {
				return echo.ErrUnauthorized
			}
			//account is deactive
			if account.Token == "" && !account.IsActive {
				return echo.ErrUnauthorized
			}

			//Add Account Object To Context
			c.Set("account", account)
			return next(c)

		} else {
			return echo.ErrUnauthorized
		}
	}
}
