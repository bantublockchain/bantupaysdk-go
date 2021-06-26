package merchants

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/bantublockchain/bantupaysdk-go/security"
	"github.com/dghubble/sling"
	"github.com/stellar/go/keypair"
)

//NewMerchant retrieves merchant Info from the bantupay API
func NewMerchant(baseUrl, merchantBantupayUsername, secretKey string) (merchantInfo *Merchant, err error) {

	if baseUrl == "" {
		baseUrl = "https://api-alpha.dev.bantupay.org"
	}
	if merchantBantupayUsername == "" {
		return nil, errors.New("owner username is empty")
	}
	if secretKey == "" {
		return nil, errors.New("secretKey is empty")
	}
	kp := keypair.MustParseFull(secretKey)

	merchantInfo = &Merchant{
		BantupayUsername: strings.ToLower(merchantBantupayUsername),
		BaseURL:          strings.ToLower(baseUrl),
		KP:               kp,
	}
	return merchantInfo, nil
}

//VerifyLoginRequest verifies if loginID is authorized and resturns nil
func (m *Merchant) VerifyLoginRequest(targetUser, loginID string) (userDetail *MerchantBudsInfo, err error) {

	if m.BaseURL == "" {
		m.BaseURL = "https://api-alpha.dev.bantupay.org"
	}
	if m.BantupayUsername == "" {
		return nil, errors.New("owner username is empty")
	}
	if targetUser == "" {
		return nil, errors.New("target username is empty")
	}
	if loginID == "" {
		return nil, errors.New("loginID is empty")
	}
	kp := keypair.MustParseFull(m.KP.Seed())
	// log.Println(kp.Address())

	fullPath := fmt.Sprintf("/v2/merchants/%v/%v/login/%v", m.BantupayUsername, targetUser, loginID)
	signedHttpHeader, err := security.SignHttp(fullPath, "", kp.Seed())
	if err != nil {
		return nil, err
	}
	errorResponse := new(ErrorResponse)
	userDetail = new(MerchantBudsInfo)

	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", kp.Address()).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(m.BaseURL).
		Get(fullPath).Receive(userDetail, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[VerifyLoginRequest]server response error:", *errorResponse)
		return nil, errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[VerifyLoginRequest]request error:", err)
		return nil, err
	}
	// log.Printf("[VerifyLoginRequest][%+v]\n", *userDetail)

	return userDetail, nil
}

//SendLoginRequest sends login request and retrieves QRCode and dynamic link
func (m *Merchant) SendLoginRequest(targetUser, deviceInfo, callbackUrl string) (loginInfo *LoginWithBantupayData, err error) {

	if m.BaseURL == "" {
		m.BaseURL = "https://api-alpha.dev.bantupay.org"
	}
	if m.BantupayUsername == "" {
		return nil, errors.New("owner username is empty")
	}
	if targetUser == "" {
		return nil, errors.New("target username is empty")
	}
	kp := keypair.MustParseFull(m.KP.Seed())
	// log.Println(kp.Address())
	encDeviceInfo := url.QueryEscape(deviceInfo)
	encCallbackUrl := url.QueryEscape(callbackUrl)
	fullPath := fmt.Sprintf("/v2/merchants/%v/%v/login?deviceInfo=%v&callbackUrl=%v", m.BantupayUsername, targetUser, encDeviceInfo, encCallbackUrl)
	signedHttpHeader, err := security.SignHttp(fullPath, "", kp.Seed())
	if err != nil {
		return nil, err
	}
	errorResponse := new(ErrorResponse)
	loginInfo = new(LoginWithBantupayData)
	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", kp.Address()).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(m.BaseURL).
		Post(fullPath).Receive(loginInfo, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[SendLoginRequest]server response error:", *errorResponse)
		return nil, errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[SendLoginRequest]request error:", err)
		return nil, err
	}
	// log.Printf("[SendLoginRequest][%+v]\n", *loginInfo)

	return loginInfo, nil
}

//SendAuthorizationRequest sends 2FA authorization request and retrieves QRCode and dynamic link
func (m *Merchant) SendAuthorizationRequest(targetUser, authDescription, deviceInfo, callbackUrl string) (authInfo *BantupayAuthorizationData, err error) {

	if m.BaseURL == "" {
		m.BaseURL = "https://api-alpha.dev.bantupay.org"
	}
	if m.BantupayUsername == "" {
		return nil, errors.New("owner username is empty")
	}
	if targetUser == "" {
		return nil, errors.New("target username is empty")
	}
	kp := keypair.MustParseFull(m.KP.Seed())
	// log.Println(kp.Address())
	encAuthDescription := url.QueryEscape(authDescription)
	encDeviceInfo := url.QueryEscape(deviceInfo)
	encCallbackUrl := url.QueryEscape(callbackUrl)
	fullPath := fmt.Sprintf("/v2/merchants/%v/%v/authorize?authDescription=%v&deviceInfo=%v&callbackUrl=%v", m.BantupayUsername, targetUser, encAuthDescription, encDeviceInfo, encCallbackUrl)
	signedHttpHeader, err := security.SignHttp(fullPath, "", kp.Seed())
	if err != nil {
		return nil, err
	}
	// type QueryStruct struct {
	// 	AuthDescription string `url:"authDescription"`
	// 	DeviceInfo      string `url:"deviceInfo"`
	// }
	// var q QueryStruct
	// q.AuthDescription = authDescription
	// q.DeviceInfo = deviceInfo
	errorResponse := new(ErrorResponse)
	authInfo = new(BantupayAuthorizationData)
	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", kp.Address()).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(m.BaseURL).
		Post(fullPath).Receive(authInfo, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[SendAuthorizationRequest]server response error:", *errorResponse)
		return nil, errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[SendAuthorizationRequest]request error:", err)
		return nil, err
	}
	// log.Printf("[SendAuthorizationRequest][%+v]\n", *loginInfo)

	return authInfo, nil
}

