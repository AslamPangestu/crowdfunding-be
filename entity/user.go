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
	AvatarURL    string
	Token        string
	RoleID       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//RegisterRequest : Mapping Register Request
type RegisterRequest struct {
	Name       string `json:"name" binding:"required"`
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
