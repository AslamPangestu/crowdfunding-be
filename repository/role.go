package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// RoleRepository Contract
type RoleRepository interface {
	Create(role entity.Role) (entity.Role, error)
	FindAll() ([]entity.Role, error)
	FindOneByName(name string) (entity.Role, error)
	FindOneByID(id int) (entity.Role, error)
	// Update(role entity.Role) (entity.Role, error)
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

func (r *roleRepository) FindAll() ([]entity.Role, error) {
	var model []entity.Role
	err := r.db.Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepository) FindOneByName(name string) (entity.Role, error) {
	var model entity.Role
	err := r.db.Where("name = ?", name).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepository) FindOneByID(id int) (entity.Role, error) {
	var model entity.Role
	err := r.db.Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
