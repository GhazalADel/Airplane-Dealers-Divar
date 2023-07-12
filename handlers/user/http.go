package handlers

import (
	"Airplane-Divar/consts"
	"Airplane-Divar/datastore"
	"Airplane-Divar/models"
	"Airplane-Divar/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// define this structs for swagger docs
type UserResponse struct {
	ID       uint   `json:"ID"`
	UserID   uint   `json:"UserID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	IsActive bool   `json:"IsActive"`
}
type ErrorResponseRegisterLogin struct {
	ResponseCode int    `json:"responsecode"`
	Message      string `json:"message"`
}
type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type BudgetAmountResponse struct {
	Amount int `json:"amount"`
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
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponseRegisterLogin
// @Failure 422 {object} ErrorResponseRegisterLogin
// @Failure 500 {object} ErrorResponseRegisterLogin
// @Router /users/register [post]
func (a UserHandler) RegisterHandler(c echo.Context) error {
	// Read Request Body
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	//check json format
	jsonFormatValidationMsg, jsonFormatErr := utils.ValidateJsonFormat(jsonBody, "username", "password", "role")
	if jsonFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: jsonFormatValidationMsg})
	}

	//check unique User
	userUniqueMsg, userUniqueErr := a.data_user.CheckUnique(jsonBody["username"].(string))
	if userUniqueErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: userUniqueMsg})
	}

	//role
	role := string(consts.ROLE_AIRLINE)
	if r, ok := jsonBody["role"]; ok {
		if r.(string) != string(consts.ROLE_ADMIN) && r.(string) != (consts.ROLE_EXPERT) {
			return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid Role"})
		} else {
			role = r.(string)
		}
	}
	if role == string(consts.ROLE_EXPERT) || role == string(consts.ROLE_ADMIN) {
		if code, ok := jsonBody["code"]; !ok {
			return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "You have to enter code!"})
		} else {
			if role == string(consts.ROLE_EXPERT) {
				if code.(string) != "expert123" {
					return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "WRONG code"})
				}
			} else {
				if code.(string) != "admin123" {
					return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "WRONG code"})
				}
			}
		}
	}

	//create user
	userCreationMsg, user, userCreationErr := a.data_user.Create(jsonBody["username"].(string), jsonBody["password"].(string), role)
	if userCreationErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: userCreationMsg})
	}

	return c.JSON(http.StatusOK, user)
}

// LoginHandler handles user login
// @Summary User login
// @Description Login with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login request body"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponseRegisterLogin
// @Failure 422 {object} ErrorResponseRegisterLogin
// @Router  /users/login [post]
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

	//find user based on username and check password correction
	findUserMsg, user, findUserErr := a.data_user.Login(jsonBody["username"].(string), jsonBody["password"].(string))
	if findUserErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: findUserMsg})
	}

	return c.JSON(http.StatusOK, user)
}
