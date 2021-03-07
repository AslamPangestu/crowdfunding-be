package entity

import (
	"time"
)

//Transaction : Mapping Transaction DB
type Transaction struct {
	ID          int
	CampaignID  int
	BackerID    int
	Amount      int
	Status      string
	TRXCode     string
	PaymentCode string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User     `gorm:"foreignKey:BackerID"`
	Campaign    Campaign `gorm:"foreignKey:CampaignID"`
}

//CampaignTransactionsRequest : Request Get Transactions
type CampaignTransactionsRequest struct {
	ID           int `uri:"id" binding:"required"`
	CampaignerID int
}

//CampaignTransactionsResponse : Response Get Transactions for CampaignTransactionsRequest
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
	TRXCode   string    `json:"trx_code"`
	CreatedAt time.Time `json:"created_at"`
	Campaign  CampaignTransaction
}

//CampaignTransaction : Detail Campaign Transaction for UserTransactionsResponse
type CampaignTransaction struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}

//TransactionRequest : Transaction Request
type TransactionRequest struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	BackerID   int
}
