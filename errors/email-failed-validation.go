package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorEmailFailedValidation = "error-email-failed-validation"

//ErrorEmailFailedValidation for invalid email formats
type ErrorEmailFailedValidation struct {
	Email  string
	Detail string
}

//Error returns the error string
func (e *ErrorEmailFailedValidation) Error() string {
	return errorEmailFailedValidation
}

//ErrorType returns error type as string
func (e *ErrorEmailFailedValidation) ErrorType() string {
	return errorEmailFailedValidation
}

//Data returns data of the error
func (e *ErrorEmailFailedValidation) Data() string {
	return errorInvalidEmail
}

//Message returns tring message of error
func (e *ErrorEmailFailedValidation) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorEmailFailedValidation) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "email", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorEmailFailedValidation) HTTPCode() int {
	return http.StatusBadRequest
}
