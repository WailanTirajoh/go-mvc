package controller

import (
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/labstack/echo"
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

	return c.JSON(http.StatusOK, helper.SuccessResponse(users))
}

func (userController *UserController) Show(c echo.Context) (err error) {
	user, err := userController.UserService.GetUser(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(user))
}

func (userController *UserController) Store(c echo.Context) (err error) {
	userRequest := new(model.StoreUserRequest)
	if err = c.Bind(&userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
	}

	if err = userController.UserService.StoreUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(user))
}

func (userController *UserController) Update(c echo.Context) (err error) {
	userRequest := new(model.UpdateUserRequest)
	if err = c.Bind(&userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
	}

	if err = userController.UserService.UpdateUser(c.Param("id"), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(user))
}

func (userController *UserController) Destroy(c echo.Context) (err error) {
	if err := userController.UserService.DeleteUser(c.Param("id")); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]string{
		"Message": "User deleted successfully",
	}))

}
