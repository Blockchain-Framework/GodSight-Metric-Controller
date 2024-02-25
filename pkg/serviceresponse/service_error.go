package serviceresponse

import (
	"fmt"
	"net/http"
)

type ServiceError struct {
	Code    int
	Message string
	Detail  string
}

func New() *ServiceError {
	return &ServiceError{
		Code:    0,
		Message: "",
		Detail:  "",
	}
}

func (se *ServiceError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", se.Code, se.Message)
}

func (se *ServiceError) DecodeHttpCode(statusCode int, message string) {

	if v := http.StatusText(statusCode); len(v) != 0 {
		se.Code = statusCode
		se.Message = v
		se.Detail = message
	} else {
		se.Code = 0
		se.Message = "unknown http code"
		se.Detail = message
	}
}
