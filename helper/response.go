package helper

import (
	"github.com/go-playground/validator/v10"
)

// Repsonse : Mapping Response
type Repsonse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta : Mapping Meta Response
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// ResponseHandler : Handler response
func ResponseHandler(message string, code int, status string, data interface{}) Repsonse {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	res := Repsonse{
		Meta: meta,
		Data: data,
	}
	return res
}

func ErrResponseValidationHandler(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
