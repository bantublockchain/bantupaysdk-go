package funcs

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/dghubble/sling"
	"github.com/shopspring/decimal"
	"github.com/stellar/go/keypair"
)

//PaymentInstance returns a new payment struct
func PaymentInstance() (p *PaymentInfo) {

	return new(PaymentInfo)
}

//ConfirmPaymentDetail calls bantupayAPI service to generate transactions and confirm payment details. This is the first method to be invoked during payment. the result is returned for UI display of payment details
func (p *PaymentInfo) ConfirmPaymentDetail(baseUrl, ownerUsername, secretKey, ownerPublicKey, channelAccountSecret string) (err error) {
	if strings.Contains(ownerUsername, "%") {
		u, e := url.QueryUnescape(ownerUsername)
		if e == nil {
			ownerUsername = u
		}
	}
	if p == nil {
		return errors.New("paymentInfo struct is nil")
	}
	if baseUrl == "" {
		baseUrl = "https://api-alpha.dev.bantupay.org"
	}

	if ownerUsername == "" {
		return errors.New("owner username is empty")
	}
	if secretKey == "" {
		return errors.New("secretKey is empty")
	}
	kp := keypair.MustParseFull(secretKey)
	// log.Println(kp.Address())
	if len(channelAccountSecret) > 0 {
		ckp := keypair.MustParseFull(channelAccountSecret)
		p.ChannelAccount = ckp.Address()
		// log.Println(ckp.Address())

	}
	if len(p.Transaction) > 0 {
		return errors.New("transaction must be empty")
	}
	if len(p.Destination) == 0 {
		return errors.New("destination cannot be empty")
	}
	if decimal.RequireFromString(p.Amount).LessThan(decimal.NewFromFloat(0.0000001)) {
		return errors.New("amount must be at least 0.0000001")
	}

	if len(p.AssetCode) > 0 {
		_, e := keypair.Parse(p.AssetIssuer)
		if e != nil {
			return errors.New("assetIssuer " + p.AssetIssuer + " is invalid")
		}

		if len(p.AssetCode) > 12 {
			return errors.New("assetCode " + p.AssetCode + " is invalid")
		}
	}
	errorResponse := new(ErrorResponse)
	escapedOwnerUsername := url.QueryEscape(ownerUsername)
	fullPath := "/v2/users/" + escapedOwnerUsername + "/payments"
	// log.Println("fullPath:", fullPath)
	body, err := json.Marshal(*p)

	if err != nil {
		return err
	}
	if ownerPublicKey == "" {
		ownerPublicKey = kp.Address()
	}
	// log.Println("body:", body)
	// log.Println("string(body):", string(body))
	signedHttpHeader, err := SignHttp(fullPath, string(body), secretKey)
	if err != nil {
		return err
	}
	// log.Println("X-BANTUPAY-SIGNATURE:", signedHttpHeader)
	// log.Println("X-BANTUPAY-PUBLIC_KEY:", ownerPublicKey)
	// log.Println("X-BANTUPAY-SIGNER:", kp.Address())
	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", ownerPublicKey).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(baseUrl).
		Post(fullPath).BodyJSON(*p).Receive(p, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[ConfirmPaymentDetail]server response error:", *errorResponse)
		return errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[ConfirmPaymentDetail]request error:", err)
		return err
	}
	// log.Printf("[ConfirmPaymentDetail][%+v]\n", *p)
	return nil
}

