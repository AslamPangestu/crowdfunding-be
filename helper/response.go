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

// PaginationAdapterResponse : Mapping PaginationAdapterResponse Response
type paginationAdapterResponse struct {
	Total       int
	PerPage     int
	CurrentPage int
	Pages       []int
	LastPage    int
	NextPage    int
	PrevPage    int
	HasNext     bool
	HasPrev     bool
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

// ErrResponseValidationHandler : Handler Error validation response
func ErrResponseValidationHandler(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

// PaginationAdapterHandler : Handler Error validation response
func PaginationAdapterHandler(pagination Pagination) paginationAdapterResponse {
	var pages []int
	for i := 1; i <= pagination.LastPage; i++ {
		pages = append(pages, i)
	}
	return paginationAdapterResponse{
		Total:       pagination.Total,
		PerPage:     pagination.PerPage,
		CurrentPage: pagination.CurrentPage,
		LastPage:    pagination.LastPage,
		NextPage:    pagination.CurrentPage + 1,
		PrevPage:    pagination.CurrentPage - 1,
		HasNext:     pagination.HasNext,
		HasPrev:     pagination.HasPrev,
		Pages:       pages,
	}
}
