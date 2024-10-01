package repository

import (
	"crowdfunding/entity"
	"crowdfunding/helper"

	"gorm.io/gorm"
)

// RoleInteractor Contract
type RoleInteractor interface {
	//Get Many
	FindAll(query entity.Paginate) (helper.ResponsePagination, error)
	//Get One
	FindOneByID(id string) (entity.Role, error)
	FindOneByName(name string) (entity.Role, error)
	//Action
	Create(model entity.Role) (entity.Role, error)
	Update(model entity.Role) (entity.Role, error)
	Delete(id string) error
}

type roleRepo struct {
	db *gorm.DB
}

// NewRoleRepository Initiation
func NewRoleRepository(db *gorm.DB) *roleRepo {
	return &roleRepo{db}
}

const TABLE_ROLES = "roles"

// Get Many
func (r *roleRepo) FindAll(query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.Role
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Scopes(helper.PaginationScope(query.Page, query.PageSize)).Find(&models).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_ROLES).Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

// Get One
func (r *roleRepo) FindOneByID(id string) (entity.Role, error) {
	var model entity.Role
	err := r.db.Find(&model).Where("xata_id = ?", id).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepo) FindOneByName(name string) (entity.Role, error) {
	var model entity.Role
	err := r.db.Where("name = ?", name).First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

// Action
func (r *roleRepo) Create(model entity.Role) (entity.Role, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepo) Update(model entity.Role) (entity.Role, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *roleRepo) Delete(id string) error {
	err := r.db.Delete(&entity.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
