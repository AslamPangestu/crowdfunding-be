package repository

import (
	"crowdfunding/entity"
	"crowdfunding/helper"

	"gorm.io/gorm"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	Create(model entity.Campaign) (entity.Campaign, error)
	FindAll(query entity.Paginate) ([]entity.Campaign, error)
	FindOneByID(id int) (entity.Campaign, error)
	FindManyByCampaignerID(userID int) ([]entity.Campaign, error)
	Update(model entity.Campaign) (entity.Campaign, error)
	//IMAGE
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
	TBL_CAMPAIGN_IMAGES   = "CampaignImages"
	QUERY_CAMPAIGN_IMAGES = "campaign_images.is_primary = 1"
)

func (r *campaignRepository) Create(model entity.Campaign) (entity.Campaign, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) FindAll(query entity.Paginate) ([]entity.Campaign, error) {
	var models []entity.Campaign
	err := r.db.Scopes(helper.Pagination(query.Page, query.PageSize)).Order("created_at desc").Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&models).Error
	if err != nil {
		return models, err
	}
	return models, nil
}

func (r *campaignRepository) FindOneByID(id int) (entity.Campaign, error) {
	var model entity.Campaign
	err := r.db.Preload("User").Preload(TBL_CAMPAIGN_IMAGES).Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) FindManyByCampaignerID(userID int) ([]entity.Campaign, error) {
	var models []entity.Campaign
	err := r.db.Where("campaigner_id = ?", userID).Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&models).Error
	if err != nil {
		return models, err
	}
	return models, nil
}

func (r *campaignRepository) Update(model entity.Campaign) (entity.Campaign, error) {
	err := r.db.Save(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

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
