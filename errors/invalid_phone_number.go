package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidPhoneNumber = "error-invalid-phone-number"

//ErrorInvalidPhoneNumber  error
type ErrorInvalidPhoneNumber struct {
}

//Error returns the error string
func (e *ErrorInvalidPhoneNumber) Error() string {
	return errorInvalidPhoneNumber
}

//ErrorType returns error type as string
func (e *ErrorInvalidPhoneNumber) ErrorType() string {
	return errorInvalidPhoneNumber
}

//Data returns data of the error
func (e *ErrorInvalidPhoneNumber) Data() string {
	return "phoneNumber"
}

//Message returns tring message of error
func (e *ErrorInvalidPhoneNumber) Message() string {
	return "Invalid phone number."
}

//JSONError returns json of the error
func (e *ErrorInvalidPhoneNumber) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidPhoneNumber) HTTPCode() int {
	return http.StatusBadRequest
}
