package services

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/repository"
	"errors"
	"fmt"
	"strings"
	"time"
)

// TransactionInteractor Contract
type TransactionInteractor interface {
	//Get Many
	GetTransactions(page int, pageSize int) (helper.ResponsePagination, error)
	GetTransactionsByCampaignID(request entity.CampaignTransactionsRequest, page int, pageSize int) (helper.ResponsePagination, error)
	GetTransactionsByUserID(userID int, page int, pageSize int) (helper.ResponsePagination, error)
	//Action
	MakeTransaction(form entity.TransactionRequest) (entity.Transaction, error)
}

type transactionService struct {
	repository         repository.TransactionInteractor
	campaignRepository repository.CampaignInteractor
	paymentService     PaymentInteractor
}

// NewTransactionService Initiation
func NewTransactionService(repository repository.TransactionInteractor, campaignRepository repository.CampaignInteractor, paymentService PaymentInteractor) *transactionService {
	return &transactionService{repository, campaignRepository, paymentService}
}

//Get Many
func (s *transactionService) GetTransactions(page int, pageSize int) (helper.ResponsePagination, error) {
	query := entity.Paginate{
		Page:     page,
		PageSize: pageSize,
	}
	models, err := s.repository.FindAll(query)
	if err != nil {
		return models, err
	}
	return models, nil
}

func (s *transactionService) GetTransactionsByCampaignID(request entity.CampaignTransactionsRequest, page int, pageSize int) (helper.ResponsePagination, error) {
	campaign, err := s.campaignRepository.FindOneByID(request.ID)
	if err != nil {
		return helper.ResponsePagination{}, err
	}
	if campaign.CampaignerID != request.CampaignerID {
		return helper.ResponsePagination{}, errors.New("Not an owner of campaign")
	}
	query := entity.Paginate{
		Page:     page,
		PageSize: pageSize,
	}
	models, err := s.repository.FindManyByCampaignID(request.ID, query)
	if err != nil {
		return models, err
	}
	return models, nil
}

func (s *transactionService) GetTransactionsByUserID(userID int, page int, pageSize int) (helper.ResponsePagination, error) {
	query := entity.Paginate{
		Page:     page,
		PageSize: pageSize,
	}
	models, err := s.repository.FindManyByUserID(userID, query)
	if err != nil {
		return models, err
	}
	return models, nil
}

//Action
func (s *transactionService) MakeTransaction(form entity.TransactionRequest) (entity.Transaction, error) {
	model := entity.Transaction{
		CampaignID: form.CampaignID,
		BackerID:   form.Backer.ID,
		Amount:     form.Amount,
		Status:     "pending",
	}

	newTransaction, err := s.repository.Create(model)
	if err != nil {
		return newTransaction, err
	}
	newTransaction.TRXCode = generateTRXCode(model.BackerID, newTransaction.ID, form.CampaignID)
	paymentURL, err := s.paymentService.GeneratePaymentURL(newTransaction, form.Backer)
	if err != nil {
		return newTransaction, err
	}
	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}

func generateTRXCode(userID int, transactionID int, campaignID int) string {
	currentDateTime := time.Now()
	formatedDateTime := strings.ReplaceAll(currentDateTime.Format("2006-01-02"), "-", "")
	return fmt.Sprintf("%d%d%d%s", transactionID, campaignID, userID, formatedDateTime)
}
