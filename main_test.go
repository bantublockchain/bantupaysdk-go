package main

import (
	"log"
	"os"
	"testing"

	"github.com/bantublockchain/bantupaysdk-go/merchants"
	// payments "github.com/bantublockchain/bantupaysdk-go/payments"
)

// func TestExpressPay(t *testing.T) {
// 	paymentDetail := payments.PaymentInstance()
// 	paymentDetail.Destination = "ric7"
// 	paymentDetail.Amount = "2"
// 	paymentDetail.Memo = "BantuPaySDKGo ExpressPay"
// 	baseURL := "" //default when empty

// 	err := paymentDetail.ExpressPay(baseURL, os.Getenv("BANTUPAY_USER"), os.Getenv("BANTUPAY_USER_SECRET"), "", os.Getenv("CHANNEL_ACCOUNT_SECRET"))
// 	// err := paymentDetail.ExpressPay(baseURL, "", os.Getenv("BANTUPAY_USER_SECRET"), "", os.Getenv("CHANNEL_ACCOUNT_SECRET"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// }

// func TestSendLoginRequest(t *testing.T) {
// 	merchant, err := merchants.NewMerchant("", "xbnp2p", os.Getenv("MERCHANTSK"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	loginInfo, err := merchant.SendLoginRequest("ric")
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	log.Printf("qrCode:[%v]\n", loginInfo.QRCode)
// 	log.Printf("Dynamic Link:[%v]\n", loginInfo.DynamicLink)
// 	log.Printf("Login ID:[%v]\n", loginInfo.LoginID)

// }

// func TestVerifyLoginRequest(t *testing.T) {
// 	merchant, err := merchants.NewMerchant("", "xbnp2p", os.Getenv("MERCHANTSK"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	userInfo, err := merchant.VerifyLoginRequest("ric", "849f2249-3df3-4b64-ae8d-14f204493f17")
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	log.Printf("userInfo:[%+v]\n", *userInfo)

// }

// func TestSendAuthorizationRequest(t *testing.T) {
// 	merchant, err := merchants.NewMerchant("", "xbnp2p", os.Getenv("MERCHANTSK"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	authData, err := merchant.SendAuthorizationRequest("ric", "Add New Payment Method: Account-1004847601, Bank-Zenith PLC, Country - NG", "")
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	log.Printf("QRCode:[%+v]\n", authData.QRCode)
// 	log.Printf("DynamicLink:[%+v]\n", authData.DynamicLink)
// 	log.Printf("AuthID:[%+v]\n", authData.AuthID)

// }

// func TestVerifyAuthorizationRequest(t *testing.T) {
// 	merchant, err := merchants.NewMerchant("", "xbnp2p", os.Getenv("MERCHANTSK"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	authID := "8ba94e6d-a046-4d38-8dd9-a46c3524f89e"
// 	err = merchant.VerifyAuthorizationRequest("ric", authID)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	log.Printf("Authorization [%v] successfully verified.\n", authID)

// }
func TestGetPaymentData(t *testing.T) {
	merchant, err := merchants.NewMerchant("", "xbnp2p", os.Getenv("MERCHANTSK"))
	if err != nil {
		t.Errorf(err.Error())
	}
	assetCode := ""
	assetIssuer := ""
	amount := "0.02"
	memo := "34hggf76398hfggsf736bvg"
	paymentDestination := "xbnp2p"
	paymentData, err := merchant.GetPaymentData("ric", paymentDestination, assetCode, assetIssuer, amount, memo)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	log.Printf("Payment QRCode:[%+v]\n\n", paymentData.QRCode)
	log.Printf("Payment DynamicLink:[%+v]\n", paymentData.DynamicLink)

}
