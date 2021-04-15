package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidTransaction = "error-invalid-transaction"

//ErrorInvalidTransaction error
type ErrorInvalidTransaction struct {
}

//Error returns the error string
func (e *ErrorInvalidTransaction) Error() string {
	return errorInvalidTransaction
}

//ErrorType returns error type as string
func (e *ErrorInvalidTransaction) ErrorType() string {
	return errorInvalidTransaction
}

//Data returns data of the error
func (e *ErrorInvalidTransaction) Data() string {
	return "transaction"
}

//Message returns tring message of error
func (e *ErrorInvalidTransaction) Message() string {
	return "Invalid Transaction"
}

//JSONError returns json of the error
func (e *ErrorInvalidTransaction) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidTransaction) HTTPCode() int {
	return http.StatusBadRequest
}
