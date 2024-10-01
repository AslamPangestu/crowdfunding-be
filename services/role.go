package services

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/repository"
	"errors"
)

// RoleInteractor Contract
type RoleInteractor interface {
	//Get Many
	GetRoles(page int, pageSize int) (helper.ResponsePagination, error)
	//Get One
	GetRoleByID(id string) (entity.Role, error)
	//Action
	AddRole(form entity.FormRoleRequest) (entity.Role, error)
	EditRole(id string, form entity.FormRoleRequest) (entity.Role, error)
	RemoveRole(id string) error
}

type roleService struct {
	repository repository.RoleInteractor
}

// NewRoleService Initiation
func NewRoleService(repository repository.RoleInteractor) *roleService {
	return &roleService{repository}
}

func (s *roleService) GetRoles(page int, pageSize int) (helper.ResponsePagination, error) {
	request := entity.Paginate{Page: page, PageSize: pageSize}
	models, err := s.repository.FindAll(request)
	if err != nil {
		return models, err
	}
	return models, nil
}

func (s *roleService) GetRoleByID(id string) (entity.Role, error) {
	model, err := s.repository.FindOneByID(id)
	if err != nil {
		return model, err
	}
	if model.ID == "" {
		return model, errors.New("ROLE NOT FOUND")
	}
	return model, nil
}

func (s *roleService) AddRole(form entity.FormRoleRequest) (entity.Role, error) {
	model := entity.Role{
		Name: form.Name,
	}

	newRole, err := s.repository.Create(model)
	if err != nil {
		return newRole, err
	}
	return newRole, nil
}

func (s *roleService) EditRole(id string, form entity.FormRoleRequest) (entity.Role, error) {
	model, err := s.repository.FindOneByID(id)
	if err != nil {
		return model, err
	}
	if model.ID == "" {
		return model, errors.New("ROLE NOT FOUND")
	}
	model.Name = form.Name
	updatedData, err := s.repository.Update(model)
	if err != nil {
		return updatedData, err
	}
	return updatedData, nil
}

func (s *roleService) RemoveRole(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
