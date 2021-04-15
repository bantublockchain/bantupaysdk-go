package main

import (
	"os"
	"testing"

	"github.com/bantublockchain/bantupaysdk-go/funcs"
)

func TestExpressPay(t *testing.T) {
	var paymentDetail funcs.PaymentInfo
	paymentDetail.Destination = "ric"
	paymentDetail.Amount = "2"
	paymentDetail.Memo = "BantuPaySDKGo ExpressPay"
	baseURL := "https://api-alpha.dev.bantupay.org"

	err := paymentDetail.ExpressPay(baseURL, os.Getenv("BANTUPAY_USER"), os.Getenv("BANTUPAY_USER_SECRET"), "", os.Getenv("CHANNEL_ACCOUNT_SECRET"))
	if err != nil {
		t.Errorf(err.Error())
	}

}
