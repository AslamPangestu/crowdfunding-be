package entity

import "time"

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
}

//CampaignImage : Mapping CampaignImage DB
type CampaignImage struct {
	ID         int
	CampaignID int
	ImagePath  string
	IsPrimary  bool
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
	ImageURL := ""
	if len(campaign.CampaignImages) > 0 {
		ImageURL = campaign.CampaignImages[0].ImagePath
	}
	res := CampaignResponse{
		ID:               campaign.ID,
		Title:            campaign.Title,
		CampaignerID:     campaign.CampaignerID,
		ShortDescription: campaign.ShortDescription,
		TargetAmount:     campaign.TargetAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageURL:         ImageURL,
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
