package funcs

//PaymentInfo holds payment information
type PaymentInfo struct {
	Destination             string   `json:"destination"`
	Memo                    string   `json:"memo,omitempty"`
	AssetIssuer             string   `json:"assetIssuer"`
	AssetCode               string   `json:"assetCode"`
	Amount                  string   `json:"amount"`
	Transaction             string   `json:"transaction,omitempty"`
	TransactionSignature    string   `json:"transactionSignature,omitempty"`
	TransactionID           string   `json:"transactionId,omitempty"`
	NetworkPassPhrase       string   `json:"networkPassPhrase,omitempty"`
	DestinationFirstName    string   `json:"destinationFirstName,omitempty"`
	DestinationLastName     string   `json:"destinationLastName,omitempty"`
	DestinationThumbnail    string   `json:"destinationThumbnail,omitempty"`
	ChannelAccount          string   `json:"channelAccount,omitempty"`
	ChannelAccountSignature string   `json:"channelAccountSignature,omitempty"`
	Messages                []string `json:"messages,omitempty"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
