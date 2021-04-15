package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidGender = "error-invalid-gender"

//ErrorInvalidGender error
type ErrorInvalidGender struct {
}

//Error returns the error string
func (e *ErrorInvalidGender) Error() string {
	return errorInvalidGender
}

//ErrorType returns error type as string
func (e *ErrorInvalidGender) ErrorType() string {
	return errorInvalidGender
}

//Data returns data of the error
func (e *ErrorInvalidGender) Data() string {
	return "gender"
}

//Message returns tring message of error
func (e *ErrorInvalidGender) Message() string {
	return "invalid gender"
}

//JSONError returns json of the error
func (e *ErrorInvalidGender) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidGender) HTTPCode() int {
	return http.StatusBadRequest
}
