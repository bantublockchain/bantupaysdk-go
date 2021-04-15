package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUsernameIsReserved = "error-username-is-reserved"

//ErrorUsernameIsReserved is returned during registration when a username already exists
type ErrorUsernameIsReserved struct {
	Username string
}

//Error returns the error string
func (e *ErrorUsernameIsReserved) Error() string {
	return errorUsernameIsReserved
}

//ErrorType returns error type as string
func (e *ErrorUsernameIsReserved) ErrorType() string {
	return errorUsernameIsReserved
}

//Data returns data of the error
func (e *ErrorUsernameIsReserved) Data() string {
	return "username"
}

//Message returns tring message of error
func (e *ErrorUsernameIsReserved) Message() string {
	return "Username is reserved"
}

//JSONError returns json of the error
func (e *ErrorUsernameIsReserved) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUsernameIsReserved) HTTPCode() int {
	return http.StatusConflict
}
