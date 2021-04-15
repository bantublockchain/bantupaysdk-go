package errors

import (
	"github.com/gin-gonic/gin"
)

const errorWalletMigrationFailure = "error-wallet-migration-failure"

//ErrorWalletMigrationFailure holds error struct
type ErrorWalletMigrationFailure struct {
}

//Error returns the error string
func (e *ErrorWalletMigrationFailure) Error() string {
	return errorWalletMigrationFailure
}

//ErrorType returns error type as string
func (e *ErrorWalletMigrationFailure) ErrorType() string {
	return errorWalletMigrationFailure
}

//Data returns data of the error
func (e *ErrorWalletMigrationFailure) Data() string {
	return "failed to migrate wallet"
}

//Message returns tring message of error
func (e *ErrorWalletMigrationFailure) Message() string {
	return "failure restoring your wallet"
}

//JSONError returns json of the error
func (e *ErrorWalletMigrationFailure) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "Failure", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorWalletMigrationFailure) HTTPCode() int {
	return 500
}
