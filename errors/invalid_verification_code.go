package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidVerificationCode = "error-invalid-verification-code"

//ErrorInvalidVerificationCode for invalid email formats
type ErrorInvalidVerificationCode struct {
}

//Error returns the error string
func (e *ErrorInvalidVerificationCode) Error() string {
	return errorInvalidVerificationCode
}

//ErrorType returns error type as string
func (e *ErrorInvalidVerificationCode) ErrorType() string {
	return errorInvalidVerificationCode
}

//Data returns data of the error
func (e *ErrorInvalidVerificationCode) Data() string {
	return "verificationCode"
}

//Message returns tring message of error
func (e *ErrorInvalidVerificationCode) Message() string {
	return "Invalid verification code. Please ensure your information has not changed and try again later."
}

//JSONError returns json of the error
func (e *ErrorInvalidVerificationCode) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidVerificationCode) HTTPCode() int {
	return http.StatusBadRequest
}
