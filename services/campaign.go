package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	CreateCampaign(form entity.FormCampaignRequest) (entity.Campaign, error)
	GetCampaigns(userID int) ([]entity.Campaign, error)
	GetCampaignByID(uri entity.CampaignIDRequest) (entity.Campaign, error)
	EditCampaign(uri entity.CampaignIDRequest, form entity.FormCampaignRequest) (entity.Campaign, error)
	UploadCampaignImages(form entity.UploadCampaignImageRequest, fileLocation string) (entity.CampaignImage, error)
	// Search(form entity.FormRoleRequest) (entity.Role, error)
	// Remove(form entity.FormRoleRequest) (entity.Role, error)
}

type campaignService struct {
	repository repository.CampaignInteractor
}

// NewCampaignService Initiation
func NewCampaignService(repository repository.CampaignInteractor) *campaignService {
	return &campaignService{repository}
}

func (s *campaignService) CreateCampaign(form entity.FormCampaignRequest) (entity.Campaign, error) {
	slugString := fmt.Sprintf("%d %s", form.CampaignerID, form.Title)
	model := entity.Campaign{
		Title:            form.Title,
		ShortDescription: form.ShortDescription,
		Description:      form.Description,
		TargetAmount:     form.TargetAmount,
		Perks:            form.Perks,
		Slug:             slug.Make(slugString),
		CampaignerID:     form.CampaignerID,
	}
	newCampaign, err := s.repository.Create(model)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}

func (s *campaignService) GetCampaigns(userID int) ([]entity.Campaign, error) {
	if userID != 0 {
		models, err := s.repository.FindManyByCampaignerID(userID)
		if err != nil {
			return models, err
		}
		return models, nil
	}
	models, err := s.repository.FindAll()
	if err != nil {
		return models, err
	}
	return models, nil
}

func (s *campaignService) GetCampaignByID(uri entity.CampaignIDRequest) (entity.Campaign, error) {
	model, err := s.repository.FindOneByID(uri.ID)
	if err != nil {
		return model, err
	}
	return model, nil
}

func (s *campaignService) EditCampaign(uri entity.CampaignIDRequest, form entity.FormCampaignRequest) (entity.Campaign, error) {
	model, err := s.repository.FindOneByID(uri.ID)
	if err != nil {
		return model, err
	}

	if model.CampaignerID != form.CampaignerID {
		return model, errors.New("User not Authorize for this action")
	}
	model.Title = form.Title
	model.ShortDescription = form.ShortDescription
	model.Description = form.Description
	model.Perks = form.Perks
	model.TargetAmount = form.TargetAmount

	updateCampaign, err := s.repository.Update(model)
	if err != nil {
		return updateCampaign, err
	}
	return updateCampaign, nil
}

func (s *campaignService) UploadCampaignImages(form entity.UploadCampaignImageRequest, fileLocation string) (entity.CampaignImage, error) {
	campaign, err := s.repository.FindOneByID(form.CampaignID)
	if err != nil {
		return entity.CampaignImage{}, err
	}
	if campaign.CampaignerID != form.UserID {
		return entity.CampaignImage{}, errors.New("User not Authorize for this action")
	}
	isPrimary := 0

	if form.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(form.CampaignID)
		if err != nil {
			return entity.CampaignImage{}, err
		}
	}

	model := entity.CampaignImage{
		CampaignID: form.CampaignID,
		IsPrimary:  isPrimary,
		ImagePath:  fileLocation,
	}

	newCampaignImage, err := s.repository.CreateImage(model)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil

}
