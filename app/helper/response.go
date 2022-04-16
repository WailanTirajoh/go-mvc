package helper

type (
	SuccessResponseImpl struct {
		Data interface{} `json:"data"`
	}
)

func SuccessResponse(data interface{}) SuccessResponseImpl {
	return SuccessResponseImpl{
		Data: data,
	}
}
