package payments

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/bantublockchain/bantupaysdk-go/security"
	"github.com/dghubble/sling"
	"github.com/shopspring/decimal"
	"github.com/stellar/go/keypair"
)

// var acceptedChars = "abcdefghijklmnopqrstuvwxyz_1234567890/"

//SwapInstance returns a new payment struct
func SwapInstance() (p *SwapSendInfo) {

	return new(SwapSendInfo)
}

//ConfirmSwapDetail calls bantupayAPI service to generate transactions and confirm swap details. This is the first method to be invoked during swap. the result is returned for UI display of swap details
func (p *SwapSendInfo) ConfirmSwapDetail(baseUrl, ownerUsername, secretKey, ownerPublicKey string) (err error) {

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

	if len(p.Transaction) > 0 {
		return errors.New("transaction must be empty")
	}
	if p.DestinationAssetCode == "XBN" {
		p.DestinationAssetIssuer = ""
	}
	if p.SourceAssetCode == "XBN" {
		p.SourceAssetIssuer = ""
	}
	if decimal.RequireFromString(p.SourceAmount).LessThan(decimal.NewFromFloat(0.0000001)) {
		return errors.New("amount must be at least 0.0000001")
	}

	if len(p.DestinationAssetCode) > 0 {
		_, e := keypair.Parse(p.DestinationAssetIssuer)
		if e != nil {
			return errors.New("destinationAssetIssuer " + p.DestinationAssetIssuer + " is invalid")
		}

		if len(p.DestinationAssetCode) > 12 {
			return errors.New("destinationAssetCode " + p.DestinationAssetCode + " is invalid")
		}
	}
	if len(p.SourceAssetIssuer) > 0 {
		_, e := keypair.Parse(p.SourceAssetIssuer)
		if e != nil {
			return errors.New("sourceAssetIssuer " + p.SourceAssetIssuer + " is invalid")
		}

		if len(p.SourceAssetCode) > 12 {
			return errors.New("sourceAssetCode " + p.SourceAssetCode + " is invalid")
		}
	}
	errorResponse := new(ErrorResponse)
	var sEnc string
	if strings.Contains(ownerUsername, "/") {
		sEnc = base64.URLEncoding.EncodeToString([]byte(ownerUsername))

	} else {
		sEnc = ownerUsername
	}
	// log.Printf("encoded username[%v]: [%v]\n", ownerUsername, sEnc)
	// /v2/users/:targetUser/swaps
	fullPath := "/v2/users/" + sEnc + "/swaps"
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
	signedHttpHeader, err := security.SignHttp(fullPath, string(body), secretKey)
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
		// log.Println("[ConfirmSwapDetail]server response error:", *errorResponse)
		return errors.New(errorResponse.Message)
	}
	if err != nil {
		// log.Println("[ConfirmSwapDetail]request error:", err)
		return err
	}
	// log.Printf("[ConfirmSwapDetail][%+v]\n", *p)
	return nil
}

//MakeSwap makes payment after transaction has been generated  using ConfirmSwap method
func (p *SwapSendInfo) MakeSwap(baseUrl, ownerUsername, secretKey, ownerPublicKey string) (err error) {

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

	if len(p.Transaction) > 0 {
		return errors.New("transaction must be empty")
	}
	if p.DestinationAssetCode == "XBN" {
		p.DestinationAssetIssuer = ""
	}
	if p.SourceAssetCode == "XBN" {
		p.SourceAssetIssuer = ""
	}
	if decimal.RequireFromString(p.SourceAmount).LessThan(decimal.NewFromFloat(0.0000001)) {
		return errors.New("amount must be at least 0.0000001")
	}

	if len(p.DestinationAssetCode) > 0 {
		_, e := keypair.Parse(p.DestinationAssetIssuer)
		if e != nil {
			return errors.New("destinationAssetIssuer " + p.DestinationAssetIssuer + " is invalid")
		}

		if len(p.DestinationAssetCode) > 12 {
			return errors.New("destinationAssetCode " + p.DestinationAssetCode + " is invalid")
		}
	}
	if len(p.SourceAssetIssuer) > 0 {
		_, e := keypair.Parse(p.SourceAssetIssuer)
		if e != nil {
			return errors.New("sourceAssetIssuer " + p.SourceAssetIssuer + " is invalid")
		}

		if len(p.SourceAssetCode) > 12 {
			return errors.New("sourceAssetCode " + p.SourceAssetCode + " is invalid")
		}
	}
	errorResponse := new(ErrorResponse)
	var sEnc string
	if strings.Contains(ownerUsername, "/") {
		sEnc = base64.URLEncoding.EncodeToString([]byte(ownerUsername))

	} else {
		sEnc = ownerUsername
	}

	signedBase64, err := security.SignBase64Txn(kp.Seed(), p.Transaction, p.NetworkPassPhrase)
	if err != nil {
		// log.Println("[MakeSwap] Transaction signing failed:", err)
		return err
	}

	p.TransactionSignature = signedBase64

	fullPath := "/v2/users/" + sEnc + "/swaps"
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
	signedHttpHeader, err := security.SignHttp(fullPath, string(body), secretKey)
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
		// log.Println("[MakeSwap]server response error:", *errorResponse)
		return errors.New(errorResponse.Message)
	}
	if err != nil {
		// log.Println("[MakeSwap]request error:", err)
		return err
	}
	// log.Printf("[MakeSwap][%+v]\n", *p)
	return nil
}

//ExpressSwap used by bots and bulk payment systems. this makes payments bypassing any form of confirmation. It chains together ConfirmPaymentDetails and MakePayment methods
func (p *SwapSendInfo) ExpressSwap(baseUrl, ownerUsername, secretKey, ownerPublicKey string) (err error) {

	if baseUrl == "" {
		baseUrl = "https://api-alpha.dev.bantupay.org"
	}
	err = p.ConfirmSwapDetail(baseUrl, ownerUsername, secretKey, ownerPublicKey)
	if err != nil {
		return err
	}
	return p.MakeSwap(baseUrl, ownerUsername, secretKey, ownerPublicKey)

}
