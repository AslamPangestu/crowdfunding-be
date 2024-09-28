package services

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/repository"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

// CampaignInteractor Contract
type CampaignInteractor interface {
	//Get Many
	GetCampaigns(userID int, page int, pageSize int) (helper.ResponsePagination, error)
	//Get One
	GetCampaignByID(uri entity.CampaignIDRequest) (entity.Campaign, error)
	//Action
	CreateCampaign(form entity.FormCampaignRequest) (entity.Campaign, error)
	EditCampaign(uri entity.CampaignIDRequest, form entity.FormCampaignRequest) (entity.Campaign, error)
	UploadCampaignImages(form entity.UploadCampaignImageRequest, fileLocation string) (entity.CampaignImage, error)
	// RemoveCampaign(form entity.FormCampaignRequest) (entity.Campaign, error)
}

type campaignService struct {
	repository repository.CampaignInteractor
}

// NewCampaignService Initiation
func NewCampaignService(repository repository.CampaignInteractor) *campaignService {
	return &campaignService{repository}
}

// Get Many
func (s *campaignService) GetCampaigns(userID int, page int, pageSize int) (helper.ResponsePagination, error) {
	query := entity.Paginate{
		Page:     page,
		PageSize: pageSize,
	}
	if userID != 0 {
		models, err := s.repository.FindManyByCampaignerID(userID, query)
		if err != nil {
			return models, err
		}
		return models, nil
	}
	models, err := s.repository.FindAll(query)
	if err != nil {
		return models, err
	}
	return models, nil
}

// Get One
func (s *campaignService) GetCampaignByID(uri entity.CampaignIDRequest) (entity.Campaign, error) {
	model, err := s.repository.FindOneByID(uri.ID)
	if err != nil {
		return model, err
	}
	if model.ID == 0 {
		return model, errors.New("CAMPAIGN NOT FOUND")
	}
	return model, nil
}

// Action
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

func (s *campaignService) EditCampaign(uri entity.CampaignIDRequest, form entity.FormCampaignRequest) (entity.Campaign, error) {
	model, err := s.repository.FindOneByID(uri.ID)
	if err != nil {
		return model, err
	}
	if model.ID == 0 {
		return model, errors.New("CAMPAIGN NOT FOUND")
	}

	if model.CampaignerID != form.CampaignerID {
		return model, errors.New("USER NOT AUTHORIZE FOR THIS ACTION")
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
	if campaign.ID == 0 {
		return entity.CampaignImage{}, errors.New("CAMPAIGN NOT FOUND")
	}
	if campaign.CampaignerID != form.UserID {
		return entity.CampaignImage{}, errors.New("USER NOT AUTHORIZE FOR THIS ACTION")
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
