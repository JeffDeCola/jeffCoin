// jeffCoin 1. BLOCKCHAIN transactions.go

package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"strconv"

	wallet "github.com/JeffDeCola/jeffCoin/wallet"

	log "github.com/sirupsen/logrus"
)

// TRANSACTIONS **********************************************************************************************************

// processTxRequestMessage - Request to transfer jeffCoins to a jeffCoin Address
func (trms txRequestMessageSignedStruct) processTxRequestMessage() string {

	s := "START  processTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("TRANSACTION:          " + s)

	// PRINT OUT TX REQUEST MESSAGE
	fmt.Printf("\nThe theTxRequestMessageSignedStruct:\n\n")
	js, _ := json.MarshalIndent(trms, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// EXTRACT WHAT YOU NEED
	signature := trms.Signature
	theTxRequestMessageStruct := trms.TxRequestMessage
	theTxRequestMessageByte, err := json.Marshal(theTxRequestMessageStruct)
	checkErr(err)
	theTxRequestMessage := string(theTxRequestMessageByte)

	// ---------------------------------------------------------------------------
	// STEP 1 - VERIFY SIGNATURE

	// GET THE PUBLIC KEY (ADDRESS) FROM BLOCKCHAIN ???????????????????????????UPDATE THIS
	theWallet := wallet.GetWallet()
	publicKeyHex := theWallet.PublicKeyHex

	// VERIFY THIS IS FROM THE SENDER
	verifyStatus := verifySignature(publicKeyHex, signature, theTxRequestMessage)
	status := strconv.FormatBool(verifyStatus)

	// ---------------------------------------------------------------------------
	// STEP 2 - CHECK BALANCE TO SEE IF YOU HAVE THE MONEY
	//getBalance()

	// ---------------------------------------------------------------------------
	// STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins

	// ---------------------------------------------------------------------------
	// STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE
	//pickUnspentOutputs()

	// ---------------------------------------------------------------------------
	// STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE
	//addTransactionToPendingBlock()

	// ????????

	// CHECK BALANCE???????????????????

	// ADD TRANSACTIONS TO pendingBlock???????????????????

	s = "END    processTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("TRANSACTION:          " + s)

	return status

}

// SIGNATURE  ************************************************************************************************************

// verifySignature - Verifies a ECDSA Digital Signature
func verifySignature(publicKeyHex string, signature string, plainText string) bool {

	s := "START  verifySignature() - Verifies a ECDSA Digital Signature"
	log.Trace("TRANSACTION:          " + s)

	// DECODE PUBLIC KEY
	publicKeyPEM, _ := hex.DecodeString(publicKeyHex)
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	publicKeyx509Encoded := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(publicKeyx509Encoded)
	publicKeyRaw := genericPublicKey.(*ecdsa.PublicKey)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	// DECODE signature
	signatureByte, _ := hex.DecodeString(signature)

	// EXTRACT R & S
	r := big.NewInt(0)
	ss := big.NewInt(0)
	sigLen := len(signatureByte)
	r.SetBytes(signatureByte[:(sigLen / 2)])
	ss.SetBytes(signatureByte[(sigLen / 2):])

	// VERIFY SIGNATURE
	verifyStatus := ecdsa.Verify(
		publicKeyRaw,
		hashedPlainTextByte,
		r,
		ss,
	)

	s = "Verified status is: " + strconv.FormatBool(verifyStatus)
	log.Info("TRANSACTION:                 " + s)

	s = "END    verifySignature() - Verifies a ECDSA Digital Signature"
	log.Trace("TRANSACTION:          " + s)

	return verifyStatus

}
