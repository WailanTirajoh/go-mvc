package router

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/middleware"
	"github.com/labstack/echo/v4"
)

func Setup(
	e *echo.Echo,
	userController *controller.UserController,
	authController *controller.AuthContoroller,
	roleController *controller.RoleController,
) *echo.Echo {
	authRoute := e.Group("/v1")
	{
		auth := middleware.NewAuth(authController)
		authRoute.Use(auth.Authenticate)

		authRoute.GET("/users", userController.Index)
		authRoute.GET("/users/:id", userController.Show)
		authRoute.POST("/users", userController.Store)
		authRoute.PUT("/users/:id", userController.Update)
		authRoute.DELETE("/users/:id", userController.Destroy)

		authRoute.GET("/roles", roleController.Index)
		authRoute.GET("/roles/:id", roleController.Show)
		authRoute.POST("/roles", roleController.Store)
		authRoute.PUT("/roles/:id", roleController.Update)
		authRoute.DELETE("/roles/:id", roleController.Destroy)

		authRoute.POST("/logout", authController.Logout)
	}

	guestRoute := e.Group("/v1")
	{
		guestRoute.POST("/login", authController.Login)
		guestRoute.POST("/register", authController.Register)
	}

	return e
}
