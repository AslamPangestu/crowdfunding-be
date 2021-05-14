package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
)

// RoleInteractor Contract
type RoleInteractor interface {
	AddRole(form entity.FormRoleRequest) (entity.Role, error)
	GetRoles() ([]entity.Role, error)
	GetRoleByID(id int) (entity.Role, error)
	GetRolesByName(form entity.RoleNameRequest) ([]entity.Role, error)
	EditRole(id int, form entity.FormRoleRequest) (entity.Role, error)
	// RemoveRoles(form entity.FormRoleRequest) (entity.Role, error)
}

type roleService struct {
	repository repository.RoleInteractor
}

// NewRoleService Initiation
func NewRoleService(repository repository.RoleInteractor) *roleService {
	return &roleService{repository}
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

func (s *roleService) GetRoles() ([]entity.Role, error) {
	models, err := s.repository.FindAll()
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

func (s *roleService) GetRolesByName(form entity.RoleNameRequest) ([]entity.Role, error) {
	models, err := s.repository.FindManyByName(form.Name)
	if err != nil {
		return models, err
	}
	return models, nil
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
