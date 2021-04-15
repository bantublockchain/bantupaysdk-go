package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUsernameAlreadyExists = "error-username-already-exists"

//ErrorUsernameAlreadyExists is returned during registration when a username already exists
type ErrorUsernameAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorUsernameAlreadyExists) Error() string {
	return errorUsernameAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorUsernameAlreadyExists) ErrorType() string {
	return errorUsernameAlreadyExists
}

//Data returns data of the error
func (e *ErrorUsernameAlreadyExists) Data() string {
	return "username"
}

//Message returns tring message of error
func (e *ErrorUsernameAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorUsernameAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUsernameAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
