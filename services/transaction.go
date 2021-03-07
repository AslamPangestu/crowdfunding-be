package services

import (
	"crowdfunding/entity"
	"crowdfunding/repository"
	"errors"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	GetTransactionsByCampaignID(request entity.CampaignTransactionsRequest) ([]entity.Transaction, error)
	GetTransactionsByUserID(userID int) ([]entity.Transaction, error)
}

type transactionService struct {
	repository         repository.TransactionInteractor
	campaignRepository repository.CampaignInteractor
}

// NewTransactionService Initiation
func NewTransactionService(repository repository.TransactionInteractor, campaignRepository repository.CampaignInteractor) *transactionService {
	return &transactionService{repository, campaignRepository}
}

func (s *transactionService) GetTransactionsByCampaignID(request entity.CampaignTransactionsRequest) ([]entity.Transaction, error) {
	campaign, err := s.campaignRepository.FindOneByID(request.ID)
	if err != nil {
		return []entity.Transaction{}, err
	}
	if campaign.CampaignerID == request.CampaignerID {
		return []entity.Transaction{}, errors.New("Not an owner of campaign")
	}
	models, err := s.repository.FindManyByCampaignID(request.ID)
	if err != nil {
		return models, err
	}
	return models, nil
}

func (s *transactionService) GetTransactionsByUserID(userID int) ([]entity.Transaction, error) {
	models, err := s.repository.FindManyByUserID(userID)
	if err != nil {
		return models, err
	}
	return models, nil
}
