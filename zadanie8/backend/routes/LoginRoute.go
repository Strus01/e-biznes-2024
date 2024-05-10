package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func LoginRoute(e *echo.Echo) {
	loginController := controllers.LoginController{}

	e.POST("/register", loginController.CreateUser)
	e.POST("/login", loginController.AuthenticateUser)
}
