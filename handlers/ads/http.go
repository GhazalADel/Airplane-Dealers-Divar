package ads

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdsHandler struct {
	datastore datastore.Ad
}

func New(ads datastore.Ad) *AdsHandler {
	return &AdsHandler{datastore: ads}
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
	var ad models.Ad
	// TODO
	//check ad properties validation
	// adFormatValidationMsg, ad, adFormatErr := utils.ValidateAd(jsonBody, a.db)
	// if adFormatErr != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: adFormatValidationMsg})
	// }

	//set user id
	// user := c.Get("user")
	// user = user.(models.User)
	// id := uint(user.(models.User).ID)
	// ad.UserID = id

	// TODO
	// createdAd := a.datastore.Create(&ad)
	// if createdAd.Error != nil {
	// 	return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: "Ad Cration Failed"})
	// }
	return c.JSON(http.StatusOK, ad)
}

// Get retrieves an ad by ID.
// @Summary Get ad by ID
// @Description Retrieves an ad based on the provided ID
// @Tags ads
// @Accept json
// @Produce json
// @Param id query int true "Ad ID"
// @Success 200 {object} models.Ad
// @Failure 400 {string} string "Invalid parameter id"
// @Failure 500 {string} string "Could not retrieve ads"
// @Router /ads/{id} [get]
func (a AdsHandler) Get(c echo.Context) error {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid parameter id")
	}

	userRole := c.Get("account").(models.User).Role

	resp, err := a.datastore.Get(index, userRole)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not retrieve ads")
	}

	return c.JSON(http.StatusOK, resp)
}

// List retrieves all ads.
// @Summary List of ads.
// @Description Retrieves ads from database and accept query params.
// @Tags ads
// @Accept json
// @Produce json
// @Success 200 {object} []models.Ad
// @Failure 500 {string} string "Could not retrieve ads"
// @Router /ads [get]
func (a AdsHandler) List(c echo.Context) error {
	filter := filter.NewAdsFilter(c.QueryParams())

	filter.Base.UserRole = c.Get("account").(models.User).Role

	if len(filter.Base.Sort) != 0 {
		resp, err := a.datastore.ListFilterSort(&filter.Base)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "could not retrieve ads")
		}
		return c.JSON(http.StatusOK, resp)
	}

	resp, err := a.datastore.ListFilterByColumn(filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not retrieve ads")
	}
	return c.JSON(http.StatusOK, resp)
}
