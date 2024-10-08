package entity

import (
	"time"

	"github.com/leekchan/accounting"
)

// Transaction : Mapping Transaction DB
type Transaction struct {
	ID         string `gorm:"column:xata_id"`
	CampaignID string
	BackerID   string
	Amount     int
	Status     string
	TRXCode    string
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User     `gorm:"foreignKey:BackerID"`
	Campaign   Campaign `gorm:"foreignKey:CampaignID"`
}

// CampaignTransactionsRequest : Request Get Transactions
type CampaignTransactionsRequest struct {
	ID           string `uri:"id" binding:"required"`
	CampaignerID string
}

// CampaignTransactionsResponse : Response Get Transactions for CampaignTransactionsRequest
type CampaignTransactionsResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// UserTransactionsResponse : Response Get Transactions for User
type UserTransactionsResponse struct {
	ID        string              `json:"id"`
	Amount    int                 `json:"amount"`
	Status    string              `json:"status"`
	TRXCode   string              `json:"trx_code"`
	CreatedAt time.Time           `json:"created_at"`
	Campaign  CampaignTransaction `json:"campaign"`
}

// CampaignTransaction : Detail Campaign Transaction for UserTransactionsResponse
type CampaignTransaction struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}

// TransactionRequest : Transaction Request
type TransactionRequest struct {
	Amount     int    `json:"amount" binding:"required"`
	CampaignID string `json:"campaign_id" binding:"required"`
	Backer     User
}

// TransactionResponse : Transaction Response
type TransactionResponse struct {
	ID         string `json:"id"`
	CampaignID string `json:"campaign_id"`
	BackerID   string `json:"backer_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	TRXCode    string `json:"trx_code"`
	PaymentURL string `json:"payment_url"`
}

// TransactionNotificationRequest : Transaction Notif Req
type TransactionNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

// TargetAmountFormatIDR : Adapter Campaign Detail
func (t Transaction) AmountFormatIDR() string {
	format := accounting.Accounting{Symbol: "IDR ", Precision: 2, Thousand: ".", Decimal: ","}
	return format.FormatMoney(t.Amount)
}
