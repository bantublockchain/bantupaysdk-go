package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorInvalidImageThumbnailSize = "error-invalid-image-thumbnail-size"

//ErrorInvalidImageThumbnailSize error
type ErrorInvalidImageThumbnailSize struct {
}

//Error returns the error string
func (e *ErrorInvalidImageThumbnailSize) Error() string {
	return errorInvalidImageThumbnailSize
}

//ErrorType returns error type as string
func (e *ErrorInvalidImageThumbnailSize) ErrorType() string {
	return errorInvalidImageThumbnailSize
}

//Data returns data of the error
func (e *ErrorInvalidImageThumbnailSize) Data() string {
	return "imageThumbnailSize"
}

//Message returns tring message of error
func (e *ErrorInvalidImageThumbnailSize) Message() string {
	return "image thumbnail size more than 25kbytes"
}

//JSONError returns json of the error
func (e *ErrorInvalidImageThumbnailSize) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorInvalidImageThumbnailSize) HTTPCode() int {
	return http.StatusBadRequest
}
