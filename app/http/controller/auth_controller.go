package controller

import (
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"github.com/labstack/echo/v4"
)

type AuthContoroller struct {
	AuthService service.AuthService
}

// To initialize auth controller
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

	user, err := a.AuthService.Login(loginRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	jwt := service.NewJWT(map[string]interface{}{
		"id":  user.ID,
		"key": user.Password,
	})

	jwt.SetSecret(config.GetEnv("APP_KEY", "mysecretpassword")).
		GenerateToken()

	token := jwt.GetToken()

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"token": token,
	}))
}
