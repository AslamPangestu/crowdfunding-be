package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
	"errors"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	GetTransactionsByCampaignID(request entity.GetCampaignTransactionsRequest) ([]entity.Transaction, error)
}

type transactionService struct {
	repository         repository.TransactionInteractor
	campaignRepository repository.CampaignInteractor
}

// NewTransactionService Initiation
func NewTransactionService(repository repository.TransactionInteractor, campaignRepository repository.CampaignInteractor) *transactionService {
	return &transactionService{repository, campaignRepository}
}

func (s *transactionService) GetTransactionsByCampaignID(request entity.GetCampaignTransactionsRequest) ([]entity.Transaction, error) {
	campaign, err := s.campaignRepository.FindOneByID(request.ID)
	if err != nil {
		return []entity.Transaction{}, err
	}
	if campaign.CampaignerID == request.User.ID {
		return []entity.Transaction{}, errors.New("Not an owner of campaign")
	}
	transactions, err := s.repository.FindManyByCampaignID(request.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
