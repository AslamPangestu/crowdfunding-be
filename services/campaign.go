package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
)

// CampaignService Contract
type CampaignService interface {
	GetCampaigns(userID int) ([]entity.Campaign, error)
	// Create(form entity.RoleRequest) (entity.Role, error)
	// Search(form entity.RoleRequest) (entity.Role, error)
	// Remove(form entity.RoleRequest) (entity.Role, error)
}

type campaignService struct {
	repository repository.CampaignRepository
}

// CampaignServiceInit Initiation
func CampaignServiceInit(repository repository.CampaignRepository) *campaignService {
	return &campaignService{repository}
}

func (s *campaignService) GetCampaigns(userID int) ([]entity.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindManyByCampaignerID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
