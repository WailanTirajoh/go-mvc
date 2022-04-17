package helper

type (
	SuccessResponseImpl struct {
		Data interface{} `json:"data"`
	}

	ValidationErrorImpl struct {
		Errors interface{} `json:"errors"`
	}
)

func SuccessResponse(data interface{}) SuccessResponseImpl {
	return SuccessResponseImpl{
		Data: data,
	}
}

func ValidationError(errors interface{}) ValidationErrorImpl {
	return ValidationErrorImpl{
		Errors: errors,
	}
}
