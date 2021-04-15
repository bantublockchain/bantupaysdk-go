package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorTemporaryServerError = "error-temporary-server-error"

//ErrorTemporaryServerError temporary server error
type ErrorTemporaryServerError struct {
}

//Error returns the error string
func (e *ErrorTemporaryServerError) Error() string {
	return errorTemporaryServerError
}

//ErrorType returns error type as string
func (e *ErrorTemporaryServerError) ErrorType() string {
	return errorTemporaryServerError
}

//Data returns data of the error
func (e *ErrorTemporaryServerError) Data() string {
	return "temporaryServerError"
}

//Message returns tring message of error
func (e *ErrorTemporaryServerError) Message() string {
	return "Temporary Server Error. Contact support."
}

//JSONError returns json of the error
func (e *ErrorTemporaryServerError) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorTemporaryServerError) HTTPCode() int {
	return http.StatusServiceUnavailable
}
