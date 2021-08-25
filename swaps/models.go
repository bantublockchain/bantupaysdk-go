package payments

//PaymentInfo holds payment information
type PaymentInfo struct {
	Destination             string   `json:"destination"`                       //Receiver of the asset
	Memo                    string   `json:"memo,omitempty"`                    //Any description for the transaction. max of 28bytes
	AssetIssuer             string   `json:"assetIssuer"`                       //issuer of the asset you are sending
	AssetCode               string   `json:"assetCode"`                         //asset code of the asset you are sending
	Amount                  string   `json:"amount"`                            //the amount in 7dp. cannot be "0"
	Transaction             string   `json:"transaction,omitempty"`             //this is returned by the API service after ConfirmPaymentDetails has been invoked
	TransactionSignature    string   `json:"transactionSignature,omitempty"`    //used to send signature of the transaction back to the server on the second call
	TransactionID           string   `json:"transactionId,omitempty"`           //used to return transactionID for a successful transaction
	NetworkPassPhrase       string   `json:"networkPassPhrase,omitempty"`       //used to return network passphrase of the network connected to
	DestinationFirstName    string   `json:"destinationFirstName,omitempty"`    //used to return destination first name
	DestinationLastName     string   `json:"destinationLastName,omitempty"`     //used to return destination last name
	DestinationThumbnail    string   `json:"destinationThumbnail,omitempty"`    //used to return destination thumbnail
	ChannelAccount          string   `json:"channelAccount,omitempty"`          //used to send channel account that will be used for this transaction if any
	ChannelAccountSignature string   `json:"channelAccountSignature,omitempty"` //used to preserve the signature from the channel accounts
	Messages                []string `json:"messages,omitempty"`                //used to return any warnings or extra messages important to the user
}

//SwapSendPathInput model struct for user input when importing keys for new user
type SwapSendInfo struct {
	DestinationAssetCode   string   `json:"destinationAssetCode"`
	DestinationAssetIssuer string   `json:"destinationAssetIssuer"`
	SwappedEstimate        string   `json:"swappedEstimate"`
	SourceAssetCode        string   `json:"sourceAssetCode"`
	SourceAssetIssuer      string   `json:"sourceAssetIssuer"`
	SourceAmount           string   `json:"sourceAmount" `
	Transaction            string   `json:"transaction"`
	TransactionSignature   string   `json:"transactionSignature"`
	TransactionID          string   `json:"transactionId"`
	NetworkPassPhrase      string   `json:"networkPassPhrase"`
	Messages               []string `json:"messages"`
	Memo                   string   `json:"-"`
}
type ErrorResponse struct {
	Error   string `json:"error"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
