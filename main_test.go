package main

import (
	"os"
	"testing"
)

func TestExpressPay(t *testing.T) {
	paymentDetail := NewPayment()
	paymentDetail.Destination = "ric"
	paymentDetail.Amount = "2"
	paymentDetail.Memo = "BantuPaySDKGo ExpressPay"
	baseURL := "" //default when empty

	err := paymentDetail.ExpressPay(baseURL, os.Getenv("BANTUPAY_USER"), os.Getenv("BANTUPAY_USER_SECRET"), "", os.Getenv("CHANNEL_ACCOUNT_SECRET"))
	if err != nil {
		t.Errorf(err.Error())
	}

}
