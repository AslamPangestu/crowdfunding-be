package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// RoleRepository Contract
type RoleRepository interface {
	Create(role entity.Role) (entity.Role, error)
	// Update(role entity.Role) (entity.Role, error)
	// View(role entity.Role) (entity.Role, error)
	// Delete(role entity.Role) (entity.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

// RoleRepositoryInit Initiation
func RoleRepositoryInit(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(role entity.Role) (entity.Role, error) {
	err := r.db.Create(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}
