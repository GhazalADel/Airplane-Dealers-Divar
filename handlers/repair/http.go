package handlers

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/datastore/repair"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type RepairHandler struct {
	RepairDatastore datastore.Repair
	UserDatastore   datastore.User
}

func NewRepairHandler(repairDS datastore.Repair, userDS datastore.User) *RepairHandler {
	return &RepairHandler{
		RepairDatastore: repairDS,
		UserDatastore:   userDS,
	}
}

// @Summary Request to repair check
// @Description Request to repair check
// @Tags repair
// @Accept json
// @Produce json
// @Param adID path int true "Ad ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /repair/ads/{adID}/check-request [post]
func (e *RepairHandler) RequestToRepairCheck(c echo.Context) error {
	ctx := c.Request().Context()
	adID, err := strconv.Atoi(c.Param("adID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	err = e.RepairDatastore.RequestToRepairCheck(ctx, adID, 2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	}

	resp := models.SuccessResponse{
		Success: true,
	}

	return c.JSON(201, resp)
}

// @Summary retrieve repair check request for repair or user
// @Description retrieve repair check request for repair or user
// @Tags repair
// @Param adID path int true "ad ID"
// @Success 200 {object} models.GetRepairRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /repair/ads/{adID} [get]
func (e *RepairHandler) GetRepairRequest(c echo.Context) error {
	ctx := c.Request().Context()
	user, _ := e.UserDatastore.Get(ctx, 2)
	adID, _ := strconv.Atoi(c.Param("adID"))

	repairRequest, err := e.RepairDatastore.GetByAd(ctx, adID, user)
	if err != nil {
		return c.JSON(
			http.StatusNotFound, models.ErrorResponse{Error: "repair request not found!"},
		)
	}

	resp := models.GetRepairRequestResponse{
		ID:        int(repairRequest.ID),
		UserID:    int(repairRequest.UserID),
		AdSubject: repairRequest.Ads.Subject,
		Status:    string(repairRequest.Status),
		CreatedAt: repairRequest.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary ListRepairRequest retrieves all repair requests for an repair
// @Description ListRepairRequest retrieves all repair requests for an repair
// @Tags repair
// @Param user_id query int false "User ID"
// @Param ads_id query int false "Ad ID"
// @Param from_date query string false "From date"
// @Success 200 {array} models.RepairRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /repair/check-requests [get]
func (e *RepairHandler) GetAllRepairRequest(c echo.Context) error {
	ctx := c.Request().Context()

	user, _ := e.UserDatastore.Get(ctx, 4)

	// params
	page, _ := strconv.Atoi(c.QueryParam("page"))
	fromDate := c.QueryParam("from_date")
	userID, _ := strconv.Atoi(c.QueryParam("user_id"))
	adsID, _ := strconv.Atoi(c.QueryParam("ads_id"))

	var (
		queryAndCondition  clause.AndConditions
		queryOrConditions  []clause.OrConditions
		filterAndCondition repair.FilterAndConditionRepairRequest
		filterNotCondtion  repair.FilterNotConditionRepairRequest
	)
	filterAndCondition = repair.FilterAndConditionRepairRequest{
		FromDate: fromDate,
		AdsID:    adsID,
		UserID:   userID,
	}
	if user.Role == 1 {
		filterNotCondtion = repair.FilterNotConditionRepairRequest{
			Status: utils.WAIT_FOR_PAYMENT_STATUS,
		}
	} else if user.Role == 4 {
		filterAndCondition.UserID = int(user.ID)
	}

	queryAndCondition, _ = filterAndCondition.ToQueryModel()
	queryNotCondition := filterNotCondtion.ToQueryModel()

	repairRequests, err := e.RepairDatastore.GetAllRepairRequests(
		ctx, queryAndCondition, queryOrConditions, queryNotCondition, page,
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})

	}

	var resp []models.RepairRequestResponse

	for _, v := range repairRequests {
		resp = append(resp, models.RepairRequestResponse{
			ID:        int(v.ID),
			UserID:    int(v.UserID),
			AdID:      int(v.AdsID),
			Status:    string(v.Status),
			CreatedAt: v.CreatedAt,
		})
	}

	return c.JSON(200, resp)
}

// @Summary Update repair request
// @Description Update repair request
// @Tags repair
// @Accept json
// @Produce json
// @Param repairRequestID path int true "repair request ID"
// @Param repairCheckRequest body models.UpdateRepairRequest true "repair object"
// @Success 200 {object} models.RepairRequestResponse
// @Failure 204 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /repair/check-request/{repairRequestID} [put]
func (e *RepairHandler) UpdateRepairRequest(c echo.Context) error {
	ctx := c.Request().Context()
	user, _ := e.UserDatastore.Get(ctx, 4)

	repairRequestID, _ := strconv.Atoi(c.Param("repairRequestID"))

	var updatedRepairRequest models.UpdateRepairRequest
	if err := c.Bind(&updatedRepairRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	repairRequest, err := e.RepairDatastore.Update(
		ctx, repairRequestID, user, updatedRepairRequest,
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp := models.RepairRequestResponse{
		ID:        int(repairRequest.ID),
		UserID:    int(repairRequest.UserID),
		AdID:      int(repairRequest.AdsID),
		Status:    string(repairRequest.Status),
		CreatedAt: repairRequest.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)
}
