package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/controller"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/request"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"github.com/WailanTirajoh/go-simple-clean-architecture/router"
	"github.com/labstack/echo/v4"
)

func NewApp() {
	app := config.NewAppConfig()
	db, err := config.NewConnection()

	if err != nil {
		panic(err)
	}

	// Setup Repository
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	validate := request.NewValidator()

	// Setup Service
	authService := service.NewAuthService(&userRepository, validate)
	userService := service.NewUserService(&userRepository, validate)
	roleService := service.NewRoleService(&roleRepository, validate)

	// Setup Controller
	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)
	roleController := controller.NewRoleController(&roleService)

	// Setup Router
	echo := echo.New()
	router := router.Setup(
		echo,
		&userController,
		&authController,
		&roleController,
	)

	// Start App
	fmt.Println("Running the server on port", app.APP_PORT)
	strPort := strconv.FormatInt(app.APP_PORT, 10)
	log.Fatal(http.ListenAndServe(":"+strPort, router))
}
