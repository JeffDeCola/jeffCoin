// jeffCoin TESTDATA load-test-data.go

package testdata

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
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
	log.Trace("*** LOAD-TEST-DATA:   " + s)

	// MOCK - RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "MOCK - RECEIVING TRANSACTION REQUEST MESSAGES"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	receivingTransaction(txRequestMessageSignedDataString1)
	receivingTransaction(txRequestMessageSignedDataString2)

	// MOCK - lockPendingBlock() - Move pendingBlock to lockedBlock
	s = "MOCK - lockPendingBlock() - Move pendingBlock to lockedBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	// GET DIFFICULTY FROM LAST LOCKED BLOCK
	theLockedBlock := blockchain.GetLockedBlock()
	blockchain.LockPendingBlock(theLockedBlock.Difficulty)

	// MOCK - ADD lockedBlock TO THE blockchain
	theLockedBlock = blockchain.GetLockedBlock()
	s = "MOCK - ADD lockedBlock TO THE blockchain. Adding block number " + fmt.Sprint(theLockedBlock.BlockID)
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.AppendLockedBlock()

	// RESET pendingBlock
	s = "MOCK - RESET pendingBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.ResetPendingBlock()

	s = "END    LoadTestDatatoBlockchain() - Load the blockchain with test data"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

}

// receivingTransaction - Place the transaction into txRequestMessageSignedStruct and process
func receivingTransaction(txRequestMessageSignedDataString string) {

	s := "START  receivingTransaction() - Place the transaction into txRequestMessageSignedStruct and process"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

	s = "--------------------------------------------"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

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

	s = "END    receivingTransaction() - Place the transaction into txRequestMessageSignedStruct and process"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

}
