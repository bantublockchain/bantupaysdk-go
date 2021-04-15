package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidInstagramHandleFormat = "error-invalid-instagram-handle-format"

//ErrorInvalidInstagramHandleFormat is returned when there is an error in the username
type ErrorInvalidInstagramHandleFormat struct {
	Handle string
	Detail string
}

//Error returns the error string
func (e *ErrorInvalidInstagramHandleFormat) Error() string {
	return errorInvalidInstagramHandleFormat
}

//ErrorType returns error type as string
func (e *ErrorInvalidInstagramHandleFormat) ErrorType() string {
	return errorInvalidInstagramHandleFormat
}

//Data returns data of the error
func (e *ErrorInvalidInstagramHandleFormat) Data() string {
	return "instagram"
}

//Message returns tring message of error
func (e *ErrorInvalidInstagramHandleFormat) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInvalidInstagramHandleFormat) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "instagram", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidInstagramHandleFormat) HTTPCode() int {
	return http.StatusBadRequest
}
