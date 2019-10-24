// jeffCoin transactions.go

package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"
)

// TRANSACTIONS ************************************************************************************************************

// transactionRequest - Request to Transfer Coins to a jeffCoin Address
func transactionRequest(transactionRequestMessageSigned string) string {

	s := "START: transactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("TRANSACTION:        " + s)

	// PLACE THIS IS A STRUCT theTransactionRequestMessageStruct
	transactionRequestMessageSignedByte := []byte(transactionRequestMessageSigned)
	var theTransactionRequestMessageStruct transactionRequestMessageStruct
	err := json.Unmarshal(transactionRequestMessageSignedByte, &theTransactionRequestMessageStruct)
	checkErr(err)

    // EXTRACT WHAT YOU NEED
    signature := theTransactionRequestMessageStruct.Signature
    theRequestMessageStruct := theTransactionRequestMessageStruct.RequestMessage
    theRequestMessageByte:= json.Marshal(theRequestMessageStruct)
    plainText := string(theRequestMessageByte)

    // GET THE KEY
    
    // VERIFY THIS IS FROM THE SENDER
    verifyStatus := verifySignature(senderPublicKeyRaw *ecdsa.PublicKey, signature string, plainText string) {

	// ADD TRANSACTIONS TO CURRENT BLOCK

	status := "Was this verified?: " + strconv.FormatBool(verifyStatus)

	s = "END:   transactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("TRANSACTION:        " + s)

	return status

}

// SIGNATURE  ************************************************************************************************************

// verifySignature - Verify a ECDSA Digital Signature
func verifySignature(senderPublicKeyRaw *ecdsa.PublicKey, signature string, plainText string) bool {

	s := "START: verifySignature - Verify a ECDSA Digital Signature"
	log.Trace("TRANSACTION:        " + s)

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

	s = "END:   verifySignature - Verify a ECDSA Digital Signature"
	log.Trace("TRANSACTION:        " + s)

	return verifyStatus

}
