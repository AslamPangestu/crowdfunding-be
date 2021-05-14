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

//FormRoleRequest : Mapping Role Request
type FormRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

//RoleResponse : Mapping Role Response
type RoleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//RoleNameRequest : Mapping Role Request by name
type CreateRoleForm struct {
	Name  string `form:"name" binding:"required"`
	Error error
}

//RoleNameRequest : Mapping Role Request by name
type EditRoleForm struct {
	ID    int
	Name  string `form:"name" binding:"required"`
	Error error
}
