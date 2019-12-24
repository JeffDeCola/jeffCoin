// jeffCoin 1. BLOCKCHAIN blockchain-interface.go

package blockchain

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// BLOCKCHAIN ************************************************************************************************************

// GetBlockchain - Gets the blockchain
func GetBlockchain() blockchainSlice {

	s := "START  GetBlockchain() - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	theBlockchain := getBlockchain()

	s = "END    GetBlockchain() - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return theBlockchain

}

// GenesisBlockchain - Creates the blockchain
func GenesisBlockchain(blockDataString string) {

	s := "START  GenesisBlockchain() - Creates the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	// LOAD pendingBlock - loadPendingBlock()
	s = "LOAD pendingBlock - loadPendingBlock()"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	loadPendingBlock(blockDataString)

	// LOCK pendingBlock - lockPendingBlock()
	s = "LOCK pendingBlock - lockPendingBlock()"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	lockPendingBlock(pendingBlock.Difficulty)

	// APPEND lockBlock - appendLockedBlock()
	s = "APPEND lockBlock - appendLockedBlock()"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	appendLockedBlock()

	firstBlock := getBlock("0")

	fmt.Printf("\nCongrats, your first block in your blockchain is:\n\n")
	js, _ := json.MarshalIndent(firstBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	// RESET pendingBlock - resetPendingBlock()
	s = "RESET pendingBlock - resetPendingBlock()"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	resetPendingBlock()

	s = "END    GenesisBlockchain() - Creates the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

}

// RequestBlockchain - Requests the blockchain and the pendingBlock from a Network Node
func RequestBlockchain(networkIP string, networkTCPPort string) error {

	s := "START  RequestBlockchain() - Requests the blockchain and the pendingBlock from a Network Node"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	//  CONN - SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	s = "-C conn   TCP Connection on " + networkIP + ":" + networkTCPPort
	log.Info("BLOCKCHAIN:  I/F   " + s)
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for command"
		log.Error("BLOCKCHAIN:  I/F             " + s)
		return errors.New(s)
	}

	// REQ - SEND-BLOCKCHAIN
	s = "-C req    - SEND-BLOCKCHAIN"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	fmt.Fprintf(conn, "SEND-BLOCKCHAIN\n")

	// RCV - GET THE blockchain
	messageBlockchain, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: (NOT SHOWN) - Received blockchain"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if messageBlockchain == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("BLOCKCHAIN:  I/F              " + s)
		return errors.New(s)
	}

	// LOAD THE blockchain
	loadBlockchain(messageBlockchain)

	// SEND - THANK YOU
	s = "-C send   - Thank you"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	fmt.Fprintf(conn, "Thank You\n")

	// RCV - GET THE pendingBlock
	messagePendingBlock, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: (NOT SHOWN) - Received pendingBlock"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if messagePendingBlock == "ERROR" {
		s = "ERROR: Could not get pendingBlock from node"
		log.Error("BLOCKCHAIN:  I/F              " + s)
		return errors.New(s)
	}

	// LOAD THE PendingBlock
	loadPendingBlock(messagePendingBlock)

	// SEND - THANK YOU
	s = "-C send   - Thank you"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	fmt.Fprintf(conn, "Thank You\n")

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("BLOCKCHAIN:  I/F              " + s)
		return errors.New(s)
	}

	// REQ - EOF (CLOSE CONNECTION)
	s = "-C req    - EOF (CLOSE CONNECTION)"
	log.Info("BLOCKCHAIN:  I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	s = "CLIENT - Closed a connection"
	log.Info("BLOCKCHAIN:  I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("BLOCKCHAIN:  I/F             " + s)

	s = "END    RequestBlockchain() - Requests the blockchain and the pendingBlock from a Network Node"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return nil

}

// BLOCK *****************************************************************************************************************

// GetBlock - Gets a block (via Index number) from the blockchain
func GetBlock(id string) blockStruct {

	s := "START  GetBlock() - Gets a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	theBlock := getBlock(id)

	// RETURN NOT FOUND
	s = "END    GetBlock() - Gets a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return theBlock

}

// LOCKED BLOCK **********************************************************************************************************

// GetLockedBlock - Gets the lockedBlock
func GetLockedBlock() blockStruct {

	s := "START  GetLockedBlock() - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	theBlock := getLockedBlock()

	s = "END    GetLockedBlock() - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return theBlock

}

// AppendLockedBlock - Appends the lockedBlock to the blockchain
func AppendLockedBlock() {

	s := "START  AppendLockedBlock() - Appends the lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	appendLockedBlock()

	s = "END    AppendLockedBlock() - Appends the lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

}

// PENDING BLOCK *********************************************************************************************************

// GetPendingBlock - Gets the pendingBlock
func GetPendingBlock() blockStruct {

	s := "START  GetPendingBlock() - Gets the pendingBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	theBlock := getPendingBlock()

	s = "END    GetPendingBlock() - Gets the pendingBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return theBlock

}

// ResetPendingBlock - Resets the pendingBlock
func ResetPendingBlock() {

	s := "START  ResetPendingBlock() - Resets the pendingBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	resetPendingBlock()

	s = "END    ResetPendingBlock() - Resets the pendingBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

}

// AddTransactionToPendingBlock - Adds a transaction to the pendingBlock
func AddTransactionToPendingBlock(transaction string) blockStruct {

	s := "START  AddTransactionToPendingBlock() - Adds a transaction to the pendingBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	thePendingBlock := addTransactionToPendingBlock(transaction)

	s = "END    AddTransactionToPendingBlock() - Adds a transaction to the pendingBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return thePendingBlock

}

// LockPendingBlock - Moves the pendingBlock to the lockedBlock
func LockPendingBlock(difficulty int) {

	s := "START  LockPendingBlock() - Moves the pendingBlock to the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	lockPendingBlock(difficulty)

	s = "END    LockPendingBlock() - Moves the pendingBlock to the lockedBlock"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

}

// JEFFCOINS *************************************************************************************************************

// GetAddressBalance - Gets the jeffCoin Address balance
func GetAddressBalance(jeffCoinAddress string) string {

	s := "START  GetAddressBalance() - Gets the jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	balance, _ := getAddressBalance(jeffCoinAddress)

	s = "END    GetAddressBalance() - Gets the jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return strconv.FormatInt(balance, 10)

}

// TRANSACTIONS **********************************************************************************************************

// ProcessTxRequestMessage - Request to transfer jeffCoins to a jeffCoin Address
func ProcessTxRequestMessage(txRequestMessageSigned string) string {

	s := "START  ProcessTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	var trms txRequestMessageSignedStruct

	// PLACE txRequestMessageSigned IN A STRUCT theTxnRequestMessageStruct
	txRequestMessageSignedByte := []byte(txRequestMessageSigned)
	err := json.Unmarshal(txRequestMessageSignedByte, &trms)
	checkErr(err)

	// PRINT NOTE
	s = "Received transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
		fmt.Sprint(trms.TxRequestMessage.Destinations)
	log.Info("BLOCKCHAIN:  I/F             " + s)

	// CREATE SIG BASED ON PRIVATE KEY
	//gotWallet := wallet.GetWallet()
	//sourceAddress := gotWallet.JeffCoinAddress
	// GET ENCODED KEYS FROM wallet
	//privateKeyHex := gotWallet.PrivateKeyHex
	// Make a long string - Remove /n and whitespace
	//txRequestMessageStruct := trms.TxRequestMessage
	//txRequestMessage, _ := json.Marshal(txRequestMessageStruct)
	//signature := wallet.CreateSignature(privateKeyHex, string(txRequestMessage))
	//fmt.Printf("\n\n%v\n\n", sourceAddress)
	//fmt.Printf("\n\n%v\n\n", signature)
	//time.Sleep(100000 * time.Second)

	// PROCESS
	status := trms.processTxRequestMessage()

	// PRINT STATUS
	s = "The status of transaction message from " + trms.TxRequestMessage.SourceAddress + " to " +
		fmt.Sprint(trms.TxRequestMessage.Destinations) + " is " + status
	log.Info("BLOCKCHAIN:  I/F             " + s)

	s = "END    ProcessTxRequestMessage() - Request to transfer jeffCoins to a jeffCoin Address"
	log.Trace("BLOCKCHAIN:  I/F      " + s)

	return status

}
