package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// UserRepository Contract
type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	// Update(user entity.User) (entity.User, error)
	// View(user entity.User) (entity.User, error)
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
