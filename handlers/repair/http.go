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
