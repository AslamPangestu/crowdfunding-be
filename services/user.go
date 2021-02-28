package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserService Contract
type UserService interface {
	Register(form entity.RegisterRequest) (entity.User, error)
	Login(form entity.LoginRequest) (entity.User, error)
	IsEmailAvaiable(form entity.EmailValidationRequest) (bool, error)
	UploadAvatar(id int, fileLocation string) (entity.User, error)
	GetUserByID(id int) (entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

// UserServiceInit Initiation
func UserServiceInit(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Register(form entity.RegisterRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = form.Name
	user.Username = form.Username
	user.Email = form.Email
	user.Occupation = form.Occupation
	user.RoleID = 2
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.MinCost)
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

func (s *userService) Login(form entity.LoginRequest) (entity.User, error) {
	//Mapping Request
	email := form.Email
	password := form.Password

	//Find
	model, err := s.repository.FindBy("email", email)
	if err != nil {
		return model, err
	}
	//Is Found?
	if model.ID == 0 {
		return model, errors.New("No user found")
	}
	//Decrypt Password Hash
	err = bcrypt.CompareHashAndPassword([]byte(model.PasswordHash), []byte(password))
	if err != nil {
		return model, errors.New("Password incorrect")
	}
	return model, nil
}

func (s *userService) IsEmailAvaiable(form entity.EmailValidationRequest) (bool, error) {
	//Mapping Request
	email := form.Email

	//Find
	model, err := s.repository.FindBy("email", email)
	if err != nil {
		return false, err
	}
	//Is Available
	if model.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *userService) UploadAvatar(id int, fileLocation string) (entity.User, error) {
	//Find
	model, err := s.repository.FindByID(id)
	if err != nil {
		return model, err
	}
	model.AvatarPath = fileLocation
	updatedData, err := s.repository.Update(model)

	if err != nil {
		return updatedData, err
	}
	return updatedData, nil
}

func (s *userService) GetUserByID(id int) (entity.User, error) {
	//Find
	model, err := s.repository.FindByID(id)
	if err != nil {
		return model, err
	}
	if model.ID == 0 {
		return model, errors.New("User not found")
	}
	return model, nil
}
