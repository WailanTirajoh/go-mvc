package request

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse map[string]interface{}
	ErrorOutput   map[string]interface{}
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}

func MessageTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	default:
		return fe.Error()
	}
}

func Output(ve validator.ValidationErrors) []ErrorResponse {
	out := make([]ErrorResponse, len(ve))
	for i, fe := range ve {
		out[i] = ErrorResponse{
			fe.Field(): MessageTag(fe),
		}
	}

	return out
}
