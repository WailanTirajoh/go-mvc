package main

import (
	"log"
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"github.com/WailanTirajoh/go-simple-clean-architecture/router"
)

func main() {
	db := config.SetupConnection()

	// Setup Repository
	userRepository := repository.NewUserRepository(db)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)

	// Setup Router
	r := router.Setup(&userController)

	// Start App
	log.Fatal(http.ListenAndServe(":9000", r))
}
