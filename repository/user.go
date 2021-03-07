package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// UserInteractor Contract
type UserInteractor interface {
	Create(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindOneByID(id int) (entity.User, error)
	FindOneByEmail(email string) (entity.User, error)
	// FindManyByQuery(user entity.User) (entity.User, error)
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

func (r *userRepository) Create(model entity.User) (entity.User, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var models []entity.User
	err := r.db.Find(&models).Where("role_id != 1").Error
	if err != nil {
		return models, err
	}
	return models, nil
}

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

func (r *userRepository) Update(model entity.User) (entity.User, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
