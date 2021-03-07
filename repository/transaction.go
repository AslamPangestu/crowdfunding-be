package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	Create(model entity.Transaction) (entity.Transaction, error)
	FindManyByCampaignID(campaignID int) ([]entity.Transaction, error)
	FindManyByUserID(userID int) ([]entity.Transaction, error)
}

type trasactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository Initiation
func NewTransactionRepository(db *gorm.DB) *trasactionRepository {
	return &trasactionRepository{db}
}

const ORDER_BY_ID_DESC = "id desc"

func (r *trasactionRepository) Create(model entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *trasactionRepository) FindManyByCampaignID(campaignID int) ([]entity.Transaction, error) {
	var models []entity.Transaction
	err := r.db.Preload("User").Find(&models).Where("campaign_id = ?", campaignID).Order(ORDER_BY_ID_DESC).Error
	if err != nil {
		return models, err
	}
	return models, nil
}

func (r *trasactionRepository) FindManyByUserID(userID int) ([]entity.Transaction, error) {
	var models []entity.Transaction
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Find(&models).Where("user_id = ?", userID).Order(ORDER_BY_ID_DESC).Error
	if err != nil {
		return models, err
	}
	return models, nil
}
