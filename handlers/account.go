package handlers

import (
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// define this structs for swagger docs
type AccountResponse struct {
	ID       uint   `json:"ID"`
	UserID   uint   `json:"UserID"`
	Username string `json:"Username"`
	Budget   int    `json:"Budget"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	IsActive bool   `json:"IsActive"`
}
type ErrorResponseRegisterLogin struct {
	ResponseCode int    `json:"responsecode"`
	Message      string `json:"message"`
}
type UserCreateRequest struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	NationalID string `json:"nationalid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type BudgetAmountResponse struct {
	Amount int `json:"amount"`
}

type SenderNumbersResponse struct {
	Numbers []string `json:"numbers"`
}

type AccountHandler struct {
	data_user    datastore.User
	data_account datastore.Account
}

func NewAccountHandler(user datastore.User, account datastore.Account) *AccountHandler {
	return &AccountHandler{data_user: user, data_account: account}
}

// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param body body UserCreateRequest true "User registration details"
// @Success 201 {object} AccountResponse
// @Failure 400 {object} ErrorResponseRegisterLogin
// @Failure 422 {object} ErrorResponseRegisterLogin
// @Failure 500 {object} ErrorResponseRegisterLogin
// @Router /accounts/register [post]
func (a AccountHandler) RegisterHandler(c echo.Context) error {
	// Read Request Body
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	//check json format
	jsonFormatValidationMsg, jsonFormatErr := utils.ValidateJsonFormat(jsonBody, "username", "password", "email", "phone", "nationalid", "username", "password")
	if jsonFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: jsonFormatValidationMsg})
	}

	//check user validation
	userFormatValidationMsg, user, userFormatErr := utils.ValidateUser(jsonBody)
	if userFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: userFormatValidationMsg})
	}

	//check unique User
	userUniqueMsg, userUniqueErr := a.data_user.CheckUnique(user)
	if userUniqueErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: userUniqueMsg})
	}

	//check unique Account
	accountUniqueMsg, accountUniqueErr := a.data_account.CheckUnique(jsonBody["username"].(string))
	if accountUniqueErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: accountUniqueMsg})
	}

	// Insert User Object Into Database
	_, err = a.data_user.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{ResponseCode: 500, Message: "User Cration Failed"})
	}

	//create account
	accountCreationMsg, account, accountCreationErr := a.data_account.Create(int(user.ID), jsonBody["username"].(string), false, jsonBody["password"].(string))
	if accountCreationErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: accountCreationMsg})
	}

	return c.JSON(http.StatusOK, account)
}

// LoginHandler handles user login
// @Summary User login
// @Description Login with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login request body"
// @Success 200 {object} AccountResponse
// @Failure 400 {object} ErrorResponseRegisterLogin
// @Failure 422 {object} ErrorResponseRegisterLogin
// @Router  /accounts/login [post]
func (a AccountHandler) LoginHandler(c echo.Context) error {
	// Read Request Body
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	//check json format
	jsonFormatValidationMsg, jsonFormatErr := utils.ValidateJsonFormat(jsonBody, "username", "password")
	if jsonFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: jsonFormatValidationMsg})
	}

	//find account based on username and check password correction
	findAccountMsg, account, findAccountErr := a.data_account.Login(jsonBody["username"].(string), jsonBody["password"].(string), false)
	if findAccountErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: findAccountMsg})
	}

	return c.JSON(http.StatusOK, account)
}
