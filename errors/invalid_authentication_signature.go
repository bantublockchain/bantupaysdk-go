package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidAuthenticationSignature = "error-invalid-authentication-signature"

//ErrorInvalidAuthenticationSignature holds error struct
type ErrorInvalidAuthenticationSignature struct {
}

//Error returns the error string
func (e *ErrorInvalidAuthenticationSignature) Error() string {
	return errorInvalidAuthenticationSignature
}

//ErrorType returns error type as string
func (e *ErrorInvalidAuthenticationSignature) ErrorType() string {
	return errorInvalidAuthenticationSignature
}

//Data returns data of the error
func (e *ErrorInvalidAuthenticationSignature) Data() string {
	return "X-BANTUPAY-SIGNATURE"
}

//Message returns tring message of error
func (e *ErrorInvalidAuthenticationSignature) Message() string {
	return "Authentication Signature Header [X-BANTUPAY-SIGNATURE] is invalid"
}

//JSONError returns json of the error
func (e *ErrorInvalidAuthenticationSignature) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "Authentication", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidAuthenticationSignature) HTTPCode() int {
	return http.StatusUnauthorized
}
