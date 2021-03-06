package repository

import (
	"crowdfunding/entity"
	"fmt"

	"gorm.io/gorm"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	Create(campaign entity.Campaign) (entity.Campaign, error)
	FindAll() ([]entity.Campaign, error)
	FindByID(id int) (entity.Campaign, error)
	FindManyByCampaignerID(userID int) ([]entity.Campaign, error)
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

func (r *campaignRepository) Create(campaign entity.Campaign) (entity.Campaign, error) {
	fmt.Println(campaign)
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *campaignRepository) FindAll() ([]entity.Campaign, error) {
	var model []entity.Campaign
	err := r.db.Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) FindByID(id int) (entity.Campaign, error) {
	var model entity.Campaign
	err := r.db.Preload("User").Preload(TBL_CAMPAIGN_IMAGES).Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) FindManyByCampaignerID(userID int) ([]entity.Campaign, error) {
	var model []entity.Campaign
	err := r.db.Where("campaigner_id = ?", userID).Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
