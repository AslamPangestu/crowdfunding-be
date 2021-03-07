package entity

import "time"

//Role : Mapping Role DB
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//RoleIDRequest : Mapping Role Request by ID uri
type RoleIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

//RoleNameRequest : Mapping Role Request by name
type RoleNameRequest struct {
	Name string `form:"name" binding:"required"`
}

//RoleRequest : Mapping Role Request
type RoleRequest struct {
	Name string `json:"name" binding:"required"`
}

//RoleResponse : Mapping Role Response
type RoleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
