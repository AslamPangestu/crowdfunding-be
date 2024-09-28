package helper

import (
	"gorm.io/gorm"
)

// Pagination : Mapping Role DB
type Pagination struct {
	Total       int  `json:"total"`
	PerPage     int  `json:"per_page"`
	CurrentPage int  `json:"current_page"`
	LastPage    int  `json:"last_page"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}

// ResponsePagination : Mapping Role DB
type ResponsePagination struct {
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}

func PaginationScope(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginationAdapter(page, pageSize, total int, data interface{}) ResponsePagination {
	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	//Count Last Page
	var lastPage, remaining int
	remaining = total % pageSize
	if remaining != 0 {
		lastPage = total/pageSize + 1
	} else {
		lastPage = total / pageSize
	}

	return ResponsePagination{
		Data: data,
		Pagination: Pagination{
			Total:       total,
			PerPage:     pageSize,
			CurrentPage: page,
			LastPage:    lastPage,
			HasPrev:     page > 1,
			HasNext:     page < lastPage && page >= 1,
		},
	}
}
