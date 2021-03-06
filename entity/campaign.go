package entity

import (
	"strings"
	"time"
)

//Campaign : Mapping Campaign DB
type Campaign struct {
	ID               int
	CampaignerID     int
	Title            string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	TargetAmount     int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             User
}

//CampaignImage : Mapping CampaignImage DB
type CampaignImage struct {
	ID         int
	CampaignID int
	ImagePath  string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//CampaignResponse : Mapping Campaign Response
type CampaignResponse struct {
	ID               int    `json:"id"`
	CampaignerID     int    `json:"campaigner_id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	TargetAmount     int    `json:"target_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	ImageURL         string `json:"image_url"`
}

//CampaignAdapter : Adapter Campaign
func CampaignAdapter(campaign Campaign) CampaignResponse {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].ImagePath
	}
	res := CampaignResponse{
		ID:               campaign.ID,
		Title:            campaign.Title,
		CampaignerID:     campaign.CampaignerID,
		ShortDescription: campaign.ShortDescription,
		TargetAmount:     campaign.TargetAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         imageURL,
	}
	return res
}

//CampaignsAdapter : Adapter Campaigns
func CampaignsAdapter(campaigns []Campaign) []CampaignResponse {
	campaignsAdapter := []CampaignResponse{}
	for _, campaign := range campaigns {
		campaignAdapter := CampaignAdapter(campaign)
		campaignsAdapter = append(campaignsAdapter, campaignAdapter)
	}
	return campaignsAdapter
}

//CampaignDetailRequest : Request Detail Campaign
type CampaignDetailRequest struct {
	ID int `uri:"id" binding:"required"`
}

//CampaignDetailResponse : Response Detail Campaign
type CampaignDetailResponse struct {
	ID               int                   `json:"id"`
	Title            string                `json:"title"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	TargetAmount     int                   `json:"target_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	Slug             string                `json:"slug"`
	ImageURL         string                `json:"image_url"`
	Perks            []string              `json:"perks"`
	User             UserCampaignDetail    `json:"user"`
	Images           []ImageCampaignDetail `json:"images"`
}

//UserCampaignDetail : Response Detail Campaign for User
type UserCampaignDetail struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

//ImageCampaignDetail : Response Detail Campaign for Image
type ImageCampaignDetail struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

//CampaignDetailAdapter : Adapter Campaign Detail
func CampaignDetailAdapter(campaign Campaign) CampaignDetailResponse {
	imageURL := ""
	var perks []string
	images := []ImageCampaignDetail{}
	//SET User
	user := UserCampaignDetail{
		Name:     campaign.User.Name,
		ImageURL: campaign.User.AvatarPath,
	}
	//SET Image URL
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].ImagePath
	}
	//SET Perks
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	//SET Images
	for _, image := range campaign.CampaignImages {
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		images = append(images, ImageCampaignDetail{
			ImageURL:  image.ImagePath,
			IsPrimary: isPrimary,
		})

	}
	campaignResponse := CampaignDetailResponse{
		ID:               campaign.ID,
		Title:            campaign.Title,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		TargetAmount:     campaign.TargetAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         imageURL,
		Perks:            perks,
		User:             user,
	}
	return campaignResponse
}
