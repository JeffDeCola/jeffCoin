// jeffCoin TESTDATA load-test-data.go

package testdata

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"
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
	mockReceivingTransaction(txRequestMessageSignedDataString1)

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

	// MOCK - RESET pendingBlock
	s = "MOCK - RESET pendingBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.ResetPendingBlock()

	// MOCK - RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "MOCK - RECEIVING TRANSACTION REQUEST MESSAGES"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	mockReceivingTransaction(txRequestMessageSignedDataString2)

	// MOCK - lockPendingBlock() - Move pendingBlock to lockedBlock
	s = "MOCK - lockPendingBlock() - Move pendingBlock to lockedBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	// GET DIFFICULTY FROM LAST LOCKED BLOCK
	theLockedBlock = blockchain.GetLockedBlock()
	blockchain.LockPendingBlock(theLockedBlock.Difficulty)

	// MOCK - ADD lockedBlock TO THE blockchain
	theLockedBlock = blockchain.GetLockedBlock()
	s = "MOCK - ADD lockedBlock TO THE blockchain. Adding block number " + fmt.Sprint(theLockedBlock.BlockID)
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.AppendLockedBlock()

	// MOCK - RESET pendingBlock
	s = "MOCK - RESET pendingBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.ResetPendingBlock()

	// MOCK - RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "MOCK - RECEIVING TRANSACTION REQUEST MESSAGES"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	mockReceivingTransaction(txRequestMessageSignedDataString3)
	mockReceivingTransaction(txRequestMessageSignedDataString4)

	// MOCK - lockPendingBlock() - Move pendingBlock to lockedBlock
	s = "MOCK - lockPendingBlock() - Move pendingBlock to lockedBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	// GET DIFFICULTY FROM LAST LOCKED BLOCK
	theLockedBlock = blockchain.GetLockedBlock()
	blockchain.LockPendingBlock(theLockedBlock.Difficulty)

	// MOCK - ADD lockedBlock TO THE blockchain
	theLockedBlock = blockchain.GetLockedBlock()
	s = "MOCK - ADD lockedBlock TO THE blockchain. Adding block number " + fmt.Sprint(theLockedBlock.BlockID)
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.AppendLockedBlock()

	// MOCK - RESET pendingBlock
	s = "MOCK - RESET pendingBlock"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	blockchain.ResetPendingBlock()

	// MOCK - RECEIVING SOME TRANSACTION REQUEST MESSAGES
	s = "MOCK - RECEIVING TRANSACTION REQUEST MESSAGES"
	log.Trace("*** LOAD-TEST-DATA:          " + s)
	mockReceivingTransaction(txRequestMessageSignedDataString5)
	mockReceivingTransaction(txRequestMessageSignedDataString6)

	s = "END    LoadTestDatatoBlockchain() - Load the blockchain with test data"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

}

// mockReceivingTransaction - Place the transaction into txRequestMessageSignedStruct and process
func mockReceivingTransaction(txRequestMessageSignedDataString string) {

	s := "START  mockReceivingTransaction() - Place the transaction into txRequestMessageSignedStruct and process"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

	// Remove all spaces and returns, etc... - squish it
	// Make a long string - Remove /n and whitespace
	txRequestMessageSignedDataString = strings.Replace(txRequestMessageSignedDataString, "\n", "", -1)
	txRequestMessageSignedDataString = strings.Replace(txRequestMessageSignedDataString, " ", "", -1)

	s = "--------------------------------------------"
	log.Trace("*** LOAD-TEST-DATA:          " + s)

	// GET nodeIP & nodeTCPPort from thisNode
	thisNode := routingnode.GetThisNode()
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// REQUEST TRANSACTION
	status, err := wallet.TransactionRequest(nodeIP, nodeTCPPort, txRequestMessageSignedDataString)
	checkErr(err)

	s = "Status is " + status
	log.Trace("*** LOAD-TEST-DATA:          " + s)

	// Check is message valid, get balance and add to pendingBlock
	//s = "Received transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
	//		fmt.Sprint(trms.TxRequestMessage.Destinations)
	//	log.Info("mockReceivingTransaction()           " + s)
	// status := trms.processTransactionRequest()
	//s = "The status of transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
	//	fmt.Sprint(trms.TxRequestMessage.Destinations) + " is " + status
	//log.Info("mockReceivingTransaction()           " + s)

	s = "END    mockReceivingTransaction() - Place the transaction into txRequestMessageSignedStruct and process"
	log.Trace("*** LOAD-TEST-DATA:   " + s)

}
