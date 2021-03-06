package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	GetCampaigns(userID int) ([]entity.Campaign, error)
	GetCampaignByID(request entity.CampaignDetailRequest) (entity.Campaign, error)
	// Create(form entity.RolesRequest) (entity.Role, error)
	// Search(form entity.RolesRequest) (entity.Role, error)
	// Remove(form entity.RolesRequest) (entity.Role, error)
}

type campaignService struct {
	repository repository.CampaignInteractor
}

// CampaignServiceInit Initiation
func CampaignServiceInit(repository repository.CampaignInteractor) *campaignService {
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

func (s *campaignService) GetCampaignByID(request entity.CampaignDetailRequest) (entity.Campaign, error) {
	campaign, err := s.repository.FindByID(request.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
