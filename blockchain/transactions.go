// jeffCoin 1. BLOCKCHAIN transactions.go

package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strconv"

	wallet "github.com/JeffDeCola/jeffCoin/wallet"

	log "github.com/sirupsen/logrus"
)

// TRANSACTIONS **********************************************************************************************************

// transactionRequest - Request to transfer jeffCoins to a jeffCoin Address
func transactionRequest(transactionRequestMessageSigned string) string {

	s := "START  transactionRequest() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("TRANSACTION:          " + s)

	// PLACE THIS IS A STRUCT theTransactionRequestMessageStruct
	transactionRequestMessageSignedByte := []byte(transactionRequestMessageSigned)
	var theTransactionRequestMessageStruct transactionRequestMessageStruct
	err := json.Unmarshal(transactionRequestMessageSignedByte, &theTransactionRequestMessageStruct)
	checkErr(err)

	// EXTRACT WHAT YOU NEED
	signature := theTransactionRequestMessageStruct.Signature
	theRequestMessageStruct := theTransactionRequestMessageStruct.RequestMessage
	theRequestMessageByte, err := json.Marshal(theRequestMessageStruct)
	checkErr(err)
	theRequestMessage := string(theRequestMessageByte)

	// GET THE PUBLIC KEY ?????????????????????????????????????????????
	// FIX THIS THE RIGHT WAY
	theWallet := wallet.GetWallet()
	privateKeyHex := theWallet.PrivateKeyHex
	publicKeyHex := theWallet.PublicKeyHex
	_, publicKeyRawCheck := wallet.DecodeKeys(privateKeyHex, publicKeyHex)

	// VERIFY THIS IS FROM THE SENDER
	verifyStatus := verifySignature(publicKeyRawCheck, signature, theRequestMessage)

	// CHECK BALANCE???????????????????

	// ADD TRANSACTIONS TO pendingBlock???????????????????

	status := "Was this verified?: " + strconv.FormatBool(verifyStatus)

	s = "END    transactionRequest() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("TRANSACTION:          " + s)

	return status

}

// SIGNATURE  ************************************************************************************************************

// verifySignature - Verifies a ECDSA Digital Signature
func verifySignature(senderPublicKeyRaw *ecdsa.PublicKey, signature string, plainText string) bool {

	s := "START  verifySignature() - Verifies a ECDSA Digital Signature"
	log.Trace("TRANSACTION:          " + s)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	// DECODE signature
	signatureByte, _ := hex.DecodeString(signature)

	// EXTRACT R & S
	rSign := big.NewInt(0)
	sSign := big.NewInt(0)
	sigLen := len(signatureByte)
	rSign.SetBytes(signatureByte[:(sigLen / 2)])
	sSign.SetBytes(signatureByte[(sigLen / 2):])

	// VERIFY SIGNATURE
	verifyStatus := ecdsa.Verify(
		senderPublicKeyRaw,
		hashedPlainTextByte,
		rSign,
		sSign,
	)

	s = "Verified status is: " + strconv.FormatBool(verifyStatus)
	log.Info("TRANSACTION:                 " + s)

	s = "END    verifySignature() - Verifies a ECDSA Digital Signature"
	log.Trace("TRANSACTION:          " + s)

	return verifyStatus

}
