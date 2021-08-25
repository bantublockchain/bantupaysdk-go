package merchants

import "github.com/stellar/go/keypair"

//Merchant holds Merchant data model
type Merchant struct {
	BantupayUsername string `json:"bantupayUsername" `
	BaseURL          string `json:"baseUrl"`
	KP               *keypair.Full
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type PayWithBantupayData struct {
	DynamicLink string `json:"dynamicLink"`
	QRCode      string `json:"qrCode"`
}
type LoginWithBantupayData struct {
	DynamicLink string `json:"dynamicLink"`
	QRCode      string `json:"qrCode"`
	LoginID     string `json:"loginId"`
}

type BantupayAuthorizationData struct {
	DynamicLink string `json:"dynamicLink"`
	QRCode      string `json:"qrCode"`
	AuthID      string `json:"authId"`
}

//MerchantBudsInfo model for bantu user directory info
type MerchantBudsInfo struct {
	Username       string `json:"username"`
	PublicKey      string `json:"publicKey"`
	Email          string `json:"email"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	MiddleName     string `json:"middleName"`
	Mobile         string `json:"mobile"`
	BantuTalk      string `json:"bantuTalk"`
	ImageThumbnail string `json:"imageThumbnail"`
	Verified       uint   `json:"verified"`
	Suspended      uint   `json:"suspended"`
}
type MerchantRequestInput struct {
	AuthDescription   string `json:"authDescription,omitempty"`
	DeviceInfo        string `json:"deviceInfo,omitempty"`
	CallbackURL       string `json:"callbackUrl,omitempty"`
	ValidityInMinutes int    `json:"validityInMinutes,omitempty"`
}
type MerchantPushNotification struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}
