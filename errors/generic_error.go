package errors

import "github.com/gin-gonic/gin"

//GenericError interface for all BantuPay API errors
type GenericError interface {
	ErrorType() string
	Message() string
	Data() string
	JSONError() gin.H
	HTTPCode() int
}
