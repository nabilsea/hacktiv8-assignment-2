package util

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors,omitempty"`
	Data   interface{} `json:"data"`
}

func APIResponse(code int, status string, errors interface{}, data interface{}) Response {
	return Response{
		Code:   code,
		Status: status,
		Errors: errors,
		Data:   data,
	}
}
