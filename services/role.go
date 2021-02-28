package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
)

// RoleService Contract
type RoleService interface {
	Create(input entity.RoleRequest) (entity.Role, error)
}

type roleService struct {
	repository repository.RoleRepository
}

// RoleServiceInit Initiation
func RoleServiceInit(repository repository.RoleRepository) *roleService {
	return &roleService{repository}
}

func (s *roleService) Create(input entity.RoleRequest) (entity.Role, error) {
	role := entity.Role{}
	role.Name = input.Name

	newRole, err := s.repository.Create(role)
	if err != nil {
		return newRole, err
	}
	return newRole, nil
}
