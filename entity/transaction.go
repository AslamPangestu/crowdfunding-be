package entity

import (
	"time"
)

//Transaction : Mapping Transaction DB
type Transaction struct {
	ID         int
	CampaignID int
	BackerID   int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User     `gorm:"foreignKey:BackerID"`
	Campaign   Campaign `gorm:"foreignKey:CampaignID"`
}

//CampaignTransactionsRequest : Request Get Transactions
type CampaignTransactionsRequest struct {
	ID           int `uri:"id" binding:"required"`
	CampaignerID int
}

//CampaignTransactionsResponse : Response Get Transactions for Campaign
type CampaignTransactionsResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

//UserTransactionsResponse : Response Get Transactions for User
type UserTransactionsResponse struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	Campaign  CampaignTransaction
}

//CampaignTransaction : Detail Campaign Transaction
type CampaignTransaction struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}
