package repository

import (
	"crowdfunding/entity"
	"fmt"

	"gorm.io/gorm"
)

// RoleRepository Contract
type RoleRepository interface {
	Create(role entity.Role) (entity.Role, error)
	FindBy(key string, value string) (entity.Role, error)
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

func (r *roleRepository) FindBy(key string, value string) (entity.Role, error) {
	var model entity.Role
	query := fmt.Sprintf("%s = ?", key)
	err := r.db.Where(query, value).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
