package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorEmailAlreadyExists = "error-email-already-exists"

//ErrorEmailAlreadyExists is returned during registration when an email is already in use
type ErrorEmailAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorEmailAlreadyExists) Error() string {
	return errorEmailAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorEmailAlreadyExists) ErrorType() string {
	return errorEmailAlreadyExists
}

//Data returns data of the error
func (e *ErrorEmailAlreadyExists) Data() string {
	return "email"
}

//Message returns tring message of error
func (e *ErrorEmailAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorEmailAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorEmailAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
