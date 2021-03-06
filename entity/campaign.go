package entity

import (
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
	User             User `gorm:"foreignKey:CampaignerID"`
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

//DETAIL CAMPAIGN ENTIY
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

//CREATE CAMPAIGN ENTITY
//CreateCampaignRequest
type CreateCampaignRequest struct {
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	TargetAmount     int    `json:"target_amount" binding:"required"`
	CampaignerID     int
}
