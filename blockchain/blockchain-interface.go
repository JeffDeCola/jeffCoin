// jeffCoin blockchain-interface.go

package blockchain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	errors "github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// BLOCKCHAIN ************************************************************************************************************

// GenesisBlockchain - Creates the blockchain
func GenesisBlockchain(transaction string, difficulty int) {

	s := "START: GenesisBlockchain - Creates the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	resetCurrentBlock(transaction)
	lockCurrentBlock(difficulty)
	newBlock := appendLockedBlock()
	// resetCurrentBlock(transaction)

	fmt.Printf("\nCongrats, your first block in your blockchain is:\n\n")
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END:   GenesisBlockchain - Creates the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

}

// LoadBlockchain - Receives the blockchain and the currentBlock
func LoadBlockchain(networkIP string, networkTCPPort string) error {

	s := "START: LoadBlockchain - Receives the blockchain and the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F           " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("BLOCKCHAIN:  I/F           " + s)
		return errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "SEND-BLOCKCHAIN\n")

	// GET THE blockchain
	messageBlockchain, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F           " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("BLOCKCHAIN:  I/F           " + s)
		return errors.New(s)
	}

	// GET THE currentBlock
	messageCurrentBlock, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F           " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("BLOCKCHAIN:  I/F           " + s)
		return errors.New(s)
	}

	// LOAD THE blockchain
	loadBlockchain(messageBlockchain)

	// LOAD THE CurrentBlock
	loadCurrentBlock(messageCurrentBlock)

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END:   LoadBlockchain - Receives the blockchain and the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return nil

}

// GetBlockchain - Gets the blockchain
func GetBlockchain() blockchainSlice {

	s := "START: GetBlockchain - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlockchain := getBlockchain()

	s = "END:   GetBlockchain - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlockchain

}

// BLOCK *****************************************************************************************************************

// GetBlock - Gets a block (via Index number) from the blockchain
func GetBlock(id string) blockStruct {

	s := "START: GetBlock - Gets a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getBlock(id)

	// RETURN NOT FOUND
	s = "END:   GetBlock - Gets a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// LOCKED BLOCK **********************************************************************************************************

// GetLockedBlock - Gets the lockedBlock
func GetLockedBlock() blockStruct {

	s := "START: GetLockedBlock - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getLockedBlock()

	// RETURN NOT FOUND
	s = "END:   GetLockedBlock - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// CURRENT BLOCK *********************************************************************************************************

// GetCurrentBlock - Gets the currentBlock
func GetCurrentBlock() blockStruct {

	s := "START: GetCurrentBlock - Gets the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getCurrentBlock()

	// RETURN NOT FOUND
	s = "END:   GetCurrentBlock - Gets the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// AddTransactionToCurrentBlock - Adds a transaction to the currentBlock
func AddTransactionToCurrentBlock(transaction string) blockStruct {

	s := "START: AddTransactionToCurrentBlock - Adds a transaction to the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theCurrentBlock := addTransactionToCurrentBlock(transaction)

	s = "END:   AddTransactionToCurrentBlock - Adds a transaction to the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theCurrentBlock

}

// COINS ****************************************************************************************************************

// GetAddressBalance - Gets jeffCoin Address balance
func GetAddressBalance(jeffCoinAddress string) string {

	s := "START: GetAddressBalance - Gets jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	balance := getAddressBalance(jeffCoinAddress)

	s = "END:   GetAddressBalance - Gets jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return balance

}

// TRANSACTIONS ****************************************************************************************************************

// TransactionRequest - Request to Transfer Coins to a jeffCoin Address
func TransactionRequest(transactionRequestMessageSigned string) string {

	s := "START: TransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	status := transactionRequest(transactionRequestMessageSigned)

	s = "END:   TransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return status

}
