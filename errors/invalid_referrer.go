package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidReferrer = "error-invalid-referrer"

//ErrorInvalidReferrer  error
type ErrorInvalidReferrer struct {
}

//Error returns the error string
func (e *ErrorInvalidReferrer) Error() string {
	return errorInvalidReferrer
}

//ErrorType returns error type as string
func (e *ErrorInvalidReferrer) ErrorType() string {
	return errorInvalidReferrer
}

//Data returns data of the error
func (e *ErrorInvalidReferrer) Data() string {
	return "referrer"
}

//Message returns tring message of error
func (e *ErrorInvalidReferrer) Message() string {
	return "Invalid referrer."
}

//JSONError returns json of the error
func (e *ErrorInvalidReferrer) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidReferrer) HTTPCode() int {
	return http.StatusBadRequest
}
