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

	log "github.com/sirupsen/logrus"
)

// TRANSACTIONS **********************************************************************************************************

// processTxRequestMessage - Request to transfer jeffCoins to a jeffCoin Address
func (trms txRequestMessageSignedStruct) processTxRequestMessage() string {

	s := "START  processTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Debug("TRANSACTION:          " + s)

	// PRINT OUT TX REQUEST MESSAGE
	s = "The theTxRequestMessageSignedStruct (-loglevel trace to display)"
	log.Info("TRANSACTION:                 " + s)
	js, _ := json.MarshalIndent(trms, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	// EXTRACT WHAT YOU NEED
	signature := trms.Signature
	publicKeyHex := trms.TxRequestMessage.SourceAddress
	theTxRequestMessageStruct := trms.TxRequestMessage
	theTxRequestMessageByte, err := json.Marshal(theTxRequestMessageStruct)
	checkErr(err)
	theTxRequestMessage := string(theTxRequestMessageByte)

	// ---------------------------------------------------------------------------
	// STEP 1 - VERIFY SIGNATURE
	s = "STEP 1 - VERIFY SIGNATURE"
	log.Info("TRANSACTION:                 " + s)
	verifyStatus := verifySignature(publicKeyHex, signature, theTxRequestMessage)
	if !(verifyStatus) {
		s = "Signature Failed"
		log.Warn("TRANSACTION:                 " + s)
		return "Signature Failed"
	}

	// ---------------------------------------------------------------------------
	// STEP 2 - GET BALANCE AND A LIST OF UNSPENT OUTPUTS
	s = "STEP 2 - GET BALANCE AND A LIST OF UNSPENT OUTPUTS"
	log.Info("TRANSACTION:                 " + s)
	balance, unspentOutput := getAddressBalance(trms.TxRequestMessage.SourceAddress)

	// ---------------------------------------------------------------------------
	// STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins
	s = "STEP 3 - CHECK IF YOU HAVE ENOUGH jeffCoins"
	log.Info("TRANSACTION:                 " + s)
	var value int64
	for _, destinations := range trms.TxRequestMessage.Destinations {
		value = value + destinations.Value
	}
	s = "The balance for " + trms.TxRequestMessage.SourceAddress[0:50] + "... " + "is " + strconv.FormatInt(balance, 10)
	log.Info("TRANSACTION:                 " + s)
	s = "The value to remove is " + strconv.FormatInt(value, 10) + " from " + fmt.Sprint(unspentOutput)
	log.Info("TRANSACTION:                 " + s)
	if balance < value {
		s = "Not Enough jeffCoins/Value"
		log.Warn("TRANSACTION:                 " + s)
		return "Not Enough jeffCoins/Value"
	}

	// ---------------------------------------------------------------------------
	// STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE
	s = "STEP 4 - PICK THE UNSPENT OUTPUTS TO USE AND PROVIDE CHANGE"
	log.Info("TRANSACTION:                 " + s)
	useUnspentOutput, change := pickUnspentOutputs(unspentOutput, value)
	s = "You are using unspent outputs " + fmt.Sprint(useUnspentOutput)
	log.Info("TRANSACTION:                 " + s)
	s = "The change will be " + strconv.FormatInt(change, 10)
	log.Info("TRANSACTION:                 " + s)

	// ---------------------------------------------------------------------------
	// STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE
	s = "STEP 5 - ADD TRANSACTION to pendingBlock and MAKE CHANGE"
	log.Info("TRANSACTION:                 " + s)
	trms.addTransactionToPendingBlock(useUnspentOutput, change)

	s = "END    processTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Debug("TRANSACTION:          " + s)

	return "Pending Transaction"

}

// SIGNATURE *************************************************************************************************************

// verifySignature - Verifies a ECDSA Digital Signature
func verifySignature(publicKeyHex string, signature string, plainText string) bool {

	s := "START  verifySignature() - Verifies a ECDSA Digital Signature"
	log.Debug("TRANSACTION:          " + s)

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
	log.Debug("TRANSACTION:          " + s)

	return verifyStatus

}

// UNSPENT OUTPUTS *******************************************************************************************************

// pickUnspentOutputs - Pick the Unspent Outputs to use and provide change
func pickUnspentOutputs(pickUnspentOutputSlice []unspentOutputStruct, value int64) ([]unspentOutputStruct, int64) {

	s := "START  pickUnspentOutputs() - Pick the Unspent Outputs to use and provide change"
	log.Debug("TRANSACTION:          " + s)

	var unspentOutputStructTemp = unspentOutputStruct{}
	var useUnspentOutputSlice []unspentOutputStruct

	var change int64
	var runningTotal int64

	// Once you hit the value, stop
	for _, unspentOutput := range pickUnspentOutputSlice {

		unspentOutputStructTemp.TxID = unspentOutput.TxID
		unspentOutputStructTemp.Value = unspentOutput.Value

		// Place in slice
		useUnspentOutputSlice = append(useUnspentOutputSlice, unspentOutputStructTemp)

		runningTotal = runningTotal + unspentOutput.Value

		// did you get enough - If yes, provide change
		if value < runningTotal {
			change = runningTotal - value
			break
		}

	}

	s = "END    pickUnspentOutputs() - Pick the Unspent Outputs to use and provide change"
	log.Debug("TRANSACTION:          " + s)

	return useUnspentOutputSlice, change

}
