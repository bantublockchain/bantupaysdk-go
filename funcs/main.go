package funcs

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"

	bantupaysdkErrors "github.com/bantublockchain/bantupaysdk-go/errors"

	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
)

//SignString returns a signed base64 encoded string of toSign
func SignString(toSign string, secretKey string) (string, error) {
	kp, keyPairError := keypair.ParseFull(secretKey)
	if keyPairError != nil {
		return "", keyPairError
	}
	toSign = strings.TrimSpace(toSign)
	// log.Printf("toSign:[%v]\n", toSign)
	signature, err := kp.SignBase64([]byte(toSign))
	// ps, _ := base64.StdEncoding.DecodeString(signature)
	// log.Printf("provided signature:[%v]\n", ps)

	if err != nil {
		return "", err
	}

	return signature, nil
}

//SignHttp returns a signed base64 encoded string of fullPathWithQuery+body
func SignHttp(fullPathWithQuery string, body string, secretKey string) (string, error) {
	// log.Printf("path + string:[%v]\n", fullPathWithQuery+body)
	body = strings.TrimSpace(body)
	fullPathWithQuery = strings.TrimSpace(fullPathWithQuery)
	signature, err := SignString(fullPathWithQuery+body, secretKey)

	if err != nil {
		return "", err
	}

	return signature, nil
}

//SignBase64Txn signs the transaction hash from base64Txn string using the secret key
func SignBase64Txn(secretKey string, base64Txn string, networkPassPhrase string) (string, error) {

	kp, keyPairError := keypair.ParseFull(secretKey)
	if keyPairError != nil {
		return "", keyPairError
	}

	tx, err := txnbuild.TransactionFromXDR(base64Txn)
	if err != nil {
		return "", err
	}

	txn, b := tx.Transaction()

	if !b {
		return "", errors.New("not a txn")
	}

	bytes, err := txn.Hash(networkPassPhrase)

	if err != nil {
		return "", errors.New("could not hash txn")
	}

	signature, err := kp.SignBase64(bytes[:])

	if err != nil {
		return "", err
	}

	return signature, nil

}

//VerifySignatureString verifies if the signatures match with the one to eb generated from toSign
func VerifySignatureString(toSign string, base64Signature string, publicKey string) error {
	kp, errParsingPublicKey := keypair.ParseAddress(publicKey)
	toSign = strings.TrimSpace(toSign)
	if errParsingPublicKey != nil {
		return &bantupaysdkErrors.ErrorInvalidPublicKey{}
	}

	providedSignature, errDecoding := base64.StdEncoding.DecodeString(base64Signature)

	if errDecoding != nil {
		return &bantupaysdkErrors.ErrorInvalidAuthenticationSignature{}
	}

	signatureError := kp.Verify([]byte(toSign), providedSignature)

	if signatureError != nil {
		log.Printf("invalid signature: %s\n", signatureError)
		return &bantupaysdkErrors.ErrorInvalidAuthenticationSignature{}
	}

	return nil

}

//VerifyHttpSignature verifies the httpRequest signation retrieved from X-BANTUPAY-SIGNATURE header
func VerifyHttpSignature(fullPathWithQuery string, body string, base64Signature string, publicKey string) error {
	body = strings.TrimSpace(body)
	fullPathWithQuery = strings.TrimSpace(fullPathWithQuery)
	signatureError := VerifySignatureString(fullPathWithQuery+body, base64Signature, publicKey)

	if signatureError != nil {
		log.Printf("invalid signature: %s\n", signatureError)
		return &bantupaysdkErrors.ErrorInvalidAuthenticationSignature{}
	}

	return nil

}

//NewPayment returns a new payment struct
func NewPayment() (p *PaymentInfo) {
	return new(PaymentInfo)
}
