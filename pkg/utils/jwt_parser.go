package utils

import (
	"fmt"
	"fuji-auth/pkg/constants"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type TokenMetadata struct {
	UserID  string
	Expires int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c echo.Context) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		log.Error("Failed to verify token - Reason: %v", err)
		return nil, err
	}
	log.Debug(fmt.Sprintf("Token metadata extracted - token: %#v", token))
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID := claims["id"].(string)

		// Expires time.
		expires := int64(claims["expires"].(float64))

		return &TokenMetadata{
			UserID:  userID,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func extractToken(c echo.Context) string {
	bearToken := c.Request().Header.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c echo.Context) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv(constants.JwtSecret)), nil
}
