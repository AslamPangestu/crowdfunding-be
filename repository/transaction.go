package repository

import (
	"crowdfunding/entity"
	"crowdfunding/helper"

	"gorm.io/gorm"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	//Get Many
	FindAll(query entity.Paginate) (helper.ResponsePagination, error)
	FindManyByCampaignID(campaignID int, query entity.Paginate) (helper.ResponsePagination, error)
	FindManyByUserID(userID int, query entity.Paginate) (helper.ResponsePagination, error)
	//Get One
	FindOneByTransactionID(transactionID int) (entity.Transaction, error)
	FindOneByTrxCode(trxCode string) (entity.Transaction, error)
	//Action
	Create(model entity.Transaction) (entity.Transaction, error)
	Update(model entity.Transaction) (entity.Transaction, error)
}

type trasactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository Initiation
func NewTransactionRepository(db *gorm.DB) *trasactionRepository {
	return &trasactionRepository{db}
}

const (
	TABLE_TRANSACTIONS = "transactions"
	ORDER_BY_ID_DESC   = "id desc"
)

//Get Many
func (r *trasactionRepository) FindAll(query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.Transaction
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Scopes(helper.PaginationScope(query.Page, query.PageSize)).Order("created_at desc").Preload("Campaign").Find(&models).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_TRANSACTIONS).Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

func (r *trasactionRepository) FindManyByCampaignID(campaignID int, query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.Transaction
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Preload("User").Find(&models).Where("campaign_id = ?", campaignID).Order(ORDER_BY_ID_DESC).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_TRANSACTIONS).Where("campaign_id = ?", campaignID).Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

func (r *trasactionRepository) FindManyByUserID(userID int, query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.Transaction
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Scopes(helper.PaginationScope(query.Page, query.PageSize)).Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Find(&models).Where("user_id = ?", userID).Order(ORDER_BY_ID_DESC).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_TRANSACTIONS).Where("user_id = ?", userID).Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

//Get One
func (r *trasactionRepository) FindOneByTransactionID(transactionID int) (entity.Transaction, error) {
	var model entity.Transaction
	err := r.db.Where("id = ?", transactionID).Find(&model).Where("id = ?", transactionID).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
func (r *trasactionRepository) FindOneByTrxCode(trxCode string) (entity.Transaction, error) {
	var model entity.Transaction
	err := r.db.Where("trx_code = ?", trxCode).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

//Action
func (r *trasactionRepository) Create(model entity.Transaction) (entity.Transaction, error) {
	err := r.db.Create(&model).Error
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
