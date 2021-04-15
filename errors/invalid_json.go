package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidJSON = "error-invalid-json"

//ErrorInvalidJSON is error struct
type ErrorInvalidJSON struct {
}

//Error returns the error string
func (e *ErrorInvalidJSON) Error() string {
	return errorMissingParameter
}

//ErrorType returns error type as string
func (e *ErrorInvalidJSON) ErrorType() string {
	return errorInvalidJSON
}

//Data returns data of the error
func (e *ErrorInvalidJSON) Data() string {
	return errorInvalidJSON
}

//Message returns tring message of error
func (e *ErrorInvalidJSON) Message() string {
	return "The JSON data you submitted is not well formed JSON!"
}

//JSONError returns json of the error
func (e *ErrorInvalidJSON) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": nil, "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidJSON) HTTPCode() int {
	return http.StatusBadRequest
}
