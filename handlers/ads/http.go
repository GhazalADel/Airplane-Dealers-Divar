package ads

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/datastore"
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	logging_service "Airplane-Divar/service/logging"
	"Airplane-Divar/utils"
	"encoding/json"
	"fmt"
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

type AdRequest struct {
	Image         string `json:"Image"`
	Description   string `json:"Description"`
	Subject       string `json:"Subject"`
	Price         uint64 `json:"Price"`
	Category      string `json:"Category"`
	FlyTime       uint   `json:"FlyTime"`
	AirplaneModel string `json:"AirplaneModel"`
	RepairCheck   bool   `json:"RepairCheck"`
	ExpertCheck   bool   `json:"ExpertCheck"`
	PlaneAge      uint   `json:"PlaneAge"`
}

type AdResponse struct {
	ID            uint   `json:"ID"`
	UserID        uint   `json:"UserID"`
	Image         string `json:"Image"`
	Description   string `json:"Description"`
	Subject       string `json:"Subject"`
	Price         uint64 `json:"Price"`
	CategoryID    uint   `json:"CategoryID"`
	Status        string `json:"Status"`
	FlyTime       uint   `json:"FlyTime"`
	AirplaneModel string `json:"AirplaneModel"`
	RepairCheck   bool   `json:"RepairCheck"`
	ExpertCheck   bool   `json:"ExpertCheck"`
	PlaneAge      uint   `json:"PlaneAge"`
}
type ErrorAddAd struct {
	ResponseCode int    `json:"responsecode"`
	Message      string `json:"message"`
}

// Create a new ad by an airline.
// @Summary Create an ad
// @Description Create new ad by given properties
// @Tags Ads
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Param AdRequest body AdRequest true "Ad details"
// @Success 200 {object} AdResponse
// @Failure 422 {object} ErrorAddAd
// @Failure 403 {object} ErrorAddAd
// @Failure 500 {object} ErrorAddAd
// @Router /ads/add [post]
func (a AdsHandler) AddAdHandler(c echo.Context) error {
	// Read Request Body
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	//check json format
	jsonFormatValidationMsg, jsonFormatErr := utils.ValidateJsonFormat(jsonBody, "Price", "Category", "FlyTime", "AirplaneModel", "RepairCheck", "ExpertCheck", "PlaneAge")
	if jsonFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: jsonFormatValidationMsg})
	}

	user := c.Get("user").(models.User)
	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(http.StatusForbidden, models.Response{ResponseCode: 403, Message: "Airlines Can Add an ad!"})
	}

	//validate and initialize categoryID in ad object
	category_name := ""
	if cat, ok := jsonBody["Category"].(string); ok {
		category_name = cat
	} else {
		msg := "Category should be string !"
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: msg})
	}
	categoryObj, err := a.datastore.GetCategoryByName(category_name)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid Category Name"})
	}

	var ad models.Ad

	//check ad properties validation
	adFormatValidationMsg, ad, adFormatErr := utils.ValidateAd(jsonBody, categoryObj)
	if adFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: adFormatValidationMsg})
	}

	//set user id

	id := user.ID
	ad.UserID = id

	ad.Status = string(consts.INACTIVE)

	//Create Ad
	createdAd, err := a.datastore.CreateAd(&ad)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: "Ad Cration Failed"})
	}
	adRes := models.AdResponse{
		ID:            createdAd.ID,
		UserID:        createdAd.UserID,
		Image:         createdAd.Image,
		Description:   createdAd.Description,
		Subject:       createdAd.Subject,
		Price:         createdAd.Price,
		CategoryID:    createdAd.CategoryID,
		Status:        createdAd.Status,
		FlyTime:       createdAd.FlyTime,
		AirplaneModel: createdAd.AirplaneModel,
		RepairCheck:   createdAd.RepairCheck,
		ExpertCheck:   createdAd.ExpertCheck,
		PlaneAge:      createdAd.PlaneAge,
	}

	// ____ Report Log ____
	logService := logging_service.GetInstance()
	if logService != (*logging_service.Logging)(nil) {
		err = logService.ReportActivity(user.Role, user.ID, "Ads", createdAd.ID, consts.LOG_CREATE_AD, "")
		if err != nil {
			_ = fmt.Errorf("cannot log activity %v", consts.LOG_CREATE_AD)
		}
	}
	// ____ Report Log ____

	return c.JSON(http.StatusOK, adRes)
}

