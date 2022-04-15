package controller

import (
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

type UserController struct {
	UserService service.UserService
}

func (userController *UserController) Index(c echo.Context) (err error) {
	users := userController.UserService.GetUsers()

	var usersResponse []model.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, user.Response())
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(usersResponse))
}

func (userController *UserController) Show(c echo.Context) (err error) {
	user, err := userController.UserService.GetUser(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(user))
}

func (userController *UserController) Store(c echo.Context) (err error) {
	var user model.User

	userRequest := new(model.StoreUserRequest)
	if err = c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	validate := validator.New()
	if err = validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse(validationErrors.Error()))
	}

	user, err = userController.UserService.StoreUser(userRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(user.Response()))
}

func (userController *UserController) Update(c echo.Context) (err error) {
	var user model.User

	userRequest := new(model.UpdateUserRequest)
	if err = c.Bind(&userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err = validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusUnprocessableEntity, helper.ErrorResponse(validationErrors.Error()))
	}

	user, err = userController.UserService.UpdateUser(userRequest, c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(user.Response()))
}

func (userController *UserController) Destroy(c echo.Context) (err error) {
	if err := userController.UserService.DeleteUser(c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"Message": "User deleted successfully",
	}))
}
