package config

import (
	"os"

	"github.com/veritrans/go-midtrans"
)

// NewPayment : Initialize Midtrans
func NewPayment() midtrans.SnapGateway {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("PAYMENT_SERVER_KEY")
	midclient.ClientKey = os.Getenv("PAYMENT_CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox
	return midtrans.SnapGateway{
		Client: midclient,
	}
}
