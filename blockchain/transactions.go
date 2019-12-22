// jeffCoin 1. BLOCKCHAIN transactions.go

package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	wallet "github.com/JeffDeCola/jeffCoin/wallet"

	log "github.com/sirupsen/logrus"
)

// TRANSACTIONS **********************************************************************************************************

// processTxRequestMessage - Request to transfer jeffCoins to a jeffCoin Address
func processTxRequestMessage(txRequestMessageSigned string) string {

	s := "START  processTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("TRANSACTION:          " + s)

	// PLACE THIS IS A STRUCT theTxnRequestMessageStruct
	txRequestMessageSignedByte := []byte(txRequestMessageSigned)
	var theTxRequestMessageSignedStruct txRequestMessageSignedStruct
	err := json.Unmarshal(txRequestMessageSignedByte, &theTxRequestMessageSignedStruct)
	checkErr(err)

	// PRINT IT OUT
	fmt.Printf("\nThe theTxRequestMessageSignedStruct:\n\n")
	js, _ := json.MarshalIndent(theTxRequestMessageSignedStruct, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// EXTRACT WHAT YOU NEED
	signature := theTxRequestMessageSignedStruct.Signature
	theTxRequestMessageStruct := theTxRequestMessageSignedStruct.TxRequestMessage
	theTxRequestMessageByte, err := json.Marshal(theTxRequestMessageStruct)
	checkErr(err)
	theTxRequestMessage := string(theTxRequestMessageByte)

	// GET THE PUBLIC KEY FROM BLOCKCHAIN ???????????????????????????
	theWallet := wallet.GetWallet()
	privateKeyHex := theWallet.PrivateKeyHex
	publicKeyHex := theWallet.PublicKeyHex
	_, publicKeyRawCheck := wallet.DecodeKeys(privateKeyHex, publicKeyHex)

	// VERIFY THIS IS FROM THE SENDER
	verifyStatus := verifySignature(publicKeyRawCheck, signature, theTxRequestMessage)

	// CHECK BALANCE???????????????????

	// ADD TRANSACTIONS TO pendingBlock???????????????????

	status := "Was this verified?: " + strconv.FormatBool(verifyStatus)

	s = "END    processTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
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
