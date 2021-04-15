package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorVerificationCodeNotSent = "error-verification-code-not-sent"

//ErrorVerificationCodeNotSent for verification codes that couldn't be sent for one reason or another
type ErrorVerificationCodeNotSent struct {
}

//Error returns the error string
func (e *ErrorVerificationCodeNotSent) Error() string {
	return errorVerificationCodeNotSent
}

//ErrorType returns error type as string
func (e *ErrorVerificationCodeNotSent) ErrorType() string {
	return errorVerificationCodeNotSent
}

//Data returns data of the error
func (e *ErrorVerificationCodeNotSent) Data() string {
	return "verificationCode"
}

//Message returns tring message of error
func (e *ErrorVerificationCodeNotSent) Message() string {
	return "Verification code not sent. Contact support."
}

//JSONError returns json of the error
func (e *ErrorVerificationCodeNotSent) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorVerificationCodeNotSent) HTTPCode() int {
	return http.StatusServiceUnavailable
}
