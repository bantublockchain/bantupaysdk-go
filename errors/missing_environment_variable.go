package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorMissingEnvironmentVariable = "error-missing-environment-variable"

//ErrorMissingEnvironmentVariable is returned when an environment variable has not yet been set
type ErrorMissingEnvironmentVariable struct {
	EnvironmentVariable string
}

//Error returns the error string
func (e *ErrorMissingEnvironmentVariable) Error() string {
	return errorMissingEnvironmentVariable
}

//ErrorType returns error type as string
func (e *ErrorMissingEnvironmentVariable) ErrorType() string {
	return errorMissingEnvironmentVariable
}

//Data returns data of the error
func (e *ErrorMissingEnvironmentVariable) Data() string {
	return e.EnvironmentVariable
}

//Message returns tring message of error
func (e *ErrorMissingEnvironmentVariable) Message() string {
	return "Missing environment variable [" + e.EnvironmentVariable + "]"
}

//JSONError returns json of the error
func (e *ErrorMissingEnvironmentVariable) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorMissingEnvironmentVariable) HTTPCode() int {
	return http.StatusBadRequest
}
