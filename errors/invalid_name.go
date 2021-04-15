package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidName = "error-invalid-name"

//ErrorInvalidName invalid name error
type ErrorInvalidName struct {
	Field string
}

//Error returns the error string
func (e *ErrorInvalidName) Error() string {
	return errorInvalidName
}

//ErrorType returns error type as string
func (e *ErrorInvalidName) ErrorType() string {
	return errorInvalidName
}

//Data returns data of the error
func (e *ErrorInvalidName) Data() string {
	return e.Field
}

//Message returns tring message of error
func (e *ErrorInvalidName) Message() string {
	return fmt.Sprintf("Invalid %v", e.Field)
}

//JSONError returns json of the error
func (e *ErrorInvalidName) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidName) HTTPCode() int {
	return http.StatusBadRequest
}
