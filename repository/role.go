package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// RoleInteractor Contract
type RoleInteractor interface {
	Create(model entity.Role) (entity.Role, error)
	FindAll() ([]entity.Role, error)
	FindOneByID(id int) (entity.Role, error)
	FindManyByName(name string) ([]entity.Role, error)
	Update(model entity.Role) (entity.Role, error)
	// Delete(id int) (entity.Role, error)
}

type roleRepo struct {
	db *gorm.DB
}

// NewRoleRepository Initiation
func NewRoleRepository(db *gorm.DB) *roleRepo {
	return &roleRepo{db}
}

func (r *roleRepo) Create(model entity.Role) (entity.Role, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepo) FindAll() ([]entity.Role, error) {
	var models []entity.Role
	err := r.db.Find(&models).Error
	if err != nil {
		return models, err
	}
	return models, nil
}

func (r *roleRepo) FindOneByID(id int) (entity.Role, error) {
	var model entity.Role
	err := r.db.Find(&model).Where("id = ?", id).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepo) FindManyByName(name string) ([]entity.Role, error) {
	var models []entity.Role
	err := r.db.Find(&models).Where("name LIKE %?%", name).Error
	if err != nil {
		return models, err
	}
	return models, nil
}

func (r *roleRepo) Update(model entity.Role) (entity.Role, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
