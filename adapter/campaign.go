package adapter

import (
	"crowdfunding/entity"
	"strings"
)

// CampaignsAdapter : Adapter Campaigns
func CampaignsAdapter(campaigns []entity.Campaign) []entity.CampaignResponse {
	campaignsAdapter := []entity.CampaignResponse{}
	for _, campaign := range campaigns {
		campaignAdapter := CampaignAdapter(campaign)
		campaignsAdapter = append(campaignsAdapter, campaignAdapter)
	}
	return campaignsAdapter
}

// CampaignAdapter : Adapter Campaign for Campaigns Adapter
func CampaignAdapter(campaign entity.Campaign) entity.CampaignResponse {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].ImagePath
	}
	return entity.CampaignResponse{
		ID:               campaign.ID,
		Title:            campaign.Title,
		CampaignerID:     campaign.CampaignerID,
		ShortDescription: campaign.ShortDescription,
		TargetAmount:     campaign.TargetAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         imageURL,
	}
}

// CampaignDetailAdapter : Adapter Campaign Detail
func CampaignDetailAdapter(campaign entity.Campaign) entity.CampaignDetailResponse {
	imageURL := ""
	var perks []string
	images := []entity.ImageCampaignDetail{}
	//SET User
	user := entity.UserCampaignDetail{
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
		images = append(images, entity.ImageCampaignDetail{
			ImageURL:  image.ImagePath,
			IsPrimary: isPrimary,
		})

	}
	return entity.CampaignDetailResponse{
		ID:               campaign.ID,
		Title:            campaign.Title,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		TargetAmount:     campaign.TargetAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		BackerCount:      campaign.BackerCount,
		Images:           images,
		ImageURL:         imageURL,
		Perks:            perks,
		User:             user,
	}
}
