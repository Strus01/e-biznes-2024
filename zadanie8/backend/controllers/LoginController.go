package controllers

import (
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var users = map[string]*models.User{}

type LoginController struct{}

func (loginController *LoginController) CreateUser(context echo.Context) error {
	user := new(models.User)
	if err := context.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid data")
	}

	if _, exists := users[user.Username]; exists {
		return echo.NewHTTPError(http.StatusConflict, "Username already in use")
	}

	if err := user.HashPassword(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not hash password")
	}

	user.ID = uint64(len(users) + 1)
	users[user.Username] = user
	return context.JSON(http.StatusCreated, user)
}

func (loginController *LoginController) AuthenticateUser(context echo.Context) error {
	loginReq := new(models.User)
	if err := context.Bind(loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	user, ok := users[loginReq.Username]
	if !ok || !user.CheckPassword(loginReq.Password) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return context.JSON(http.StatusOK, user.ID)
}
