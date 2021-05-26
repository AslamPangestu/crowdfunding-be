package services

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserInteractor Contract
type UserInteractor interface {
	//Get Many
	GetAllUsers(page int, pageSize int) (helper.ResponsePagination, error)
	//Get One
	GetUserByID(id int) (entity.User, error)
	IsEmailAvaiable(form entity.EmailValidationRequest) (bool, error)
	//Action
	Register(form entity.RegisterRequest) (entity.User, error)
	Login(form entity.LoginRequest) (entity.User, error)
	UploadAvatar(id int, fileLocation string) (entity.User, error)
	UpdateUser(form entity.EditUserForm) (entity.User, error)
}

type userService struct {
	repository repository.UserInteractor
}

// NewUserService Initiation
func NewUserService(repository repository.UserInteractor) *userService {
	return &userService{repository}
}

//Get Many
func (s *userService) GetAllUsers(page int, pageSize int) (helper.ResponsePagination, error) {
	//Find
	request := entity.Paginate{Page: page, PageSize: pageSize}
	model, err := s.repository.FindAll(request)
	if err != nil {
		return model, err
	}
	return model, nil
}

//Get One
func (s *userService) GetUserByID(id int) (entity.User, error) {
	//Find
	model, err := s.repository.FindOneByID(id)
	if err != nil {
		return model, err
	}
	//Is Found?
	if model.ID == 0 {
		return model, errors.New("User not found")
	}
	return model, nil
}

func (s *userService) IsEmailAvaiable(form entity.EmailValidationRequest) (bool, error) {
	//Find
	model, err := s.repository.FindOneByEmail(form.Email)
	if err != nil {
		return false, err
	}
	//Is Available
	if model.ID != 0 {
		return false, nil
	}
	return true, nil
}

//Action
func (s *userService) Register(form entity.RegisterRequest) (entity.User, error) {
	var model entity.User
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.MinCost)
	if err != nil {
		return model, err
	}
	model = entity.User{
		Name:         form.Name,
		Username:     form.Username,
		Email:        form.Email,
		Occupation:   form.Occupation,
		RoleID:       2,
		PasswordHash: string(passwordHash),
	}

	newUser, err := s.repository.Create(model)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userService) Login(form entity.LoginRequest) (entity.User, error) {
	//Find User
	model, err := s.repository.FindOneByEmail(form.Email)
	if err != nil {
		return model, err
	}
	//Is Found?
	if model.ID == 0 {
		return model, errors.New("User not found")
	}
	//Decrypt Password Hash
	err = bcrypt.CompareHashAndPassword([]byte(model.PasswordHash), []byte(form.Password))
	if err != nil {
		return model, errors.New("Password incorrect")
	}
	return model, nil
}

func (s *userService) UploadAvatar(id int, fileLocation string) (entity.User, error) {
	//Find
	model, err := s.repository.FindOneByID(id)
	if err != nil {
		return model, err
	}
	//Update Path
	model.AvatarPath = fileLocation
	//Update DB
	updatedData, err := s.repository.Update(model)
	if err != nil {
		return updatedData, err
	}
	return updatedData, nil
}

func (s *userService) UpdateUser(form entity.EditUserForm) (entity.User, error) {
	model, err := s.repository.FindOneByID(form.ID)
	if err != nil {
		return model, err
	}
	model.Name = form.Name
	model.Username = form.Username
	model.Email = form.Email
	model.Occupation = form.Occupation

	updatedData, err := s.repository.Update(model)
	if err != nil {
		return updatedData, err
	}
	return updatedData, nil
}
