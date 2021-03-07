package repository

import (
	"crowdfunding/entity"

	"gorm.io/gorm"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	FindManyByCampaignID(campaignID int) ([]entity.Transaction, error)
}

type trasactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository Initiation
func NewTransactionRepository(db *gorm.DB) *trasactionRepository {
	return &trasactionRepository{db}
}

func (r *trasactionRepository) FindManyByCampaignID(campaignID int) ([]entity.Transaction, error) {
	var model []entity.Transaction
	err := r.db.Preload("User").Find(&model).Where("campaign_id = ?", campaignID).Order("id desc").Error
	if err != nil {
		return model, err
	}
	return model, nil
}
