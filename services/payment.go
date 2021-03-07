package services

import (
	"crowdfunding/config"
	"crowdfunding/entity"
	"fmt"

	"github.com/veritrans/go-midtrans"
)

// PaymentInteractor Contract
type PaymentInteractor interface {
	GeneratePaymentURL(transaction entity.Transaction, user entity.User) (string, error)
}

type paymentService struct {
}

// NewPaymentService Initiation
func NewPaymentService() *paymentService {
	return &paymentService{}
}

func (s *paymentService) GeneratePaymentURL(transaction entity.Transaction, user entity.User) (string, error) {
	fmt.Println(config.NewPayment())
	snapGateway := midtrans.SnapGateway{
		Client: config.NewPayment(),
	}

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
