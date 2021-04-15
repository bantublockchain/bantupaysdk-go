package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidEmail = "error-invalid-email-format"

//ErrorInvalidEmailFormat for invalid email formats
type ErrorInvalidEmailFormat struct {
	Email  string
	Detail string
}

//Error returns the error string
func (e *ErrorInvalidEmailFormat) Error() string {
	return errorInvalidEmail
}

//ErrorType returns error type as string
func (e *ErrorInvalidEmailFormat) ErrorType() string {
	return errorInvalidEmail
}

//Data returns data of the error
func (e *ErrorInvalidEmailFormat) Data() string {
	return errorInvalidEmail
}

//Message returns tring message of error
func (e *ErrorInvalidEmailFormat) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInvalidEmailFormat) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "email", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidEmailFormat) HTTPCode() int {
	return http.StatusBadRequest
}