// Status updates the status of an ad.
//
// This endpoint is used to update the status of an ad based on the provided ad ID.
//
// @Summary Update ad status
// @Description Update the status of an ad
// @Tags Ads
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Param id path integer true "Ad ID"
// @Param status body models.UpdateAdsStatusRequest true "status object"
// @Success 200 {string} string "Updated successfully"
// @Failure 400 {string} string "Invalid parameter id"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Could not update ads status"
// @Router /ads/{id}/status [put]
func (a AdsHandler) Status(c echo.Context) error {
	userRole := c.Get("user").(models.User).Role
	if userRole == "Admin" {
		id := c.Param("id")
		index, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid parameter id")
		}

		var status models.UpdateAdsStatusRequest
		if err := c.Bind(&status); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		_, err = a.datastore.UpdateStatus(index, status.Status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "could not update ads status")
		}

		// ____ Report Log ____
		logService := logging_service.GetInstance()
		logName := consts.LOG_ADMIN_APPROVE
		if status.Status == consts.INACTIVE {
			logName = consts.LOG_ADMIN_REJECT
		}
		if logService != (*logging_service.Logging)(nil) {
			err = logService.ReportActivity(userRole, c.Get("user").(models.User).ID, "Ads", uint(index), logName, "")
			if err != nil {
				_ = fmt.Errorf("cannot log activity %v", logName)
			}
		}
		// ____ Report Log ____

		return c.JSON(http.StatusOK, "Updated successfuly")
	} else {
		return c.JSON(http.StatusNotFound, "Not Found")
	}
}

// Get retrieves an ad by ID.
// @Summary Get ad by ID
// @Description Retrieves an ad based on the provided ID
// @Tags Ads
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Param id path int true "Ad ID"
// @Success 200 {object} models.Ad
// @Failure 400 {string} string "Invalid parameter id"
// @Failure 500 {string} string "Could not retrieve ads"
// @Router /ads/{id} [get]
func (a AdsHandler) Get(c echo.Context) error {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	fmt.Println("--THIS--", err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid parameter id")
	}

	userRole := c.Get("user").(models.User).Role

	resp, err := a.datastore.Get(index, userRole)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "could not retrieve ads")
	}

	// ____ Get Logs of Ads ____
	logService := logging_service.GetInstance()
	if logService != (*logging_service.Logging)(nil) {
		adsLogs, err := logService.GetAdsActivity(index)
		if err != nil {
			_ = fmt.Errorf("could not retrieve ads activity: %v", err)
		}
		fmt.Printf("---- Ads Logs ---- \n Ads %v Activity Logs: \n %v \n ---- Ads Logs ---- \n", index, string(adsLogs))
	}
	// ____ Get Logs of Ads ____

	return c.JSON(http.StatusOK, resp)
}

// ListAds retrieves a list of ads.
// @Summary List ads
// @Description Retrieves ads from the database and accepts query parameters for filtering and sorting.
// @Tags Ads
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "User Token"
// @Param filter query filter.AdsFilter true "Query parameters for filtering ads"
// @Success 200 {object} []models.Ad "Successfully retrieved ads"
// @Failure 500 {string} string "Internal Server Error: Failed to retrieve ads"
// @Router /ads [get]
func (a AdsHandler) List(c echo.Context) error {
	filter := filter.NewAdsFilter(c.QueryParams())

	filter.Base.UserRole = c.Get("user").(models.User).Role

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
