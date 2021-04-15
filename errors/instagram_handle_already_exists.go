package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInstagramHandleAlreadyExists = "error-instagram-handle-already-exists"

//ErrorInstagramHandleAlreadyExists is returned during registration when a username already exists
type ErrorInstagramHandleAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorInstagramHandleAlreadyExists) Error() string {
	return errorInstagramHandleAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorInstagramHandleAlreadyExists) ErrorType() string {
	return errorInstagramHandleAlreadyExists
}

//Data returns data of the error
func (e *ErrorInstagramHandleAlreadyExists) Data() string {
	return "instagram"
}

//Message returns tring message of error
func (e *ErrorInstagramHandleAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInstagramHandleAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInstagramHandleAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
