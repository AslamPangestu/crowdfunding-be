package repository

import (
	"crowdfunding/entity"
	"crowdfunding/helper"

	"gorm.io/gorm"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	Create(model entity.Transaction) (entity.Transaction, error)
	FindManyByCampaignID(campaignID int) ([]entity.Transaction, error)
	FindAll(query entity.Paginate) ([]entity.Transaction, error)
	FindManyByUserID(userID int, query entity.Paginate) ([]entity.Transaction, error)
	FindOneByTransactionID(transactionID int) (entity.Transaction, error)
	Update(model entity.Transaction) (entity.Transaction, error)
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
	err := r.db.Create(&model).Error
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

func (r *trasactionRepository) FindManyByUserID(userID int, query entity.Paginate) ([]entity.Transaction, error) {
	var models []entity.Transaction
	err := r.db.Scopes(helper.Pagination(query.Page, query.PageSize)).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Find(&models).Where("user_id = ?", userID).Order(ORDER_BY_ID_DESC).Error
	if err != nil {
		return models, err
	}
	return models, nil
}

func (r *trasactionRepository) FindOneByTransactionID(transactionID int) (entity.Transaction, error) {
	var model entity.Transaction
	err := r.db.Find(&model).Where("id = ?", transactionID).Order(ORDER_BY_ID_DESC).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *trasactionRepository) Update(model entity.Transaction) (entity.Transaction, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *trasactionRepository) FindAll(query entity.Paginate) ([]entity.Transaction, error) {
	var models []entity.Transaction
	err := r.db.Scopes(helper.Pagination(query.Page, query.PageSize)).Order("created_at desc").Preload("Campaign").Find(&models).Error
	if err != nil {
		return models, err
	}
	return models, nil
}
