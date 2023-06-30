package routes

import (
	"fuji-auth/pkg/handlers"

	_ "fuji-auth/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func AuthRoutes(e *echo.Echo) {
	auth := e.Group("/api/v1/user")
	auth.GET("/swagger/*", echoSwagger.WrapHandler)
	auth.GET("/status", handlers.Status)
	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.LoginUser)
	auth.GET("/logout", handlers.LogoutUser)

}
