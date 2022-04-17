package bootstrap

import (
	"log"
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"github.com/WailanTirajoh/go-simple-clean-architecture/router"
)

func NewApp() {
	db := config.NewConnection()

	// Setup Repository
	userRepository := repository.NewUserRepository(db)

	// Setup Service
	authService := service.NewAuthService(&userRepository)
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)

	// Setup Router
	r := router.Setup(&userController, &authController)

	// Start App
	log.Fatal(http.ListenAndServe(":9000", r))
}
