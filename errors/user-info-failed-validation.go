package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorUserInformationFailedValidation = "error-registration-information-failed-validation"

//ErrorUserInformationFailedValidation for invalid email formats
type ErrorUserInformationFailedValidation struct {
}

//Error returns the error string
func (e *ErrorUserInformationFailedValidation) Error() string {
	return errorUserInformationFailedValidation
}

//ErrorType returns error type as string
func (e *ErrorUserInformationFailedValidation) ErrorType() string {
	return errorUserInformationFailedValidation
}

//Data returns data of the error
func (e *ErrorUserInformationFailedValidation) Data() string {
	return errorUserInformationFailedValidation
}

//Message returns tring message of error
func (e *ErrorUserInformationFailedValidation) Message() string {
	return "Sorry, your registration was not successful at this time. Please try again later"
}

//JSONError returns json of the error
func (e *ErrorUserInformationFailedValidation) JSONError() gin.H {
	return gin.H{"error": e.ErrorType(), "data": "username", "message": e.Message()}
}

//HTTPCode returns http status code
func (e *ErrorUserInformationFailedValidation) HTTPCode() int {
	return http.StatusBadRequest
}
