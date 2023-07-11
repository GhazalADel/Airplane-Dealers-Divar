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

		//User Doesn't have Token
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

			//Find User
			var user models.User
			db.First(&user, claims["id"])

			//Token And Id are not For Same User
			if user.ID == 0 {
				return echo.ErrUnauthorized
			}
			//user is deactive
			if user.Token == "" && !user.IsActive {
				return echo.ErrUnauthorized
			}

			//Add User Object To Context
			c.Set("user", user)
			return next(c)

		} else {
			return echo.ErrUnauthorized
		}
	}
}
