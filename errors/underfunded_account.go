package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUnderfundedAccount = "error-underfunded-account"

//ErrorUnderfundedAccount gets returned when an account doesn't have the necessary funds to perform a transaction
type ErrorUnderfundedAccount struct {
	Detail string
}

//Error returns the error string
func (e *ErrorUnderfundedAccount) Error() string {
	return errorUnderfundedAccount
}

//ErrorType returns error type as string
func (e *ErrorUnderfundedAccount) ErrorType() string {
	return errorUnderfundedAccount
}

//Data returns data of the error
func (e *ErrorUnderfundedAccount) Data() string {
	return "amount"
}

//Message returns tring message of error
func (e *ErrorUnderfundedAccount) Message() string {
	if len(e.Detail) > 0 {
		return e.Detail
	}
	return "Not enough funds"
}

//JSONError returns json of the error
func (e *ErrorUnderfundedAccount) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUnderfundedAccount) HTTPCode() int {
	return http.StatusBadRequest
}
