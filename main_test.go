package main

import (
	"os"
	"testing"

	"github.com/bantublockchain/bantupaysdk-go/funcs"
)

func TestExpressPay(t *testing.T) {
	paymentDetail := funcs.PaymentInstance()
	paymentDetail.Destination = "ric7"
	paymentDetail.Amount = "2"
	paymentDetail.Memo = "BantuPaySDKGo ExpressPay"
	baseURL := "" //default when empty

	err := paymentDetail.ExpressPay(baseURL, os.Getenv("BANTUPAY_USER"), os.Getenv("BANTUPAY_USER_SECRET"), "", os.Getenv("CHANNEL_ACCOUNT_SECRET"))
	// err := paymentDetail.ExpressPay(baseURL, "", os.Getenv("BANTUPAY_USER_SECRET"), "", os.Getenv("CHANNEL_ACCOUNT_SECRET"))
	if err != nil {
		t.Errorf(err.Error())
	}

}
