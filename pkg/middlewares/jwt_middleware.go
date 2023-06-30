package middlewares

import (
	"fuji-auth/pkg/constants"
	models "fuji-auth/pkg/models"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware() echo.MiddlewareFunc {

	jwtSecret := os.Getenv(constants.JwtSecret)
	jwtConfig := echojwt.Config{
		ErrorHandler: jwtError,
		SigningKey:   jwtSecret,
		ContextKey:   "jwt",
	}

	return echojwt.WithConfig(jwtConfig)
}

func jwtError(c echo.Context, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.JSON(http.StatusBadRequest, &models.ResponseModel[any]{
			Error:      true,
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Payload:    nil,
		})
	}

	// Return status 403 and failed authentication error.
	return c.JSON(http.StatusUnauthorized, &models.ResponseModel[any]{
		Error:      true,
		StatusCode: http.StatusUnauthorized,
		Message:    err.Error(),
		Payload:    nil,
	})
}
