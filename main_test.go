package main

import (
	"log"
	"os"
	"testing"

	"github.com/bantublockchain/bantupaysdk-go/merchants"
)

// payments "github.com/bantublockchain/bantupaysdk-go/payments"

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
// 	merchant, err := merchants.NewMerchant("https://api.bantupay.org", "bantu-airdrop", os.Getenv("MERCHANTSK"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	//send auth for rewards
// 	authData, err := merchant.SendAuthorizationRequest("bantu-airdrop", "claim your Bantu Kano Hangout airdrop", "", "https://api.bantupay.org/v2/callbacks/auth/rewards", 2880)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	log.Printf("QRCode:[%+v]\n", authData.QRCode)
// 	log.Printf("DynamicLink:[%+v]\n", authData.DynamicLink)
// 	log.Printf("AuthID:[%+v]\n", authData.AuthID)

// }

func TestSendAuthorizationRequest(t *testing.T) {
	merchant, err := merchants.NewMerchant("https://api.bantupay.org", "bantu-event", os.Getenv("BANTUEVENTSK"))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	//send auth for rewards
	authData, err := merchant.SendAuthorizationRequest("bantu-event", "join the waitlist for the Bantu XBN staking service", "", "https://api.bantupay.org/v2/callbacks/auth/events", 4880)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	log.Printf("QRCode:[%+v]\n", authData.QRCode)
	log.Printf("DynamicLink:[%+v]\n", authData.DynamicLink)
	log.Printf("AuthID:[%+v]\n", authData.AuthID)

}

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

// func TestGetPaymentData(t *testing.T) {
// 	merchant, err := merchants.NewMerchant("https://api.bantupay.org", "timbuktu", os.Getenv("MERCHANTSK"))
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	assetCode := ""
// 	assetIssuer := ""
// 	amount := "400"
// 	memo := "Bantu Hangout Abuja T-Shirt"
// 	paymentDestination := "t-shirt-merch"
// 	paymentData, err := merchant.GetPaymentData("ric", paymentDestination, assetCode, assetIssuer, amount, memo)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	log.Printf("[%v] Payment QRCode: [%+v]\n\n", memo, paymentData.QRCode)
// 	log.Printf("[%v] Payment DynamicLink:  [%+v]\n", memo, paymentData.DynamicLink)

// }
