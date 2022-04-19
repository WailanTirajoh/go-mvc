package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"github.com/WailanTirajoh/go-simple-clean-architecture/router"
)

func NewApp() {
	app := config.NewAppConfig()
	db, err := config.NewConnection()
	if err != nil {
		panic(err)
	}

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
	fmt.Println("Running the server on port", app.APP_PORT)
	strPort := strconv.FormatInt(app.APP_PORT, 10)
	log.Fatal(http.ListenAndServe(":"+strPort, r))
}
