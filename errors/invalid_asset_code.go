package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidAssetCode = "error-invalid-asset-code"

//ErrorInvalidAssetCode for invalid asset codes
type ErrorInvalidAssetCode struct {
}

//Error returns the error string
func (e *ErrorInvalidAssetCode) Error() string {
	return errorInvalidAssetCode
}

//ErrorType returns error type as string
func (e *ErrorInvalidAssetCode) ErrorType() string {
	return errorInvalidAssetCode
}

//Data returns data of the error
func (e *ErrorInvalidAssetCode) Data() string {
	return "assetCode"
}

//Message sets the message on the error
func (e *ErrorInvalidAssetCode) Message() string {
	return "Invalid asset code, Asset codes can only have at most 12 latin characters"
}

//JSONError returns json of the error
func (e *ErrorInvalidAssetCode) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidAssetCode) HTTPCode() int {
	return http.StatusUnauthorized
}
