package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidLinkedinFormat = "error-invalid-linkedin-format"

//ErrorInvalidLinkedinFormat is returned when there is an error in the username
type ErrorInvalidLinkedinFormat struct {
	Handle string
	Detail string
}

//Error returns the error string
func (e *ErrorInvalidLinkedinFormat) Error() string {
	return errorInvalidLinkedinFormat
}

//ErrorType returns error type as string
func (e *ErrorInvalidLinkedinFormat) ErrorType() string {
	return errorInvalidLinkedinFormat
}

//Data returns data of the error
func (e *ErrorInvalidLinkedinFormat) Data() string {
	return "linkedin"
}

//Message returns tring message of error
func (e *ErrorInvalidLinkedinFormat) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorInvalidLinkedinFormat) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "linkedIn", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidLinkedinFormat) HTTPCode() int {
	return http.StatusBadRequest
}
