// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip blockchain-interface.go

package blockchain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var mutex = &sync.Mutex{}

// CreateBlockchain - Create a blockchain
func CreateBlockchain(transaction string, difficulty int) {

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
}

// GetBlockchain - Get the Blockchain
func GetBlockchain() BlockchainSlice {

	return Blockchain

}

// GetBlock - Get a Block from the chain
func GetBlock(id string) BlockStruct {

	log.Println("BLOCKCHAIN I/F: Get a block from blockchain")

	var item BlockStruct

	// SEARCH DATA FOR blockID
	for _, item := range Blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			return item
		}
	}

	// RETURN NOT FOUND
	return item
}

// AddBlockToChain - Add a Block to the chain
func AddBlockToChain(firstTransaction string) BlockStruct {

	log.Println("BLOCKCHAIN I/F: Started to add block to blockchain")

	var blankBlock BlockStruct

	currentBlock := Blockchain[len(Blockchain)-1]

	// JUST TO MAKE SURE CAN'T MAKE A NEW BLOCK AT THE SAME TIME
	mutex.Lock()
	newBlock := addBlock(currentBlock, firstTransaction)
	mutex.Unlock()

	// CHECK IF NEWBLOCK IS VALID
	if isBlockValid(newBlock, currentBlock) {
		log.Println("BLOCKCHAIN I/F: Block is valid")
		newBlockchain := append(Blockchain, newBlock)
		// REPLACE WITH LONGER ONE
		replaceChain(newBlockchain)
		return newBlock
	}

	log.Println("BLOCKCHAIN I/F: Block is NOT valid")
	return blankBlock

}

// AddTransactionToBlock - Add a Transaction to current Block
func AddTransactionToBlock(transaction string) BlockStruct {

	log.Println("BLOCKCHAIN I/F: Started to add transaction to current Block")

	currentBlock := Blockchain[len(Blockchain)-1]

	// JUST TO MAKE SURE CAN'T UPDATE A BLOCK AT THE SAME TIME
	mutex.Lock()
	updatedBlock := addTransaction(currentBlock, transaction)
	mutex.Unlock()

	// REPLACE CURRENT BLOCK WITH UPDATED ONE
	Blockchain[len(Blockchain)-1] = updatedBlock

	log.Println("BLOCKCHAIN I/F: Added transaction")
	return updatedBlock

}
