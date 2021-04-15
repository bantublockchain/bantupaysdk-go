package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorBantuTalkUsernameAlreadyExists = "error-bantu-talk-username-already-exists"

//ErrorBantuTalkUsernameAlreadyExists is returned during registration when a username already exists
type ErrorBantuTalkUsernameAlreadyExists struct {
	Detail string
}

//Error returns the error string
func (e *ErrorBantuTalkUsernameAlreadyExists) Error() string {
	return errorBantuTalkUsernameAlreadyExists
}

//ErrorType returns error type as string
func (e *ErrorBantuTalkUsernameAlreadyExists) ErrorType() string {
	return errorBantuTalkUsernameAlreadyExists
}

//Data returns data of the error
func (e *ErrorBantuTalkUsernameAlreadyExists) Data() string {
	return "bantuTalk"
}

//Message returns tring message of error
func (e *ErrorBantuTalkUsernameAlreadyExists) Message() string {
	return e.Detail
}

//JSONError returns json of the error
func (e *ErrorBantuTalkUsernameAlreadyExists) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": e.Data(), "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorBantuTalkUsernameAlreadyExists) HTTPCode() int {
	return http.StatusConflict
}
