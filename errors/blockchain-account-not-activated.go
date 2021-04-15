package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorBlockchainAccountNotActivated = "error-blockchain-account-not-activated"

//ErrorBlockchainAccountNotActivated is when user does not exist
type ErrorBlockchainAccountNotActivated struct {
	PublicKey string
}

//Error returns the error string
func (e *ErrorBlockchainAccountNotActivated) Error() string {
	return errorBlockchainAccountNotActivated
}

//ErrorType returns error type as string
func (e *ErrorBlockchainAccountNotActivated) ErrorType() string {
	return errorBlockchainAccountNotActivated
}

//Data returns data of the error
func (e *ErrorBlockchainAccountNotActivated) Data() string {
	return fmt.Sprintf("%v", e.PublicKey)
}

//Message returns tring message of error
func (e *ErrorBlockchainAccountNotActivated) Message() string {
	return "Account not funded"
}

//JSONError returns json of the error
func (e *ErrorBlockchainAccountNotActivated) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorBlockchainAccountNotActivated) HTTPCode() int {
	return http.StatusNotFound
}
