package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorTwitterHandleAlreadyExists = "error-twitter-handle-already-exists"

//ErrorTwitterHandleAlreadyExists is returned during registration when a username already exists
type ErrorTwitterHandleAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorTwitterHandleAlreadyExists) Error() string {
	return errorTwitterHandleAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorTwitterHandleAlreadyExists) ErrorType() string {
	return errorTwitterHandleAlreadyExists
}

//Data returns data of the error
func (e *ErrorTwitterHandleAlreadyExists) Data() string {
	return "twitter"
}

//Message returns tring message of error
func (e *ErrorTwitterHandleAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorTwitterHandleAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorTwitterHandleAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
