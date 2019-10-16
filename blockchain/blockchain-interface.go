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

// GenesisBlockchain - Creates the Blockchain (Only run once)
func GenesisBlockchain(transaction string, difficulty int) {

	s := "START: GenesisBlockchain - Creates the blockchain (Only run once)"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	refreshCurrentBlock(transaction)
	lockCurrentBlock(difficulty)
	newBlock := appendLockedBlock()
	refreshCurrentBlock("Refreshed")

	fmt.Printf("\nCongrats, your first block in your blockchain is:\n\n")
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END:   GenesisBlockchain - Creates the blockchain (Only run once)"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

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

// GetBlock - Get a Block (via Index number) from the blockchain
func GetBlock(id string) blockStruct {

	s := "START: GetBlock - Get a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theBlock := getBlock(id)

	// RETURN NOT FOUND
	s = "END:   GetBlock - Get a block (via Index number) from the blockchain"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theBlock

}

// LoadBlockchain - Loads the blockchain and CurrentBlock from a Network Node
func LoadBlockchain(networkIP string, networkTCPPort string) error {

	s := "START: LoadBlockchain - Loads the blockchain and CurrentBlock from a Network Node"
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
	fmt.Fprintf(conn, "SENDBLOCKCHAIN\n")

	// GET THE blockchain
	messageBlockchain, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("BLOCKCHAIN:  I/F           " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("BLOCKCHAIN:  I/F           " + s)
		return errors.New(s)
	}

	// GET THE CurrentBlock
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

	s = "END:   LoadBlockchain - Loads the blockchain and CurrentBlock from a Network Node"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return nil

}

// AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock
func AddTransactionToCurrentBlock(transaction string) blockStruct {

	s := "START: AddTransactionToCurrentBlock - Add a Transaction to currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	theCurrentBlock := addTransactionToCurrentBlock(transaction)

	s = "END:   AddTransactionToCurrentBlock - Add a Transaction to currentBlock"
	log.Trace("BLOCKCHAIN:  I/F    " + s)

	return theCurrentBlock

}
