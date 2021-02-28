package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
)

// RoleService Contract
type RoleService interface {
	AddRole(form entity.RoleRequest) (entity.Role, error)
	EditRole(id int, form entity.RoleRequest) (entity.Role, error)
	GetRoles() ([]entity.Role, error)
	// Search(form entity.RoleRequest) (entity.Role, error)
	// Remove(form entity.RoleRequest) (entity.Role, error)
}

type roleService struct {
	repository repository.RoleRepository
}

// RoleServiceInit Initiation
func RoleServiceInit(repository repository.RoleRepository) *roleService {
	return &roleService{repository}
}

func (s *roleService) AddRole(form entity.RoleRequest) (entity.Role, error) {
	role := entity.Role{}
	role.Name = form.Name

	newRole, err := s.repository.Create(role)
	if err != nil {
		return newRole, err
	}
	return newRole, nil
}

func (s *roleService) GetRoles() ([]entity.Role, error) {
	roles, err := s.repository.FindAll()
	if err != nil {
		return roles, err
	}
	return roles, nil
}

func (s *roleService) EditRole(id int, form entity.RoleRequest) (entity.Role, error) {
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
