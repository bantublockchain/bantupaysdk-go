package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUnimplementedFeature = "error-unimplemented-feature"

//ErrorUnimplementedFeature error struct
type ErrorUnimplementedFeature struct {
}

//Error returns the error string
func (e *ErrorUnimplementedFeature) Error() string {
	return errorUnimplementedFeature
}

//ErrorType returns error type as string
func (e *ErrorUnimplementedFeature) ErrorType() string {
	return errorUnimplementedFeature
}

//Data returns data of the error
func (e *ErrorUnimplementedFeature) Data() string {
	return ""
}

//Message returns tring message of error
func (e *ErrorUnimplementedFeature) Message() string {
	return errorUnimplementedFeature
}

//JSONError returns json of the error
func (e *ErrorUnimplementedFeature) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUnimplementedFeature) HTTPCode() int {
	return http.StatusNotImplemented
}
