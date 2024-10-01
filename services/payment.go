package services

import (
	"crowdfunding/config"
	"crowdfunding/entity"
	"crowdfunding/repository"
	"errors"

	"github.com/veritrans/go-midtrans"
)

// PaymentInteractor Contract
type PaymentInteractor interface {
	GeneratePaymentURL(transaction entity.Transaction, user entity.User) (string, error)
	ProcessPayment(form entity.TransactionNotificationRequest) error
}

type paymentService struct {
	transactionRepository repository.TransactionInteractor
	campaignRepository    repository.CampaignInteractor
}

// NewPaymentService Initiation
func NewPaymentService(transactionRepository repository.TransactionInteractor, campaignRepository repository.CampaignInteractor) *paymentService {
	return &paymentService{transactionRepository, campaignRepository}
}

func (s *paymentService) GeneratePaymentURL(transaction entity.Transaction, user entity.User) (string, error) {
	snapGateway := config.NewPayment()

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.TRXCode,
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}

func (s *paymentService) ProcessPayment(form entity.TransactionNotificationRequest) error {
	transaction, err := s.transactionRepository.FindOneByTrxCode(form.OrderID)
	if err != nil {
		return err
	}
	if transaction.ID == "" {
		return errors.New("TRANSACTION NOT FOUND")
	}
	//IF Credit Card
	if form.PaymentType == "credit_card" && form.TransactionStatus == "capture" && form.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if form.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if form.TransactionStatus == "deny" || form.TransactionStatus == "expire" || form.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindOneByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}
	if campaign.ID == "" {
		return errors.New("CAMPAIGN NOT FOUND")
	}
	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount
		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
