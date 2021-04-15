package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorNameFailedValidation = "error-name-failed-validation"

//ErrorNameFailedValidation for invalid email formats
type ErrorNameFailedValidation struct {
	Detail string
}

//Error returns the error string
func (e *ErrorNameFailedValidation) Error() string {
	return errorNameFailedValidation
}

//ErrorType returns error type as string
func (e *ErrorNameFailedValidation) ErrorType() string {
	return errorNameFailedValidation
}

//Data returns data of the error
func (e *ErrorNameFailedValidation) Data() string {
	return e.Detail
}

//Message returns tring message of error
func (e *ErrorNameFailedValidation) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorNameFailedValidation) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorNameFailedValidation) HTTPCode() int {
	return http.StatusBadRequest
}
