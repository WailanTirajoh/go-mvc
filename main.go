package main

import (
	"log"
	"net/http"

	"github.com/wailantirajoh/gorilla/app/http/controller"
	"github.com/wailantirajoh/gorilla/app/http/repository"
	"github.com/wailantirajoh/gorilla/app/http/service"
	"github.com/wailantirajoh/gorilla/config"
	"github.com/wailantirajoh/gorilla/router"
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
	r := router.Setup()

	userController.Route(r)

	// Start App
	log.Fatal(http.ListenAndServe(":9000", r))
}
