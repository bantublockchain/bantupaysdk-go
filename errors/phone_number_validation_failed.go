package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorPhoneNumberValidationFailed = "error-phone-number-validation-failed"

//ErrorPhoneNumberValidationFailed  error
type ErrorPhoneNumberValidationFailed struct {
	Detail string
}

//Error returns the error string
func (e *ErrorPhoneNumberValidationFailed) Error() string {
	return errorPhoneNumberValidationFailed
}

//ErrorType returns error type as string
func (e *ErrorPhoneNumberValidationFailed) ErrorType() string {
	return errorPhoneNumberValidationFailed
}

//Data returns data of the error
func (e *ErrorPhoneNumberValidationFailed) Data() string {
	return "mobile"
}

//Message returns tring message of error
func (e *ErrorPhoneNumberValidationFailed) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorPhoneNumberValidationFailed) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorPhoneNumberValidationFailed) HTTPCode() int {
	return http.StatusBadRequest
}
