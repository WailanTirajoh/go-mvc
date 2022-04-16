package router

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/middleware"
	"github.com/labstack/echo/v4"
)

func Setup(
	userController *controller.UserController,
	authController *controller.AuthContoroller,
) *echo.Echo {
	e := echo.New()

	auth := e.Group("/v1")
	{
		auth.Use(middleware.Authenticate)
		auth.GET("/users", userController.Index)
		auth.GET("/users/:id", userController.Show)
		auth.POST("/users", userController.Store)
		auth.PUT("/users/:id", userController.Update)
		auth.DELETE("/users/:id", userController.Destroy)
		auth.POST("/logout", authController.Logout)
	}

	guest := e.Group("/v1")
	{
		guest.POST("/login", authController.Login)
	}

	return e
}