//MakePayment makes payment after transaction has been generated  using ConfirmPaymentDetails method
func (p *PaymentInfo) MakePayment(baseUrl, ownerUsername, secretKey, ownerPublicKey, channelAccountSecret string) (err error) {
	if strings.Contains(ownerUsername, "%") {
		u, e := url.QueryUnescape(ownerUsername)
		if e == nil {
			ownerUsername = u
		}
	}

	if p == nil {
		return errors.New("paymentInfo struct is nil")
	}
	if baseUrl == "" {
		baseUrl = "https://api-alpha.dev.bantupay.org"
	}

	if ownerUsername == "" {
		return errors.New("owner username is empty")
	}
	if secretKey == "" {
		return errors.New("secretKey is empty")
	}
	kp := keypair.MustParseFull(secretKey)
	var ckp *keypair.Full
	if len(channelAccountSecret) > 0 {
		ckp = keypair.MustParseFull(channelAccountSecret)
		p.ChannelAccount = ckp.Address()
	}
	if p.Transaction == "" {
		return errors.New("please first invoke method ConfirmPaymentDetail before invoking this method")
	}
	if p.NetworkPassPhrase == "" {
		return errors.New("please first invoke method ConfirmPaymentDetail successfully before invoking this method")
	}
	if p.TransactionID != "" {
		return errors.New("transactionID must be empty")
	}
	if len(p.Destination) == 0 {
		return errors.New("destination cannot be empty")
	}
	if decimal.RequireFromString(p.Amount).LessThan(decimal.NewFromFloat(0.0000001)) {
		return errors.New("amount must be at least 0.0000001")
	}

	if len(p.AssetCode) > 0 {
		_, e := keypair.Parse(p.AssetIssuer)
		if e != nil {
			return errors.New("assetIssuer " + p.AssetIssuer + " is invalid")
		}

		if len(p.AssetCode) > 12 {
			return errors.New("assetCode " + p.AssetCode + " is invalid")
		}
	}

	//sign transaction
	if len(p.ChannelAccount) == 56 {
		// dsigned := paymentInfo.Transaction
		//channel account used. sign transaction with channel account

		dsigned, err := SignBase64Txn(ckp.Seed(), p.Transaction, p.NetworkPassPhrase)
		if err != nil {
			log.Println("[MakePayment] Channel Account Transaction signing failed:", err)
			return err
		}
		p.ChannelAccountSignature = dsigned

	}

	signedBase64, err := SignBase64Txn(kp.Seed(), p.Transaction, p.NetworkPassPhrase)
	if err != nil {
		log.Println("[MakePayment] Transaction signing failed:", err)
		return err
	}

	p.TransactionSignature = signedBase64

	errorResponse := new(ErrorResponse)
	escapedOwnerUsername := url.QueryEscape(ownerUsername)
	fullPath := "/v2/users/" + escapedOwnerUsername + "/payments"
	// log.Println("fullPath:", fullPath)
	body, err := json.Marshal(*p)

	if err != nil {
		return err
	}
	if ownerPublicKey == "" {
		ownerPublicKey = kp.Address()
	}
	// log.Println("body:", body)
	// log.Println("string(body):", string(body))
	signedHttpHeader, err := SignHttp(fullPath, string(body), secretKey)
	if err != nil {
		return err
	}
	// log.Println("X-BANTUPAY-SIGNATURE:", signedHttpHeader)
	// log.Println("X-BANTUPAY-PUBLIC_KEY:", ownerPublicKey)
	// log.Println("X-BANTUPAY-SIGNER:", kp.Address())
	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", ownerPublicKey).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(baseUrl).
		Post(fullPath).BodyJSON(*p).Receive(p, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[MakePayment]server response error:", *errorResponse)
		return errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[MakePayment]request error:", err)
		return err
	}
	// log.Printf("[MakePayment][%+v]\n", *p)
	return nil
}

//ExpressPay used by bots and bulk payment systems. this makes payments bypassing any form of confirmation. It chains together ConfirmPaymentDetails and MakePayment methods
func (p *PaymentInfo) ExpressPay(baseUrl, ownerUsername, secretKey, ownerPublicKey, channelAccountSecret string) (err error) {
	if strings.Contains(ownerUsername, "%") {
		u, e := url.QueryUnescape(ownerUsername)
		if e == nil {
			ownerUsername = u
		}
	}
	if baseUrl == "" {
		baseUrl = "https://api-alpha.dev.bantupay.org"
	}
	err = p.ConfirmPaymentDetail(baseUrl, ownerUsername, secretKey, ownerPublicKey, channelAccountSecret)
	if err != nil {
		return err
	}
	return p.MakePayment(baseUrl, ownerUsername, secretKey, ownerPublicKey, channelAccountSecret)

}
