package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidUsernameFormat = "error-invalid-username-format"

//ErrorInvalidUsernameFormat is returned when there is an error in the username
type ErrorInvalidUsernameFormat struct {
	Username string
	Detail   string
}

//Error returns the error string
func (e *ErrorInvalidUsernameFormat) Error() string {
	return errorInvalidUsernameFormat
}

//ErrorType returns error type as string
func (e *ErrorInvalidUsernameFormat) ErrorType() string {
	return errorInvalidUsernameFormat
}

//Data returns data of the error
func (e *ErrorInvalidUsernameFormat) Data() string {
	return "username"
}

//Message returns tring message of error
func (e *ErrorInvalidUsernameFormat) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInvalidUsernameFormat) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "username", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidUsernameFormat) HTTPCode() int {
	return http.StatusBadRequest
}