//VerifyAuthorizationRequest verifies if authID is authorized and resturns nil
func (m *Merchant) VerifyAuthorizationRequest(targetUser, authID string) (err error) {

	if m.BaseURL == "" {
		m.BaseURL = "https://api-alpha.dev.bantupay.org"
	}
	if m.BantupayUsername == "" {
		return errors.New("owner username is empty")
	}
	if targetUser == "" {
		return errors.New("target username is empty")
	}
	if authID == "" {
		return errors.New("authID is empty")
	}
	kp := keypair.MustParseFull(m.KP.Seed())
	// log.Println(kp.Address())

	fullPath := fmt.Sprintf("/v2/merchants/%v/%v/authorize/%v", m.BantupayUsername, targetUser, authID)
	signedHttpHeader, err := security.SignHttp(fullPath, "", kp.Seed())
	if err != nil {
		return err
	}
	type AuthVerifySuccess struct {
		Message string `json:"message"`
	}
	errorResponse := new(ErrorResponse)
	p := new(AuthVerifySuccess)

	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", kp.Address()).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(m.BaseURL).
		Get(fullPath).Receive(p, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[VerifyAuthorizationRequest]server response error:", *errorResponse)
		return errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[VerifyAuthorizationRequest]request error:", err)
		return err
	}
	// log.Printf("[VerifyAuthorizationRequest][%+v]\n", *p)

	return nil
}

//GetPaymentData verifies if authID is authorized and resturns nil
func (m *Merchant) GetPaymentData(targetUser, paymentDestination, assetCode, assetIssuer, amount, memo string) (paymentData *PayWithBantupayData, err error) {

	if m.BaseURL == "" {
		m.BaseURL = "https://api-alpha.dev.bantupay.org"
	}
	if m.BantupayUsername == "" {
		return nil, errors.New("owner username is empty")
	}
	if targetUser == "" {
		return nil, errors.New("target username is empty")
	}
	if paymentDestination == "" {
		return nil, errors.New("paymentDestination is empty")
	}
	if assetCode == "" && assetIssuer != "" {
		return nil, errors.New("assetCode is empty but assetIssuer was specified")
	}
	if len(assetCode) > 0 {
		if len(assetCode) < 3 || len(assetCode) > 12 {
			return nil, errors.New("assetCode is invalid")
		}

	}
	if assetIssuer == "" && assetCode != "" && assetCode != "XBN" {
		return nil, errors.New("assetIssuer is empty but assetCode is specified")
	}
	if amount == "" || amount == "0" {
		return nil, errors.New("amount is empty")
	}
	if len(assetIssuer) > 0 {
		_, e := keypair.ParseAddress(assetIssuer)
		if e != nil {
			return nil, errors.New("invalid assetIssuer")
		}
	}

	kp := keypair.MustParseFull(m.KP.Seed())
	// log.Println(kp.Address())
	encPaymentDestination := url.QueryEscape(paymentDestination)
	encAssetCode := url.QueryEscape(assetCode)
	encAssetIssuer := url.QueryEscape(assetIssuer)
	encAmount := url.QueryEscape(amount)
	encMemo := url.QueryEscape(memo)

	fullPath := fmt.Sprintf("/v2/merchants/%v/%v/payment?amount=%v&assetCode=%v&assetIssuer=%v&memo=%v&paymentDestination=%v", m.BantupayUsername, targetUser, encAmount, encAssetCode, encAssetIssuer, encMemo, encPaymentDestination)
	signedHttpHeader, err := security.SignHttp(fullPath, "", kp.Seed())
	if err != nil {
		return nil, err
	}

	errorResponse := new(ErrorResponse)
	paymentData = new(PayWithBantupayData)

	_, err = sling.New().Set("User-Agent", "BantuPay Go SDK").
		Set("X-BANTUPAY-PUBLIC-KEY", kp.Address()).
		Set("X-BANTUPAY-SIGNER", kp.Address()).
		Set("X-BANTUPAY-SIGNATURE", signedHttpHeader).
		Base(m.BaseURL).
		Get(fullPath).Receive(paymentData, errorResponse)
	//get payload string
	if len(errorResponse.Error) > 0 {
		log.Println("[GetPaymentData]server response error:", *errorResponse)
		return nil, errors.New(errorResponse.Message)
	}
	if err != nil {
		log.Println("[GetPaymentData]request error:", err)
		return nil, err
	}
	// log.Printf("[GetPaymentData][%+v]\n", *p)

	return paymentData, nil
}
