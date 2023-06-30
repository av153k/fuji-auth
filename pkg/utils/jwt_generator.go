package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"fuji-auth/pkg/constants"
	"os"
	"strconv"
	"strings"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/golang-jwt/jwt"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func GetTokens(userID string) (*Tokens, error) {
	log.Info("Generating access token and refresh token for userID: %v", userID)
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(userID)
	if err != nil {
		log.Error("Failed to generate access token - Reason: %v", err)
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		log.Error("Failed to generate refresh token - Reason: %v", err)
		// Return token generation error.
		return nil, err
	}

	log.Debug("Token generated for UserID: %v \nAccess Token: %v \nRefresh Token - %v", userID, accessToken, refreshToken)
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateNewAccessToken(id string) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv(constants.JwtSecret)

	// Set token validity minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv(constants.JwtValidityInMinutes))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken() (string, error) {
	// Create a new SHA256 hash.
	hash := sha256.New()

	// Create a new now date and time string with salt.
	refresh := os.Getenv(constants.JwtRefreshKey) + time.Now().String()

	// See: https://pkg.go.dev/io#Writer.Write
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		// Return error, it refresh token generation failed.
		return "", err
	}

	// Set expires hours count for refresh key from .env file.
	hoursCount, _ := strconv.Atoi(os.Getenv(constants.JwtRefreshValidityInHours))

	// Set expiration time.
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time).
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
