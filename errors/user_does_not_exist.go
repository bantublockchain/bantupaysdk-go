package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUserDoesNotExist = "error-username-does-not-exist"

//ErrorUserDoesNotExist is when user does not exist
type ErrorUserDoesNotExist struct {
	Username string
}

//Error returns the error string
func (e *ErrorUserDoesNotExist) Error() string {
	return errorUserDoesNotExist
}

//ErrorType returns error type as string
func (e *ErrorUserDoesNotExist) ErrorType() string {
	return errorUserDoesNotExist
}

//Data returns data of the error
func (e *ErrorUserDoesNotExist) Data() string {
	return fmt.Sprintf("%v", e.Username)
}

//Message returns tring message of error
func (e *ErrorUserDoesNotExist) Message() string {
	return fmt.Sprintf("username: %v does not exist", e.Username)
}

//JSONError returns json of the error
func (e *ErrorUserDoesNotExist) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUserDoesNotExist) HTTPCode() int {
	return http.StatusNotFound
}
