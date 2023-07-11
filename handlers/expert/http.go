package handlers

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/datastore"
	"Airplane-Divar/datastore/expert"
	"Airplane-Divar/models"
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
	user := c.Get("user").(models.User)

	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "Only airline user can send request for expert check",
			},
		)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	err = e.ExpertDatastore.RequestToExpertCheck(ctx, adID, user)
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
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/check-requests [get]
func (e *ExpertHandler) GetAllExpertRequest(c echo.Context) error {
	ctx := c.Request().Context()

	user := c.Get("user").(models.User)

	if user.Role == consts.ROLE_MATIN {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "You don't have access to expert requests.",
			},
		)
	}

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
		filterOrCondition  expert.FilterOrConditionExpertRequest
	)
	filterAndCondition = expert.FilterAndConditionExpertRequest{
		FromDate: fromDate,
		AdsID:    adsID,
		UserID:   userID,
	}
	if user.Role == consts.ROLE_EXPERT {
		filterOrCondition = expert.FilterOrConditionExpertRequest{
			ExpertIDList: []interface{}{user.ID, nil},
		}

		filterNotCondtion = expert.FilterNotConditionExpertRequest{
			Status: consts.WAIT_FOR_PAYMENT_STATUS,
		}
	} else if user.Role == consts.ROLE_AIRLINE {
		filterAndCondition.UserID = int(user.ID)
	}

	queryAndCondition, _ = filterAndCondition.ToQueryModel()
	queryNotCondition := filterNotCondtion.ToQueryModel()
	queryOrConditions = append(queryOrConditions, filterOrCondition.ToQueryModel())

	expertAds, err := e.ExpertDatastore.GetAllExpertRequests(
		ctx, queryAndCondition, queryOrConditions, queryNotCondition, page,
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	if len(expertAds) == 0 {
		return c.JSON(
			http.StatusOK,
			models.MessageResponse{Message: "There is no expert request."},
		)
	}

	var resp []models.ExpertRequestResponse

	for _, v := range expertAds {
		resp = append(resp, models.ExpertRequestResponse{
			ID:        int(v.ID),
			UserID:    int(v.UserID),
			AdID:      int(v.AdsID),
			ExpertID:  int(v.ExpertID),
			Status:    string(v.Status),
			Report:    v.Report,
			CreatedAt: v.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, resp)
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
	user := c.Get("user").(models.User)

	if user.Role != consts.ROLE_EXPERT {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "You can't update expert requests.",
			},
		)
	}

	expertRequestID, _ := strconv.Atoi(c.Param("expertRequestID"))

	var updatedExpertCheck models.UpdateExpertCheckRequest
	if err := c.Bind(&updatedExpertCheck); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	expertAd, err := e.ExpertDatastore.Update(ctx, expertRequestID, user, updatedExpertCheck)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	} else if expertAd.ID == 0 {
		return c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "expert check request does not exist!",
		})
	}

	resp := models.ExpertRequestResponse{
		ID:        int(expertAd.ID),
		UserID:    int(expertAd.UserID),
		AdID:      int(expertAd.AdsID),
		ExpertID:  int(expertAd.ExpertID),
		Status:    string(expertAd.Status),
		Report:    expertAd.Report,
		CreatedAt: expertAd.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary retrieve expert check request by ad for expert or user
// @Description retrieve expert check request by ad for expert or user
// @Tags expert
// @Param adID path int true "ad ID"
// @Success 200 {object} models.GetExpertRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/ads/{adID} [get]
func (e *ExpertHandler) GetExpertRequestByAd(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(models.User)
	adID, _ := strconv.Atoi(c.Param("adID"))

	if user.Role == consts.ROLE_MATIN {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "You don't have access to expert requests.",
			},
		)
	}

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
		Report:    expertAd.Report,
		CreatedAt: expertAd.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)

}

// @Summary retrieve expert check request for expert or user
// @Description retrieve expert check request for expert or user
// @Tags expert
// @Param requestID path int true "request ID"
// @Success 200 {object} models.GetExpertRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /expert/check-request/{requestID} [get]
func (e *ExpertHandler) GetExpertRequest(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(models.User)
	requestID, _ := strconv.Atoi(c.Param("requestID"))

	if user.Role == consts.ROLE_MATIN {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "You don't have access to expert requests.",
			},
		)
	}

	expertAd, err := e.ExpertDatastore.Get(ctx, requestID, user)
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
		Report:    expertAd.Report,
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
	user := c.Get("user").(models.User)
	adID, _ := strconv.Atoi(c.Param("adID"))

	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "You can't delete this expert request.",
			},
		)
	}

	err := e.ExpertDatastore.Delete(ctx, adID, user)
	if err != nil {
		return c.JSON(
			http.StatusNotFound, models.ErrorResponse{Error: "Expert request not found!"},
		)
	}

	return c.JSON(http.StatusNoContent, models.SuccessResponse{Success: true})

}
