package handlers

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	logging_service "Airplane-Divar/service/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	merchantID      = "860C78FA-D6A9-48AE-805D-2B33B52309D2"
	callbackURL     = "http://localhost:8080/users/payment/verify"
	zarinpalRequest = "https://sandbox.banktest.ir/zarinpal/api.zarinpal.com/pg/v4/payment/request.json"
	zarinpalVerify  = "https://sandbox.banktest.ir/zarinpal/api.zarinpal.com/pg/v4/payment/verify.json"
	zarinpalGateURL = "https://sandbox.banktest.ir/zarinpal/www.zarinpal.com/pg/StartPay/"
)

type ZarinpalData struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Authority string `json:"authority"`
	FeeType   string `json:"fee_type"`
	Fee       int    `json:"fee"`
}

type ZarinpalResponse struct {
	Data   ZarinpalData  `json:"data"`
	Errors []interface{} `json:"errors"`
}

type ZarinpalVerify struct {
	Data   ZarinpalVerifyData `json:"data"`
	Errors ZarinpalErrors     `json:"errors"`
}

type ZarinpalVerifyData struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CardHash string `json:"card_hash"`
	CardPan  string `json:"card_pan"`
	RefID    int    `json:"ref_id"`
	FeeType  string `json:"fee_type"`
	Fee      int    `json:"fee"`
}

type ZarinpalErrors struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type PaymentRequest struct {
	AdID             int      `json:"adID" binding:"required"`
	TransactionTypes []string `json:"transactionType" binding:"required"`
}

type TransactioTypeObject struct {
	Type     string
	ObjectID uint
}

type VerifyResponse struct {
	Status    string `json:"Status"`
	Authority string `json:"Authority"`
}

type RequestResponse struct {
	PaymentUrl string `json:"payment_url"`
}

type PaymentHandler struct {
	UserDS    datastore.User
	ExpertDS  datastore.Expert
	RepairDS  datastore.Repair
	PaymentDS datastore.Payment
}

func NewPaymentHandler(
	userDS datastore.User,
	expertDS datastore.Expert,
	repairDS datastore.Repair,
	paymentDS datastore.Payment,
) *PaymentHandler {
	return &PaymentHandler{
		ExpertDS:  expertDS,
		RepairDS:  repairDS,
		UserDS:    userDS,
		PaymentDS: paymentDS,
	}
}

// @Summary Add budget request
// @Description Zarinpal Payment to add budget to user
// @Tags payment
// @Accept json
// @Produce json
// @Param Authorization header string true "User Token"
// @Param body body PaymentRequest true "Payment request details"
// @Success 200 {object} RequestResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/payment/request [post]
func (p PaymentHandler) PaymentRequestHandler(c echo.Context) error {
	user := c.Get("user").(models.User)
	ctx := c.Request().Context()

	if user.Role != consts.ROLE_AIRLINE {
		return c.JSON(
			http.StatusForbidden,
			models.ErrorResponse{
				Error: "Not allowed!",
			},
		)
	}

	var requestBody PaymentRequest
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// get service requests
	requestIds := []TransactioTypeObject{}
	for _, tType := range requestBody.TransactionTypes {
		if tType == models.ExpertAds.TableName(models.ExpertAds{}) {
			expertAd, err := p.ExpertDS.GetByAd(ctx, requestBody.AdID, user)
			if err != nil {
				return c.JSON(
					http.StatusNotFound, models.ErrorResponse{
						Error: "You don't have any expert request",
					},
				)
			}
			if expertAd.Status != consts.WAIT_FOR_PAYMENT_STATUS {
				return c.JSON(
					http.StatusNotFound, models.ErrorResponse{
						Error: "it's already payed!",
					},
				)
			}
			requestIds = append(requestIds, TransactioTypeObject{
				Type: tType, ObjectID: expertAd.ID,
			})

		} else if tType == models.RepairRequest.TableName(models.RepairRequest{}) {
			repairRequest, err := p.RepairDS.GetByAd(ctx, requestBody.AdID, user)
			if err != nil {
				return c.JSON(
					http.StatusNotFound, models.ErrorResponse{
						Error: "You don't have any repair request",
					},
				)
			}
			if repairRequest.Status != consts.WAIT_FOR_PAYMENT_STATUS {
				return c.JSON(
					http.StatusNotFound, models.ErrorResponse{
						Error: "it's already payed!",
					},
				)
			}
			requestIds = append(requestIds, TransactioTypeObject{
				Type: tType, ObjectID: repairRequest.ID,
			})

		} else {
			return c.JSON(
				http.StatusNotFound, models.ErrorResponse{
					Error: "service not found!",
				},
			)
		}
	}

	// calculate price
	prices, err := p.PaymentDS.GetPriceByServices(ctx, requestBody.TransactionTypes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}
	total_price := p.PaymentDS.GetTotalPriceByServices(prices)

	data := map[string]interface{}{
		"merchant_id":  merchantID,
		"amount":       total_price,
		"callback_url": callbackURL,
		"description":  "Pay for service",
		"metadata": map[string]string{
			"mobile": "09121234567",
			"email":  "test@test.com",
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			ResponseCode: 500,
			Message:      "Failed to marshal JSON data",
		})
	}

	resp, err := http.Post(zarinpalRequest, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return c.JSON(http.StatusBadGateway, models.Response{
			ResponseCode: 502,
			Message:      "Failed to send POST request",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			ResponseCode: 500,
			Message:      "Failed to read body",
		})
	}

	var result ZarinpalResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			ResponseCode: 500,
			Message:      "Failed to parse response",
		})
	}

	//create payment Transaction
	for _, service := range requestIds {
		transactionCreationMsg, transactionCreationErr := p.PaymentDS.Create(
			user.ID, int64(prices[service.Type]), result.Data.Authority,
			service.Type, service.ObjectID,
		)
		if transactionCreationErr != nil {
			return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: transactionCreationMsg})
		}
	}

	var response RequestResponse
	response.PaymentUrl = fmt.Sprintf("%s%s", zarinpalGateURL, result.Data.Authority)

	// ____ Report Log ____
	logService := logging_service.GetInstance()
	err = logService.ReportActivity(user.Role, user.ID, "Ads", uint(requestBody.AdID), consts.LOG_PAYMENT, fmt.Sprintf("Total Price: %v", total_price))
	if err != nil {
		_ = fmt.Errorf("cannot report activity log %v", consts.LOG_PAYMENT)
	}
	// ____ Report Log ____

	return c.JSON(http.StatusOK, response)
}

