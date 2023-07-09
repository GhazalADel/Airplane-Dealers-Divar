package handlers

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
