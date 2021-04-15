package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidPublicKey = "error-invalid-public-key"

//ErrorInvalidPublicKey is error struct
type ErrorInvalidPublicKey struct {
	PublicKey string
}

//Error returns the error string
func (e *ErrorInvalidPublicKey) Error() string {
	return errorInvalidPublicKey
}

//ErrorType returns error type as string
func (e *ErrorInvalidPublicKey) ErrorType() string {
	return errorInvalidPublicKey
}

//Data returns data of the error
func (e *ErrorInvalidPublicKey) Data() string {
	return errorInvalidPublicKey
}

//Message returns tring message of error
func (e *ErrorInvalidPublicKey) Message() string {
	return "Invalid Public Key [" + e.PublicKey + "]"
}

//JSONError returns json of the error
func (e *ErrorInvalidPublicKey) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "publicKey", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidPublicKey) HTTPCode() int {
	return http.StatusBadRequest
}
