package handlers

import (
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AdsHandler struct {
	db *gorm.DB
}

func NewAdsHandler(db *gorm.DB) *AdsHandler {
	return &AdsHandler{db: db}
}

func (a AdsHandler) AddAdHandler(c echo.Context) error {
	// Read Request Body
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	//check json format
	jsonFormatValidationMsg, jsonFormatErr := utils.ValidateJsonFormat(jsonBody, "price", "category", "fly_time", "model", "repair_check", "expert_check", "age")
	if jsonFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: jsonFormatValidationMsg})
	}

	//check ad properties validation
	adFormatValidationMsg, ad, adFormatErr := utils.ValidateAd(jsonBody, a.db)
	if adFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: adFormatValidationMsg})
	}

	//set user id
	// user := c.Get("user")
	// user = user.(models.User)
	// id := uint(user.(models.User).ID)
	// ad.UserID = id

	createdAd := a.db.Create(&ad)
	if createdAd.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: "Ad Cration Failed"})
	}

	return c.JSON(http.StatusOK, ad)

}
