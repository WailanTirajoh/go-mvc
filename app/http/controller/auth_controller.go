package controller

import (
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/labstack/echo/v4"
)

type AuthContoroller struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthContoroller {
	return AuthContoroller{
		AuthService: *authService,
	}
}

func (authController *AuthContoroller) Login(c echo.Context) error {
	loginRequest := new(model.LoginRequest)
	if err := c.Bind(&loginRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := authController.AuthService.Login(loginRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"token": token,
	}))
}

func (authController *AuthContoroller) Logout(c echo.Context) error {
	var err error

	token := c.Request().Header.Get("token")

	if err = authController.AuthService.Logout(token); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"message": "Logout success",
	}))
}
