package repository

import (
	"crowdfunding/entity"
	"fmt"

	"gorm.io/gorm"
)

// CampaignRepository Contract
type CampaignRepository interface {
	FindAll() ([]entity.Campaign, error)
	FindManyBy(key string, value string) ([]entity.Campaign, error)
	FindByID(id int) (entity.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

const (
	TBL_CAMPAIGN_IMAGES   = "CampaignImages"
	QUERY_CAMPAIGN_IMAGES = "campaign_images.is_primary = 1"
)

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
	err := r.db.Where("id = ?", id).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) FindManyBy(key string, value string) ([]entity.Campaign, error) {
	var model []entity.Campaign
	query := fmt.Sprintf("%s = ?", key)
	err := r.db.Where(query, value).Preload(TBL_CAMPAIGN_IMAGES, QUERY_CAMPAIGN_IMAGES).Find(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
