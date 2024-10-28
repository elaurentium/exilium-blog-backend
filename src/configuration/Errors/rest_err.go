package rest_err

import "net/http"

type RestErr struct {
	Message string `JSON:"message"`
	Error string `JSON:"error"`
	Code int `JSON:"code"`
	Causes []Causes `JSON:"causes"`
}

type Causes struct {
	Field string `JSON:"field"`
	Message string `JSON:"message"`
}

func (r *RestErr) Err() string {
	return r.Message
}

func NewRestErr(message, error string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Error: error,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error: "Bad Request",
		Code: http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:	"Unauthorized",
		Code: http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Error: "Bad Request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error: "Internal Server Error",
		Code: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error: "Not Found",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error: "Forbidden",
		Code: http.StatusForbidden,
	}
}