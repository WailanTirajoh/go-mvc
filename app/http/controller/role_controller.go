package controller

import (
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/service"
	"github.com/labstack/echo/v4"
)

type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(roleService *service.RoleService) RoleController {
	return RoleController{
		RoleService: *roleService,
	}
}

func (roleController RoleController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]interface{}{
		"roles": "Index",
	}))
}

func (roleController RoleController) Store(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]interface{}{
		"roles": "Store",
	}))
}

func (roleController RoleController) Show(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]interface{}{
		"roles": "Show",
	}))
}

func (roleController RoleController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]interface{}{
		"roles": "Update",
	}))
}

func (roleController RoleController) Destroy(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SuccessResponse(map[string]interface{}{
		"roles": "Destroy",
	}))
}
