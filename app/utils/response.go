package utils

type Response struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // message
	Data    interface{} `json:"data"`    // data
}

// 成功的返回
func Success(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// 失败的返回
func Error(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
