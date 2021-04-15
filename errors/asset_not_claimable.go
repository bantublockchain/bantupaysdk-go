package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorAssetNotClaimable = "error-asset-not-claimable"

//ErrorAssetNotClaimable for unclaimable assets
type ErrorAssetNotClaimable struct {
}

//Error returns the error string
func (e *ErrorAssetNotClaimable) Error() string {
	return errorAssetNotClaimable
}

//ErrorType returns error type as string
func (e *ErrorAssetNotClaimable) ErrorType() string {
	return errorAssetNotClaimable
}

//Data returns data of the error
func (e *ErrorAssetNotClaimable) Data() string {
	return "assetCode"
}

//Message returns tring message of error
func (e *ErrorAssetNotClaimable) Message() string {
	return "The asset you are trying to claim cannot be claimed."
}

//JSONError returns json of the error
func (e *ErrorAssetNotClaimable) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorAssetNotClaimable) HTTPCode() int {
	return http.StatusBadRequest
}
