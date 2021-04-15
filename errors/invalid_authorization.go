package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidAuthorization = "error-invalid-authorization"

//ErrorInvalidAuthorization holds error struct
type ErrorInvalidAuthorization struct {
}

//Error returns the error string
func (e *ErrorInvalidAuthorization) Error() string {
	return errorInvalidAuthorization
}

//ErrorType returns error type as string
func (e *ErrorInvalidAuthorization) ErrorType() string {
	return errorInvalidAuthorization
}

//Data returns data of the error
func (e *ErrorInvalidAuthorization) Data() string {
	return "Unauthorized"
}

//Message returns tring message of error
func (e *ErrorInvalidAuthorization) Message() string {
	return "you are not authorized to view this information"
}

//JSONError returns json of the error
func (e *ErrorInvalidAuthorization) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "Authorization", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidAuthorization) HTTPCode() int {
	return http.StatusUnauthorized
}
