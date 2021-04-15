package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorMissingParameter = "error-missing-parameter"

// ErrorMissingParameter error struct
type ErrorMissingParameter struct {
	Parameter string
}

//Error returns the error string
func (e *ErrorMissingParameter) Error() string {
	return errorMissingParameter
}

//ErrorType returns error type as string
func (e *ErrorMissingParameter) ErrorType() string {
	return errorMissingParameter
}

//Data returns data of the error
func (e *ErrorMissingParameter) Data() string {
	return e.Parameter
}

//Message returns tring message of error
func (e *ErrorMissingParameter) Message() string {
	return "Missing Required Parameter"
}

//JSONError returns json of the error
func (e *ErrorMissingParameter) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorMissingParameter) HTTPCode() int {
	return http.StatusBadRequest
}
