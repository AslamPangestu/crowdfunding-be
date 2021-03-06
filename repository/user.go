package repository

import (
	"crowdfunding/entity"
	"fmt"

	"gorm.io/gorm"
)

// UserInteractor Contract
type UserInteractor interface {
	Create(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByID(id int) (entity.User, error)
	FindOneBy(key string, value string) (entity.User, error)
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

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var model []entity.User
	err := r.db.Find(&model).Where("role_id != 1").Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) FindByID(id int) (entity.User, error) {
	var model entity.User
	err := r.db.Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) FindOneBy(key string, value string) (entity.User, error) {
	var model entity.User
	query := fmt.Sprintf("%s = ?", key)
	err := r.db.Where(query, value).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
