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

type UserHandler struct {
	data_user datastore.User
}

func NewUserHandler(user datastore.User) *UserHandler {
	return &UserHandler{data_user: user}
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
func (a UserHandler) RegisterHandler(c echo.Context) error {
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

	//check unique Account
	accountUniqueMsg, accountUniqueErr := a.data_user.CheckUnique(jsonBody["username"].(string))
	if accountUniqueErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: accountUniqueMsg})
	}

	//create account
	accountCreationMsg, account, accountCreationErr := a.data_user.Create(jsonBody["username"].(string), jsonBody["password"].(string))
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
func (a UserHandler) LoginHandler(c echo.Context) error {
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
	findAccountMsg, account, findAccountErr := a.data_user.Login(jsonBody["username"].(string), jsonBody["password"].(string))
	if findAccountErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: findAccountMsg})
	}

	return c.JSON(http.StatusOK, account)
}
