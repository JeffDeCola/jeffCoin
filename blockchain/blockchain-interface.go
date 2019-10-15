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

	s := "START: GenesisBlockchain - Creates the Blockchain (Only run once)"
	log.Println("BLOCKCHAIN I/F:     " + s)

	refreshCurrentBlock(transaction)
	addTransactionCurrentBlock("remove me TEST") // REMOVE THIS
	lockCurrentBlock(difficulty)
	newBlock := appendLockedBlock()
	refreshCurrentBlock("Refreshed")

	fmt.Printf("\nCongrats, your first Block in your Blockchain is:\n\n")
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END: GenesisBlockchain - Creates the Blockchain (Only run once)"
	log.Println("BLOCKCHAIN I/F:     " + s)

}

// GetBlockchain - Gets the Blockchain
func GetBlockchain() BlockchainSlice {

	s := "START: GetBlockchain - Gets the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	theBlockchain := getBlockchain()

	s = "END: GetBlockchain - Gets the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	return theBlockchain

}

// GetBlock - Get a Block (via Index number) from the Blockchain
func GetBlock(id string) BlockStruct {

	s := "START: GetBlock - Get a Block (via Index number) from the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	theBlock := getBlock(id)

	// RETURN NOT FOUND
	s = "END GetBlock - Get a Block (via Index number) from the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	return theBlock

}

// LoadBlockchain - Loads the Blockchain, LockedBlock and CurrentBlock from a Network Node
func LoadBlockchain(networkIP string, networkTCPPort string) error {

	s := "START: LoadBlockchain - Loads the Blockchain, LockedBlock and CurrentBlock from a Network Node"
	log.Println("BLOCKCHAIN I/F:     " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("BLOCKCHAIN I/F:     " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("BLOCKCHAIN I/F:     " + s)
		return errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "SENDBLOCKCHAIN\n")

	// GET THE BLOCKCHAIN
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("BLOCKCHAIN I/F:     " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("BLOCKCHAIN I/F:     " + s)
		return errors.New(s)
	}

	// GET THE LockedBlock
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("BLOCKCHAIN I/F:     " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("BLOCKCHAIN I/F:     " + s)
		return errors.New(s)
	}

	// GET THE CurrentBlock
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("BLOCKCHAIN I/F:     " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("BLOCKCHAIN I/F:     " + s)
		return errors.New(s)
	}

	fmt.Println(message)
	// LOAD UP THE BLOCKCHAIN FROM THE STRING
	// CHANGE SEND TO GUTS????????????????????????????????????
	json.Unmarshal([]byte(message), &Blockchain)

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END: LoadBlockchain - Loads the Blockchain, LockedBlock and CurrentBlock from a Network Node"
	log.Println("BLOCKCHAIN I/F:     " + s)

	return nil

}

// AddBlockToChain - Add a Block to the Blockchain
// ?????????????????????????? UPDATE
//func AddBlockToChain(firstTransaction string) BlockStruct {

//	s := "Started to add block to blockchain"
//log.Println("BLOCKCHAIN I/F:     " + s)

//var blankBlock BlockStruct

//currentBlock := Blockchain[len(Blockchain)-1]

// JUST TO MAKE SURE CAN'T MAKE A NEW BLOCK AT THE SAME TIME
//mutex.Lock()
//newBlock := appendLockedBlock(currentBlock, firstTransaction)
//mutex.Unlock()

// CHECK IF NEWBLOCK IS VALID
//if isBlockValid(newBlock, currentBlock) {
//	log.Println("BLOCKCHAIN I/F:     Block is valid")
//	newBlockchain := append(Blockchain, newBlock)
// REPLACE WITH LONGER ONE
//	replaceChain(newBlockchain)
//	return newBlock
//}

//s = "Block is NOT valid"
//	log.Println("BLOCKCHAIN I/F:     " + s)
//	return blankBlock

//}

// AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock
func AddTransactionToCurrentBlock(transaction string) BlockStruct {

	s := "START: AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock"
	log.Println("BLOCKCHAIN I/F:     " + s)

	// JUST TO MAKE SURE CAN'T UPDATE A BLOCK AT THE SAME TIME
	// ??????????????????????????? FIX KEEP CURRENT BLOCK IN GUTS
	addTransactionCurrentBlock(transaction)

	s = "END: AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock"
	log.Println("BLOCKCHAIN I/F:     " + s)

	// FIX THIS??????????????????????
	return CurrentBlock

}
