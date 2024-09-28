package adapter

import (
	"crowdfunding/entity"
	"strconv"
)

// RegisterAdapter : Adapter Register
func RegisterAdapter(user entity.User, token string) entity.RegisterResponse {
	return entity.RegisterResponse{
		ID:         user.ID,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}
}

// LoginAdapter : Adapter Login
func LoginAdapter(user entity.User, token string) entity.LoginResponse {
	return entity.LoginResponse{
		ID:         strconv.Itoa(user.ID),
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		Token:      token,
		ImageURL:   user.AvatarPath,
		Occupation: user.Occupation,
	}
}
