package helper

type SuccessResponseImpl struct {
	IsOk bool        `json:"is_ok"`
	Data interface{} `json:"data"`
}

type ErrorResponseImpl struct {
	IsOk    bool   `json:"is_ok"`
	Message string `json:"message"`
}

func SuccessResponse(data interface{}) SuccessResponseImpl {
	return SuccessResponseImpl{
		IsOk: true,
		Data: data,
	}
}

func ErrorResponse(message string) ErrorResponseImpl {
	return ErrorResponseImpl{
		IsOk:    false,
		Message: message,
	}
}
