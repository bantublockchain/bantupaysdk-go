package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidTelegramHandleFormat = "error-invalid-telegram-handle-format"

//ErrorInvalidTelegramHandleFormat is returned when there is an error in the username
type ErrorInvalidTelegramHandleFormat struct {
	Handle string
	Detail string
}

//Error returns the error string
func (e *ErrorInvalidTelegramHandleFormat) Error() string {
	return errorInvalidTelegramHandleFormat
}

//ErrorType returns error type as string
func (e *ErrorInvalidTelegramHandleFormat) ErrorType() string {
	return errorInvalidTelegramHandleFormat
}

//Data returns data of the error
func (e *ErrorInvalidTelegramHandleFormat) Data() string {
	return "telegram"
}

//Message returns tring message of error
func (e *ErrorInvalidTelegramHandleFormat) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInvalidTelegramHandleFormat) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "telegram", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidTelegramHandleFormat) HTTPCode() int {
	return http.StatusBadRequest
}
