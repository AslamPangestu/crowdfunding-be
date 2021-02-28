package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
)

// RoleService Contract
type RoleService interface {
	Create(form entity.RoleRequest) (entity.Role, error)
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

func (s *roleService) Create(form entity.RoleRequest) (entity.Role, error) {
	role := entity.Role{}
	role.Name = form.Name

	newRole, err := s.repository.Create(role)
	if err != nil {
		return newRole, err
	}
	return newRole, nil
}
