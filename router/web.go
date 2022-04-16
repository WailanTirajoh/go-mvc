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

	authRoute := e.Group("/v1")
	{
		auth := middleware.NewAuth(authController)
		authRoute.Use(auth.Authenticate)

		authRoute.GET("/users", userController.Index)
		authRoute.GET("/users/:id", userController.Show)
		authRoute.POST("/users", userController.Store)
		authRoute.PUT("/users/:id", userController.Update)
		authRoute.DELETE("/users/:id", userController.Destroy)

		authRoute.POST("/logout", authController.Logout)
	}

	guestRoute := e.Group("/v1")
	{
		guestRoute.POST("/login", authController.Login)
	}

	return e
}
