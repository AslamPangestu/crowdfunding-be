package services

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/repository"
)

// RoleInteractor Contract
type RoleInteractor interface {
	//Get Many
	GetRoles(page int, pageSize int) (helper.ResponsePagination, error)
	//Get One
	GetRoleByID(id int) (entity.Role, error)
	//Action
	AddRole(form entity.FormRoleRequest) (entity.Role, error)
	EditRole(id int, form entity.FormRoleRequest) (entity.Role, error)
	// RemoveRole(form entity.FormRoleRequest) (entity.Role, error)
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

func (s *roleService) GetRoleByID(id int) (entity.Role, error) {
	model, err := s.repository.FindOneByID(id)
	if err != nil {
		return model, err
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

func (s *roleService) EditRole(id int, form entity.FormRoleRequest) (entity.Role, error) {
	model, err := s.repository.FindOneByID(id)
	if err != nil {
		return model, err
	}
	model.Name = form.Name
	updatedData, err := s.repository.Update(model)
	if err != nil {
		return updatedData, err
	}
	return updatedData, nil
}
