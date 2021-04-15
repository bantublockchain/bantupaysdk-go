package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorMobileNumberAlreadyExists = "error-mobile-number-already-exists"

//ErrorMobileNumberAlreadyExists is returned during registration when a username already exists
type ErrorMobileNumberAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorMobileNumberAlreadyExists) Error() string {
	return errorMobileNumberAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorMobileNumberAlreadyExists) ErrorType() string {
	return errorMobileNumberAlreadyExists
}

//Data returns data of the error
func (e *ErrorMobileNumberAlreadyExists) Data() string {
	return "mobile"
}

//Message returns tring message of error
func (e *ErrorMobileNumberAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorMobileNumberAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorMobileNumberAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
