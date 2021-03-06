package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
	"fmt"

	"github.com/gosimple/slug"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	CreateCampaign(form entity.CreateCampaignRequest) (entity.Campaign, error)
	GetCampaigns(userID int) ([]entity.Campaign, error)
	GetCampaignByID(request entity.CampaignDetailRequest) (entity.Campaign, error)
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

func (s *campaignService) CreateCampaign(form entity.CreateCampaignRequest) (entity.Campaign, error) {
	slugString := fmt.Sprintf("%d %s", form.CampaignerID, form.Title)
	campaign := entity.Campaign{
		Title:            form.Title,
		ShortDescription: form.ShortDescription,
		Description:      form.Description,
		TargetAmount:     form.TargetAmount,
		Perks:            form.Perks,
		Slug:             slug.Make(slugString),
		CampaignerID:     form.CampaignerID,
	}
	newCampaign, err := s.repository.Create(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil

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
