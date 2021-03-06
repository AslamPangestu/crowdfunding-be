package entity

import "time"

//Role : Mapping Role DB
type Role struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//RolesRequest : Mapping Role Request
type RolesRequest struct {
	Name string `json:"name" binding:"required"`
}

//RolesResponse : Mapping Role Response
type RolesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
