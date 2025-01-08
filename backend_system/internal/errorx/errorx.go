package errorx

/**
* description:
* author: wqh
* date: 2025/1/8
 */
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

func GetError(e *Error, data interface{}) *Error {
	return &Error{Code: e.Code, Message: e.Message, Data: data}
}
