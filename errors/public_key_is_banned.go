package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorAccountIsBanned = "error-Account-is-banned"

//ErrorAccountIsBanned is returned during registration when a username already exists
type ErrorAccountIsBanned struct {
	PublicKey string
}

//Error returns the error string
func (e *ErrorAccountIsBanned) Error() string {
	return errorAccountIsBanned
}

//ErrorType returns error type as string
func (e *ErrorAccountIsBanned) ErrorType() string {
	return errorAccountIsBanned
}

//Data returns data of the error
func (e *ErrorAccountIsBanned) Data() string {
	return "publicKey"
}

//Message returns tring message of error
func (e *ErrorAccountIsBanned) Message() string {
	return "Account is banned"
}

//JSONError returns json of the error
func (e *ErrorAccountIsBanned) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorAccountIsBanned) HTTPCode() int {
	return http.StatusForbidden
}
