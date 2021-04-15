package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUsernameIsSuspended = "error-username-is-suspended"

//ErrorUsernameIsSuspended is returned during registration when a username already exists
type ErrorUsernameIsSuspended struct {
	Username string
}

//Error returns the error string
func (e *ErrorUsernameIsSuspended) Error() string {
	return errorUsernameIsSuspended
}

//ErrorType returns error type as string
func (e *ErrorUsernameIsSuspended) ErrorType() string {
	return errorUsernameIsSuspended
}

//Data returns data of the error
func (e *ErrorUsernameIsSuspended) Data() string {
	return "username"
}

//Message returns tring message of error
func (e *ErrorUsernameIsSuspended) Message() string {
	return "Account is suspended"
}

//JSONError returns json of the error
func (e *ErrorUsernameIsSuspended) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUsernameIsSuspended) HTTPCode() int {
	return http.StatusForbidden
}
