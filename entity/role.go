package entity

import "time"

//Role : Mapping Role DB
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
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

//RoleAdapter : Adapter Role
func RoleAdapter(role Role) RoleResponse {
	res := RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
	return res
}
