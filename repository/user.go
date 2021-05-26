package repository

import (
	"crowdfunding/entity"
	"crowdfunding/helper"

	"gorm.io/gorm"
)

// UserInteractor Contract
type UserInteractor interface {
	//Get Many
	FindAll(query entity.Paginate) (helper.ResponsePagination, error)
	//Get One
	FindOneByID(id int) (entity.User, error)
	FindOneByEmail(email string) (entity.User, error)
	//Action
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	// Delete(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository Initiation
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

const TABLE_USERS = "users"

//Get Many
func (r *userRepository) FindAll(query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.User
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Scopes(helper.PaginationScope(query.Page, query.PageSize)).Where("role_id != 1").Find(&models).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_USERS).Where("role_id != 1").Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

//Get One
func (r *userRepository) FindOneByID(id int) (entity.User, error) {
	var model entity.User
	err := r.db.Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) FindOneByEmail(email string) (entity.User, error) {
	var model entity.User
	err := r.db.Where("email = ?", email).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

//Action
func (r *userRepository) Create(model entity.User) (entity.User, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) Update(model entity.User) (entity.User, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
