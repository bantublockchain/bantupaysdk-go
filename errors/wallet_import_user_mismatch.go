package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorWalletUserMismatch = "error-wallet-user-mismatch"

//ErrorWalletUserMismatch holds error struct
type ErrorWalletUserMismatch struct {
}

//Error returns the error string
func (e *ErrorWalletUserMismatch) Error() string {
	return errorWalletUserMismatch
}

//ErrorType returns error type as string
func (e *ErrorWalletUserMismatch) ErrorType() string {
	return errorWalletUserMismatch
}

//Data returns data of the error
func (e *ErrorWalletUserMismatch) Data() string {
	return "wallet user mismatch"
}

//Message returns tring message of error
func (e *ErrorWalletUserMismatch) Message() string {
	return "you are not authorized to import this wallet with this user"
}

//JSONError returns json of the error
func (e *ErrorWalletUserMismatch) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "Authorization", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorWalletUserMismatch) HTTPCode() int {
	return http.StatusUnauthorized
}
