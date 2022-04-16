package middleware

import (
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/labstack/echo/v4"
)

type (
	Auth interface {
		// Middleware
		Authenticate(next echo.HandlerFunc) echo.HandlerFunc

		ValidateWithToken(token string) (model.User, error)
		SetUser(user model.User)
		GetUser() model.User
	}

	AuthImpl struct {
		User           model.User
		AuthController *controller.AuthContoroller
	}
)

func NewAuth(ac *controller.AuthContoroller) Auth {
	return &AuthImpl{
		AuthController: ac,
	}
}

func (a *AuthImpl) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")

		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "token is required")
		}

		jwt := service.NewJWT()

		// Check if token is valid
		if err := jwt.ValidateToken(token); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		user, err := a.ValidateWithToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		a.SetUser(user)

		return next(c)
	}
}

func (a *AuthImpl) ValidateWithToken(token string) (model.User, error) {
	return a.AuthController.AuthService.ValidateUserToken(token)
}

func (a *AuthImpl) SetUser(user model.User) {
	a.User = user
}

func (a *AuthImpl) GetUser() model.User {
	return a.User
}
