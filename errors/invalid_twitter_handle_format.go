package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidTwitterHandleFormat = "error-invalid-twitter-handle-format"

//ErrorInvalidTwitterHandleFormat is returned when there is an error in the username
type ErrorInvalidTwitterHandleFormat struct {
	Handle string
	Detail string
}

//Error returns the error string
func (e *ErrorInvalidTwitterHandleFormat) Error() string {
	return errorInvalidTwitterHandleFormat
}

//ErrorType returns error type as string
func (e *ErrorInvalidTwitterHandleFormat) ErrorType() string {
	return errorInvalidTwitterHandleFormat
}

//Data returns data of the error
func (e *ErrorInvalidTwitterHandleFormat) Data() string {
	return "twitter"
}

//Message returns tring message of error
func (e *ErrorInvalidTwitterHandleFormat) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInvalidTwitterHandleFormat) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "twitter", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidTwitterHandleFormat) HTTPCode() int {
	return http.StatusBadRequest
}
