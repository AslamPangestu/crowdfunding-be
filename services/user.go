package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService Contract
type UserService interface {
	Register(input entity.RegisterRequest) (entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

// UserServiceInit Initiation
func UserServiceInit(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Register(input entity.RegisterRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.RoleID = 1

	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
