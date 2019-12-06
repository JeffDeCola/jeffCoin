// jeffCoin 1. BLOCKCHAIN blockchain-interface.go

package blockchain

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

// BLOCKCHAIN ************************************************************************************************************

// GetBlockchain - Gets the blockchain
func GetBlockchain() blockchainSlice {

	s := "START  GetBlockchain() - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlockchain := getBlockchain()

	s = "END    GetBlockchain() - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlockchain

}

// GenesisBlockchain - Creates the blockchain
func GenesisBlockchain(transaction string, difficulty int) {

	s := "START  GenesisBlockchain() - Creates the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	resetCurrentBlock(transaction)
	lockCurrentBlock(difficulty)
	newBlock := appendLockedBlock()

	fmt.Printf("\nCongrats, your first block in your blockchain is:\n\n")
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END    GenesisBlockchain() - Creates the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

}

// RequestBlockchain - Requests the blockchain and the currentBlock from a Network Node
func RequestBlockchain(networkIP string, networkTCPPort string) error {

	s := "START  RequestBlockchain() - Requests the blockchain and the currentBlock from a Network Node"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	// SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F           " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("BLOCKCHAIN:  I/F           " + s)
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F           " + s)
	s = "-conn   TCP Connection on " + networkIP + ":" + networkTCPPort
	log.Info("BLOCKCHAIN:  I/F   " + s)
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for command"
		log.Error("BLOCKCHAIN:  I/F           " + s)
		return errors.New(s)
	}

	// SEND-BLOCKCHAIN
	s = "-req    - SEND-BLOCKCHAIN"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	fmt.Fprintf(conn, "SEND-BLOCKCHAIN\n")

	// GET THE blockchain
	messageBlockchain, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + messageBlockchain
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if messageBlockchain == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("BLOCKCHAIN:  I/F            " + s)
		return errors.New(s)
	}

	// LOAD THE blockchain
	loadBlockchain(messageBlockchain)

	// GET THE currentBlock
	messageCurrentBlock, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + messageCurrentBlock
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if messageCurrentBlock == "ERROR" {
		s = "ERROR: Could not get currentBlock from node"
		log.Error("BLOCKCHAIN:  I/F            " + s)
		return errors.New(s)
	}

	// LOAD THE CurrentBlock
	loadCurrentBlock(messageCurrentBlock)

	// GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("BLOCKCHAIN:  I/F            " + s)
		return errors.New(s)
	}

	// EOF (CLOSE CONNECTION)
	s = "-req    - EOF (CLOSE CONNECTION)"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F           " + s)
	s = "CLIENT - Closed a connection"
	log.Info("BLOCKCHAIN:  I/F           " + s)
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F           " + s)

	s = "END    RequestBlockchain() - Requests the blockchain and the currentBlock from a Network Node"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return nil

}

// BLOCK *****************************************************************************************************************

// GetBlock - Gets a block (via Index number) from the blockchain
func GetBlock(id string) blockStruct {

	s := "START  GetBlock() - Gets a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getBlock(id)

	// RETURN NOT FOUND
	s = "END    GetBlock() - Gets a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// LOCKED BLOCK **********************************************************************************************************

// GetLockedBlock - Gets the lockedBlock
func GetLockedBlock() blockStruct {

	s := "START  GetLockedBlock() - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getLockedBlock()

	// RETURN NOT FOUND
	s = "END    GetLockedBlock() - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// CURRENT BLOCK *********************************************************************************************************

// GetCurrentBlock - Gets the currentBlock
func GetCurrentBlock() blockStruct {

	s := "START  GetCurrentBlock() - Gets the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getCurrentBlock()

	// RETURN NOT FOUND
	s = "END    GetCurrentBlock() - Gets the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// AddTransactionToCurrentBlock - Adds a transaction to the currentBlock
func AddTransactionToCurrentBlock(transaction string) blockStruct {

	s := "START  AddTransactionToCurrentBlock() - Adds a transaction to the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theCurrentBlock := addTransactionToCurrentBlock(transaction)

	s = "END    AddTransactionToCurrentBlock() - Adds a transaction to the currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theCurrentBlock

}

// JEFFCOINS *************************************************************************************************************

// GetAddressBalance - Gets the jeffCoin Address balance
func GetAddressBalance(jeffCoinAddress string) string {

	s := "START  GetAddressBalance() - Gets the jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	balance := getAddressBalance(jeffCoinAddress)

	s = "END    GetAddressBalance() - Gets the jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return balance

}

// TRANSACTIONS **********************************************************************************************************

// TransactionRequest - Request to transfer jeffCoins to a jeffCoin Address
func TransactionRequest(transactionRequestMessageSigned string) string {

	s := "START  TransactionRequest() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	status := transactionRequest(transactionRequestMessageSigned)

	s = "END    TransactionRequest() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return status

}
