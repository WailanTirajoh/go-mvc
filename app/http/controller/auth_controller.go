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

func (a *AuthContoroller) Login(c echo.Context) error {
	loginRequest := new(model.LoginRequest)
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	token, err := a.AuthService.Login(loginRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"token": token,
	}))
}

func (a *AuthContoroller) Logout(c echo.Context) error {
	var err error
	logoutRequest := new(model.LogoutRequest)
	if err = c.Bind(&logoutRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if err = a.AuthService.Logout(logoutRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"message": "Logout berhasil",
	}))
}
