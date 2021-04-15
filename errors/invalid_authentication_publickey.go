package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//errorInvalidAuthenticationPublicKey holds error of public key
const errorInvalidAuthenticationPublicKey = "error-invalid-authentication-public-key"

//ErrorInvalidAuthenticationPublicKey is struct for returning error
type ErrorInvalidAuthenticationPublicKey struct {
	PublicKey string
}

//Error returns the error string
func (e *ErrorInvalidAuthenticationPublicKey) Error() string {
	return errorInvalidPublicKey
}

//ErrorType returns error type as string
func (e *ErrorInvalidAuthenticationPublicKey) ErrorType() string {
	return errorInvalidAuthenticationPublicKey
}

//Data returns data of the error
func (e *ErrorInvalidAuthenticationPublicKey) Data() string {
	return "X-BANTUPAY-PUBLIC-KEY"
}

//Message returns tring message of error
func (e *ErrorInvalidAuthenticationPublicKey) Message() string {
	return "Invalid Header: [X-BANTUPAY-PUBLIC-KEY= " + e.PublicKey + "]"
}

//JSONError returns json of the error
func (e *ErrorInvalidAuthenticationPublicKey) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidAuthenticationPublicKey) HTTPCode() int {
	return http.StatusUnauthorized
}
