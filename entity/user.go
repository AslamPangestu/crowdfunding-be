package entity

import "time"

//User : Mapping User DB
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

//REGISTER ENTITY

//RegisterRequest : Mapping Register Request
type RegisterRequest struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Occupation string `json:"occupation"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

//RegisterResponse : Mapping Register Response
type RegisterResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

//RegsiterAdapter : Adapter Register
func RegsiterAdapter(user User, token string) RegisterResponse {
	res := RegisterResponse{
		ID:         user.ID,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}
	return res
}

//LOGIN ENTITY

//LoginRequest : Mapping Login Request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//LoginResponse : Mapping Register Response
type LoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

//LoginAdapter : Adapter Register
func LoginAdapter(user User, token string) LoginResponse {
	res := LoginResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	return res
}

//EMAIL VALIDATION ENTITY

//EmailValidationRequest : Mapping Login Request
type EmailValidationRequest struct {
	Email string `json:"email" binding:"required,email"`
}
