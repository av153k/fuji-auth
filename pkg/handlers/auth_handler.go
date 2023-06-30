package handlers

import (
	"fmt"
	"fuji-auth/pkg/models"
	"fuji-auth/pkg/services"
	"fuji-auth/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Status godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health
// @Produce json
// @Success 200 {object} models.ResponseModel[any]
// @Router /status [get]
func Status(c echo.Context) error {
	response := &models.ResponseModel[any]{
		Error:      false,
		Message:    "microservice fuji-auth is up and running!!",
		Payload:    nil,
		StatusCode: http.StatusOK,
	}

	log.Info(fmt.Sprintf("Microservice status: %#v", response))
	return c.JSON(http.StatusOK, response)
}


// RegisterUser method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param phone body string true "Phone"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} models.ResponseModel[models.User]
// @Router /register [post]
func RegisterUser(c echo.Context) error {
	registerData := &models.Register{}
	if err := c.Bind(registerData); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ResponseModel[any]{
			Error:      true,
			Message:    err.Error(),
			Payload:    nil,
			StatusCode: http.StatusBadRequest,
		})
	}
	dbService, err := services.NewDatabaseService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.ResponseModel[any]{
			Error:      true,
			Message:    err.Error(),
			Payload:    nil,
			StatusCode: http.StatusInternalServerError,
		})
	}

	role, err := utils.VerifyRole(registerData.UserRole)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ResponseModel[any]{
			Error:      true,
			Message:    err.Error(),
			Payload:    nil,
			StatusCode: http.StatusBadRequest,
		})
	}

	hashedPassword, err := utils.GeneratePasswordHash(registerData.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.ResponseModel[any]{
			Error:      true,
			Message:    err.Error(),
			Payload:    nil,
			StatusCode: http.StatusInternalServerError,
		})
	}

	user := &models.User{}

	user.ID = uuid.New()
	user.Email = registerData.Email
	user.Phone = registerData.Phone
	user.Password = hashedPassword
	user.AccountStatus = true
	user.IsOnline = true
	user.UserRole = role
	user.Name = ""
	user.PhotoUrl = ""

	if err := dbService.CreateUser(user); err != nil {
		// Return status 500 and create user process error.
		return c.JSON(http.StatusInternalServerError, &models.ResponseModel[any]{
			Error:      true,
			Message:    err.Error(),
			Payload:    nil,
			StatusCode: http.StatusInternalServerError,
		})
	}

	// Delete password field from JSON view.
	user.Password = ""

	// Return status 200 OK.
	return c.JSON(http.StatusOK, &models.ResponseModel[*models.User]{
		Error:      true,
		Message:    err.Error(),
		Payload:    user,
		StatusCode: http.StatusOK,
	})
}

// LoginUser method to login a user and send tokens in response.
// @Description login a user and send tokens in response.
// @Summary login a user and send tokens in response
// @Tags User
// @Accept json
// @Produce json
// @Param email body string false "Email"
// @Param phone body string false "Phone"
// @Param password body string true "Password"
// @Success 200 {object} models.ResponseModel[models.User]
// @Router /login [post]
func LoginUser(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.ResponseModel[any]{})
}

func LogoutUser(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.ResponseModel[any]{})
}
