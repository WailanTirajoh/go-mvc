package router

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/labstack/echo/v4"
)

func Setup(
	userController *controller.UserController,
	authController *controller.AuthContoroller,
) *echo.Echo {
	e := echo.New()
	e.GET("/users", userController.Index)
	e.GET("/users/:id", userController.Show)
	e.POST("/users", userController.Store)
	e.PUT("/users/:id", userController.Update)
	e.DELETE("/users/:id", userController.Destroy)

	e.POST("/login", authController.Login)
	e.POST("/logout", authController.Logout)

	return e
}
