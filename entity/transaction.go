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
	User       User `gorm:"foreignKey:BackerID"`
}

//List Transaciton Campaign
//GetCampaignTransactionsRequest : Request Get Transactions
type GetCampaignTransactionsRequest struct {
	ID   int `uri:"id" binding:"required"`
	User User
}

//GetCampaignTransactionsResponse : Response Get Transactions
type GetCampaignTransactionsResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
