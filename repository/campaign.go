package repository

import (
	"crowdfunding/entity"
	"crowdfunding/helper"

	"gorm.io/gorm"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	//Get Many
	FindAll(query entity.Paginate) (helper.ResponsePagination, error)
	FindManyByCampaignerID(userID int, query entity.Paginate) (helper.ResponsePagination, error)
	//Get One
	FindOneByID(id int) (entity.Campaign, error)
	//Action
	Create(model entity.Campaign) (entity.Campaign, error)
	Update(model entity.Campaign) (entity.Campaign, error)
	// Delete(id int) (entity.Role, error)
	//Action Image
	CreateImage(model entity.CampaignImage) (entity.CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}

type campaignRepository struct {
	db *gorm.DB
}

// NewCampaignRepository Initiation
func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db}
}

const (
	TABLE_CAMPAIGNS       = "campaigns"
	TBL_CAMPAIGN_IMAGES   = "CampaignImages"
	QUERY_CAMPAIGN_IMAGES = "campaign_images.is_primary = 1"
)

//Get Many
func (r *campaignRepository) FindAll(query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.Campaign
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Scopes(helper.PaginationScope(query.Page, query.PageSize)).Order("created_at desc").Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&models).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_CAMPAIGNS).Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

func (r *campaignRepository) FindManyByCampaignerID(userID int, query entity.Paginate) (helper.ResponsePagination, error) {
	var models []entity.Campaign
	var pagination helper.ResponsePagination
	var total int64
	err := r.db.Where("campaigner_id = ?", userID).Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&models).Error
	if err != nil {
		return pagination, err
	}
	r.db.Table(TABLE_CAMPAIGNS).Where("campaigner_id = ?", userID).Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Count(&total)
	pagination = helper.PaginationAdapter(query.Page, query.PageSize, int(total), models)
	return pagination, nil
}

//Get One
func (r *campaignRepository) FindOneByID(id int) (entity.Campaign, error) {
	var model entity.Campaign
	err := r.db.Preload("User").Preload(TBL_CAMPAIGN_IMAGES).Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

//Action
func (r *campaignRepository) Create(model entity.Campaign) (entity.Campaign, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) Update(model entity.Campaign) (entity.Campaign, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

//Action Image
func (r *campaignRepository) CreateImage(model entity.CampaignImage) (entity.CampaignImage, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&entity.CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
