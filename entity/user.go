package entity

import "time"

// User : Mapping User DB
type User struct {
	ID           int
	Name         string
	Occupation   string
	Username     string
	Email        string
	PasswordHash string
	AvatarPath   string
	Token        string
	RoleID       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// RegisterRequest : Mapping Register Request
type RegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Occupation string `json:"occupation"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// RegisterResponse : Mapping Register Response
type RegisterResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

// LoginRequest : Mapping Login Request
type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginResponse : Mapping Register Response
type LoginResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	ImageURL   string `json:"image_url"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

// UpdateUserRequest : Mapping Update User Request
type UpdateUserRequest struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Occupation string `json:"occupation" binding:"required"`
}

// EmailValidationRequest : Mapping Email Validation Request
type EmailValidationRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// CreateUserForm : Mapping Form Create User
type CreateUserForm struct {
	Name       string `form:"name" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
	User       User
}

// EditUserForm : Mapping Form Create User
type EditUserForm struct {
	ID         int
	Name       string `form:"name" binding:"required"`
	Username   string `form:"username" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Error      error
	User       User
}
