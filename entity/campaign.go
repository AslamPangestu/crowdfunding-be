package entity

import (
	"time"

	"github.com/leekchan/accounting"
)

//CAMPAIGN

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

//CampaignIDRequest : Request Detail Campaign by ID uri
type CampaignIDRequest struct {
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
	BackerCount      int                   `json:"backer_count"`
	Slug             string                `json:"slug"`
	ImageURL         string                `json:"image_url"`
	Perks            []string              `json:"perks"`
	User             UserCampaignDetail    `json:"user"`
	Images           []ImageCampaignDetail `json:"images"`
}

//UserCampaignDetail : User Detail Campaign
type UserCampaignDetail struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

//FormCampaignRequest : Request to create new campaign
type FormCampaignRequest struct {
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	TargetAmount     int    `json:"target_amount" binding:"required"`
	CampaignerID     int
}

//CAMPAIGN IMAGE

//CampaignImage : Mapping CampaignImage DB
type CampaignImage struct {
	ID         int
	CampaignID int
	ImagePath  string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

//ImageCampaignDetail : Response Detail Campaign for Image
type ImageCampaignDetail struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

//UploadCampaignImageRequest : Request to upload images campaign
type UploadCampaignImageRequest struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	UserID     int
}

//TargetAmountFormatIDR : Adapter Campaign Detail
func (c Campaign) TargetAmountFormatIDR() string {
	format := accounting.Accounting{Symbol: "IDR ", Precision: 2, Thousand: ".", Decimal: ","}
	return format.FormatMoney(c.TargetAmount)
}

//TargetAmountFormatIDR : Adapter Campaign Detail
func (c Campaign) CurrentAmountFormatIDR() string {
	format := accounting.Accounting{Symbol: "IDR ", Precision: 2, Thousand: ".", Decimal: ","}
	return format.FormatMoney(c.CurrentAmount)
}

//CreateUserForm : Mapping Form Create Campaign
type CreateCampaignForm struct {
	Title            string `form:"title" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	TargetAmount     int    `form:"target_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	UserID           int    `form:"user_id" binding:"required"`
	Users            []User
	Error            error
}

//EditCampaignForm : Mapping Form Edit Campaign
type EditCampaignForm struct {
	ID               int
	Title            string `form:"title" binding:"required"`
	ShortDescription string `form:"short_description" binding:"required"`
	Description      string `form:"description" binding:"required"`
	TargetAmount     int    `form:"target_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	Error            error
}
