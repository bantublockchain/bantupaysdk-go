package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorTelegramHandleAlreadyExists = "error-telegram-handle-already-exists"

//ErrorTelegramHandleAlreadyExists is returned during registration when a username already exists
type ErrorTelegramHandleAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorTelegramHandleAlreadyExists) Error() string {
	return errorTelegramHandleAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorTelegramHandleAlreadyExists) ErrorType() string {
	return errorTelegramHandleAlreadyExists
}

//Data returns data of the error
func (e *ErrorTelegramHandleAlreadyExists) Data() string {
	return "telegram"
}

//Message returns tring message of error
func (e *ErrorTelegramHandleAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorTelegramHandleAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorTelegramHandleAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
