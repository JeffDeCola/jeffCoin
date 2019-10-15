// jeffCoin blockchain-interface.go

package blockchain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	errors "github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

var mutex = &sync.Mutex{}

// GenesisBlockchain - Creates the Blockchain (Only run once)
func GenesisBlockchain(transaction string, difficulty int) {

	s := "START: GenesisBlockchain - Creates the Blockchain (Only run once)"
	log.Println("BLOCKCHAIN I/F:     " + s)

	t := time.Now()
	firstBlock := BlockStruct{}

	firstBlock = BlockStruct{
		Index:      0,
		Timestamp:  t.String(),
		Data:       append(firstBlock.Data, transaction),
		Hash:       calculateBlockHash(firstBlock),
		PrevHash:   "",
		Difficulty: difficulty,
		Nonce:      "",
	}

	fmt.Printf("\nCongrats, your first Block in your blockchain is:\n\n")
	js, _ := json.MarshalIndent(firstBlock, "", "    ")
	fmt.Printf("%v\n", string(js))

	Blockchain = append(Blockchain, firstBlock)

	s = "END: GenesisBlockchain - Creates the Blockchain (Only run once)"
	log.Println("BLOCKCHAIN I/F:     " + s)

}

// GetBlockchain - Gets the Blockchain
func GetBlockchain() BlockchainSlice {

	s := "START: GetBlockchain - Gets the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	s = "END: GetBlockchain - Gets the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	// ?????????????????? GET FROM GUTS
	return Blockchain

}

// GetBlock - Get a Block (via Index number) from the Blockchain
func GetBlock(id string) BlockStruct {

	s := "START: GetBlock - Get a Block (via Index number) from the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	var item BlockStruct

	// SEARCH DATA FOR blockID
	for _, item := range Blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END: GetBlock - Get a Block (via Index number) from the Blockchain"
			log.Println("BLOCKCHAIN I/F:     " + s)
			return item
		}
	}

	// RETURN NOT FOUND
	s = "END (ITEM NOT FOUND): GetBlock - Get a Block (via Index number) from the Blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)
	return item

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
func AddBlockToChain(firstTransaction string) BlockStruct {

	s := "Started to add block to blockchain"
	log.Println("BLOCKCHAIN I/F:     " + s)

	var blankBlock BlockStruct

	currentBlock := Blockchain[len(Blockchain)-1]

	// JUST TO MAKE SURE CAN'T MAKE A NEW BLOCK AT THE SAME TIME
	mutex.Lock()
	newBlock := addBlock(currentBlock, firstTransaction)
	mutex.Unlock()

	// CHECK IF NEWBLOCK IS VALID
	if isBlockValid(newBlock, currentBlock) {
		log.Println("BLOCKCHAIN I/F:     Block is valid")
		newBlockchain := append(Blockchain, newBlock)
		// REPLACE WITH LONGER ONE
		replaceChain(newBlockchain)
		return newBlock
	}

	s = "Block is NOT valid"
	log.Println("BLOCKCHAIN I/F:     " + s)
	return blankBlock

}

// AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock
func AddTransactionToCurrentBlock(transaction string) BlockStruct {

	s := "START: AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock"
	log.Println("BLOCKCHAIN I/F:     " + s)

	// JUST TO MAKE SURE CAN'T UPDATE A BLOCK AT THE SAME TIME
	// ??????????????????????????? FIX KEEP CURRENT BLOCK IN GUTS
	mutex.Lock()
	currentBlock := addTransaction(CurrentBlock, transaction)
	mutex.Unlock()

	s = "END: AddTransactionToCurrentBlock - Add a Transaction to CurrentBlock"
	log.Println("BLOCKCHAIN I/F:     " + s)

	return currentBlock

}
