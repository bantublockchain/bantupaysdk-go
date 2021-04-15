package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidImageThumbnail = "error-invalid-image-thumbnail"

//ErrorInvalidImageThumbnail error
type ErrorInvalidImageThumbnail struct {
}

//Error returns the error string
func (e *ErrorInvalidImageThumbnail) Error() string {
	return errorInvalidImageThumbnail
}

//ErrorType returns error type as string
func (e *ErrorInvalidImageThumbnail) ErrorType() string {
	return errorInvalidImageThumbnail
}

//Data returns data of the error
func (e *ErrorInvalidImageThumbnail) Data() string {
	return "imageThumbnail"
}

//Message returns tring message of error
func (e *ErrorInvalidImageThumbnail) Message() string {
	return "Invalid image thumbnail."
}

//JSONError returns json of the error
func (e *ErrorInvalidImageThumbnail) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidImageThumbnail) HTTPCode() int {
	return http.StatusBadRequest
}
