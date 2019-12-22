// jeffCoin TESTDATA load-test-data.go

package testdata

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// LoadTestDatatoBlockchain - Load the blockchain with test data
func LoadTestDatatoBlockchain() {

	s := "START  LoadTestDatatoBlockchain() - Load the blockchain with test data"
	log.Trace("TRANSACTION:          " + s)

	// RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "RECEIVING SOME TRANSACTION REQUEST MESSAGES"
	log.Info("LoadTestDatatoBlockchain()                " + s)
	receivingTransaction(txRequestMessageSignedDataString1)
	receivingTransaction(txRequestMessageSignedDataStringBad)

	// s = "Verified status is: " + strconv.FormatBool(verifyStatus)
	// log.Info("TRANSACTION:                 " + s)

	s = "END    LoadTestDatatoBlockchain() - Load the blockchain with test data"
	log.Trace("TRANSACTION:          " + s)
}

func receivingTransaction(txRequestMessageSignedDataString string) {

	s := "------------------------------------------------------------------"
	log.Info("receivingTransaction()           " + s)

	//var trms blockchain.txRequestMessageSignedStruct

	// RECEIVED TRANSACTION MESSAGE txRequestMessageSignedDataStringBad
	// Place transaction Request Message data in transaction Request Message struct
	//txRequestMessageSignedDataStringByte := []byte(txRequestMessageSignedDataString)
	//err := json.Unmarshal(txRequestMessageSignedDataStringByte, &trms)
	//checkErr(err)

	// Check is message valid, get balance and add to pendingBlock
	//s = "Received transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
	//		fmt.Sprint(trms.TxRequestMessage.Destinations)
	//	log.Info("receivingTransaction()           " + s)
	// status := trms.processTransactionRequest()
	//s = "The status of transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
	//	fmt.Sprint(trms.TxRequestMessage.Destinations) + " is " + status
	//log.Info("receivingTransaction()           " + s)

}
