package repository

import (
	"crowdfunding/entity"
	"fmt"

	"gorm.io/gorm"
)

// UserRepository Contract
type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindBy(key string, value string) (entity.User, error)
	// FindAll() ([]entity.Role, error)
	// View(role entity.Role) (entity.Role, error)
	// Update(user entity.User) (entity.User, error)
	// Delete(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// UserRepositoryInit Initiation
func UserRepositoryInit(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindBy(key string, value string) (entity.User, error) {
	var model entity.User
	query := fmt.Sprintf("%s = ?", key)
	err := r.db.Where(query, value).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
