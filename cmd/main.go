package main

import (
	"fmt"
	"fuji-auth/pkg/routes"
	"fuji-auth/pkg/utils"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

// @title Fuji Auth Microservice
// @version 1.0
// @description This is the auth microservice for Fuji.
// @termsOfService http://swagger.io/terms/

// @contact.name Abhishek Anand
// @contact.url http://www.github.com/av153k
// @contact.email av153k.dev@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1/user/
// @schemes http
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitLogger()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogError:     true,
		LogRequestID: true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			if values.Error != nil {
				log.WithFields(log.Fields{
					"URI":        values.URI,
					"status":     values.Status,
					"error":      values.Error,
					"formValues": values.FormValues,
					"headers":    values.Headers, "contentLength": values.ContentLength, "latency": values.Latency.Milliseconds(),
				}).Error("request")
			} else {
				log.WithFields(log.Fields{
					"URI":        values.URI,
					"status":     values.Status,
					"formValues": values.FormValues,
					"headers":    values.Headers, "contentLength": values.ContentLength, "latency": values.Latency.Milliseconds(),
				}).Info("request")
			}

			return nil
		},
	}))
	// e.Use(middlewares.JwtMiddleware())
	routes.AuthRoutes(e)
	serverConnectionUrl := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)
	log.Info("Starting server at: %v", serverConnectionUrl)
	// Start server (with or without graceful shutdown).
	if os.Getenv("FLAVOR") == "dev" {
		utils.StartServer(e, serverConnectionUrl)
	} else {
		utils.StartServerWithGracefulShutdown(e, serverConnectionUrl)
	}

}