// Todo: Talk about handling payments

// @Summary Verify budget payment and add budget
// @Description Verify Zarinpal Payment to add budget to user
// @Tags payment
// @Accept json
// @Produce json
// @Param Authorization header string true "User Token"
// @Param body body VerifyResponse true "Payment verify details"
// @Success 200 {string} string
// @Failure 400 {string} models.ErrorResponse
// @Failure 422 {string} models.ErrorResponse
// @Failure 500 {string} models.ErrorResponse
// @Router /users/payment/verify [get]
func (p PaymentHandler) PaymentVerifyHandler(c echo.Context) error {
	ctx := c.Request().Context()
	authority := c.QueryParam("Authority")
	status := c.QueryParam("Status")

	transactions, err := p.PaymentDS.FindByAuthority(authority)
	if err != nil {
		// Handle the error (e.g., transaction not found)
		return c.JSON(http.StatusNotFound, models.Response{ResponseCode: 404, Message: "Transaction Not Founded"})
	}
	transactionsIDS := []uint{}
	var totoalAmount int64 = 0
	for _, t := range transactions {
		transactionsIDS = append(transactionsIDS, t.ID)
		totoalAmount += t.Amount
	}

	// ---- Report Log ----
	var log_status_temp int = 0
	defer func(logStatus int, transactionIds []uint) {
		logService := logging_service.GetInstance()
		if logService != (*logging_service.Logging)(nil) {
			if log_status_temp == -1 {
				for _, v := range transactionIds {
					err = logService.ReportActivity("", 0, "Transaction", v, consts.LOG_PAYMENT_FAILED, "")
					if err != nil {
						_ = fmt.Errorf("cannot log activity %v", consts.LOG_PAYMENT_FAILED)
					}
				}
			} else if log_status_temp == 1 {
				for _, v := range transactionIds {
					err = logService.ReportActivity("", 0, "Transaction", v, consts.LOG_PAYMENT_SUCCESS, "")
					if err != nil {
						_ = fmt.Errorf("cannot log activity %v", consts.LOG_PAYMENT_SUCCESS)
					}
				}
			}
		}
	}(log_status_temp, transactionsIDS)
	// ____ Report Log ----

	if status == "NOK" {
		p.PaymentDS.Update("Faield", transactionsIDS)
		log_status_temp = -1

		return c.JSON(http.StatusBadRequest, "Failed Payment")
	}

	data := map[string]interface{}{
		"merchant_id": merchantID,
		"amount":      totoalAmount,
		"authority":   transactions[0].Authority,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			ResponseCode: 500,
			Message:      "Failed to marshal JSON data",
		})
	}

	resp, err := http.Post(zarinpalVerify, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return c.JSON(http.StatusBadGateway, models.Response{
			ResponseCode: 502,
			Message:      "Failed to send POST request",
		})
	}
	defer resp.Body.Close()

	// Read Request Body
	jsonBody := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	if data, ok := jsonBody["data"]; ok {
		if dataMap, ok := data.(map[string]interface{}); ok {
			if code, ok := dataMap["code"]; ok {
				if code == float64(100) {
					p.PaymentDS.Update("Successful", transactionsIDS)

					// update service requests status
					for _, t := range transactions {
						if t.TransactionType == models.ExpertAds.TableName(models.ExpertAds{}) {
							err = p.ExpertDS.Update(
								ctx, int(t.ObjectID),
								map[string]interface{}{"status": consts.EXPERT_PENDING_STATUS},
							)
						} else if t.TransactionType == models.RepairRequest.TableName(models.RepairRequest{}) {
							err = p.RepairDS.Update(
								ctx, int(t.ObjectID),
								map[string]interface{}{"status": consts.MATIN_PENDING_STATUS},
							)
						}
						if err != nil {
							return c.JSON(http.StatusInternalServerError, models.Response{
								ResponseCode: http.StatusInternalServerError,
								Message:      "Failed to update status",
							})
						}
					}
					log_status_temp = 1
					return c.JSON(http.StatusOK, "Successful Payment")
				} else if code == float64(101) {
					return c.JSON(http.StatusOK, "Transaction had verified")
				}
			}
		}

	}

	p.PaymentDS.Update("Faield", transactionsIDS)
	log_status_temp = -1

	return c.JSON(http.StatusBadRequest, "Failed Payment")
}
