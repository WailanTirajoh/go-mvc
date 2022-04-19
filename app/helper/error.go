package helper

import (
	"errors"
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/request"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func HandleError(c echo.Context, err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return c.JSON(http.StatusBadRequest, ValidationError(request.Output(ve)))
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
