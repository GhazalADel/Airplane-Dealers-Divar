package handlers

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/datastore/expert"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type ExpertHandler struct {
	ExpertDatastore datastore.Expert
	UserDatastore   datastore.User
}

func NewExpertHandler(expertDS datastore.Expert, userDS datastore.User) *ExpertHandler {
	return &ExpertHandler{
		ExpertDatastore: expertDS,
		UserDatastore:   userDS,
	}
}

// @Summary Request to expert check
// @Description Request to expert check
// @Tags expert
// @Accept json
// @Produce json
// @Param adID path int true "Ad ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/ads/{adID}/check-request [post]
func (e *ExpertHandler) RequestToExpertCheck(c echo.Context) error {
	ctx := c.Request().Context()
	adID, err := strconv.Atoi(c.Param("adID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	err = e.ExpertDatastore.RequestToExpertCheck(ctx, adID, 2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})

	}

	resp := models.SuccessResponse{
		Success: true,
	}

	return c.JSON(201, resp)
}

// @Summary ListExpertRequest retrieves all expert requests for an expert
// @Description ListExpertRequest retrieves all expert requests for an expert
// @Tags expert
// @Param user_id query int false "User ID"
// @Param ads_id query int false "Ad ID"
// @Param from_date query string false "From date"
// @Success 200 {array} models.ExpertRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/check-requests [get]
func (e *ExpertHandler) GetAllExpertRequest(c echo.Context) error {
	ctx := c.Request().Context()

	user, _ := e.UserDatastore.Get(ctx, 2)

	// params
	page, _ := strconv.Atoi(c.QueryParam("page"))
	fromDate := c.QueryParam("from_date")
	userID, _ := strconv.Atoi(c.QueryParam("user_id"))
	adsID, _ := strconv.Atoi(c.QueryParam("ads_id"))

	var (
		queryAndCondition  clause.AndConditions
		queryOrConditions  []clause.OrConditions
		filterAndCondition expert.FilterAndConditionExpertRequest
		filterNotCondtion  expert.FilterNotConditionExpertRequest
	)
	filterAndCondition = expert.FilterAndConditionExpertRequest{
		FromDate: fromDate,
		AdsID:    adsID,
	}
	if user.Role == 2 {
		filterAndCondition.UserID = userID

		filterNotCondtion = expert.FilterNotConditionExpertRequest{
			Status: utils.EXPERT_WAIT_FOR_PAYMENT_STATUS,
		}
	} else if user.Role == 4 {
		filterAndCondition.UserID = int(user.ID)
	}

	queryAndCondition, _ = filterAndCondition.ToQueryModel()
	queryNotCondition := filterNotCondtion.ToQueryModel()

	expertAds, err := e.ExpertDatastore.GetAllExpertRequests(
		ctx, queryAndCondition, queryOrConditions, queryNotCondition, page,
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})

	}

	var resp []models.ExpertRequestResponse

	for _, v := range expertAds {
		resp = append(resp, models.ExpertRequestResponse{
			ID:        int(v.ID),
			UserID:    int(v.UserID),
			AdID:      int(v.AdsID),
			ExpertID:  int(v.ExpertID),
			Status:    string(v.Status),
			CreatedAt: v.CreatedAt,
		})
	}

	return c.JSON(200, resp)
}

// @Summary Update expert check request
// @Description Update expert check request
// @Tags expert
// @Accept json
// @Produce json
// @Param expertRequestID path int true "expert request ID"
// @Param expertCheckRequest body models.UpdateExpertCheckRequest true "Expert check object"
// @Success 200 {object} models.ExpertRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/check-request/{expertRequestID} [put]
func (e *ExpertHandler) UpdateCheckExpert(c echo.Context) error {
	ctx := c.Request().Context()
	user, _ := e.UserDatastore.Get(ctx, 2)

	expertRequestID, _ := strconv.Atoi(c.Param("expertRequestID"))

	var updatedExpertCheck models.UpdateExpertCheckRequest
	if err := c.Bind(&updatedExpertCheck); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	expertAd, err := e.ExpertDatastore.Update(ctx, expertRequestID, user, updatedExpertCheck)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := models.ExpertRequestResponse{
		ID:        int(expertAd.ID),
		UserID:    int(expertAd.UserID),
		AdID:      int(expertAd.AdsID),
		ExpertID:  int(expertAd.ExpertID),
		Status:    string(expertAd.Status),
		CreatedAt: expertAd.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary retrieve expert check request for expert or user
// @Description retrieve expert check request for expert or user
// @Tags expert
// @Param adID path int true "ad ID"
// @Success 200 {object} models.GetExpertRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/ads/{adID} [get]
func (e *ExpertHandler) GetExpertRequest(c echo.Context) error {
	ctx := c.Request().Context()
	user, _ := e.UserDatastore.Get(ctx, 2)
	adID, _ := strconv.Atoi(c.Param("adID"))

	expertAd, err := e.ExpertDatastore.GetByAd(ctx, adID, user)
	if err != nil {
		return c.JSON(
			http.StatusNotFound, models.ErrorResponse{Error: "Expert request not found!"},
		)
	}

	resp := models.GetExpertRequestResponse{
		ID:        int(expertAd.ID),
		UserID:    int(expertAd.UserID),
		AdSubject: expertAd.Ads.Subject,
		ExpertID:  int(expertAd.ExpertID),
		Status:    string(expertAd.Status),
		CreatedAt: expertAd.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)

}

// @Summary delete expert check request for expert or user
// @Description delete expert check request for expert or user
// @Tags expert
// @Param adID path int true "ad ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/ads/{adID} [delete]
func (e *ExpertHandler) DeleteExpertRequest(c echo.Context) error {
	ctx := c.Request().Context()
	user, _ := e.UserDatastore.Get(ctx, 2)
	adID, _ := strconv.Atoi(c.Param("adID"))

	err := e.ExpertDatastore.Delete(ctx, adID, user)
	if err != nil {
		return c.JSON(
			http.StatusNotFound, models.ErrorResponse{Error: "Expert request not found!"},
		)
	}

	return c.JSON(http.StatusNoContent, models.SuccessResponse{Success: true})

}
