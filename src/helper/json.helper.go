package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta meta `json:"meta"`
	Data any  `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func APIResponse(code int, message, status string, data any) Response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
