// jeffCoin guts.go

package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var mutex = &sync.Mutex{}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// BLOCKCHAIN ************************************************************************************************************

// loadBlockchain - Loads the entire blockchain
func loadBlockchain(message string) {

	s := "START: loadBlockchain - Loads the entire blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	// LOAD
	json.Unmarshal([]byte(message), &blockchain)

	s = "END:   loadBlockchain - Loads the entire blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// getBlockchain - Gets the blockchain
func getBlockchain() blockchainSlice {

	s := "START: getBlockchain - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	s = "END:   getBlockchain - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return blockchain

}

// replaceBlockchain - Replaces chain with the longer one
func replaceBlockchain(newBlock blockchainSlice) {

	s := "START: replaceChain - Replaces chain with the longer one"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	if len(newBlock) > len(blockchain) {
		s = "New block added to chain"
		log.Info("BLOCKCHAIN:  GUTS           " + s)
		blockchain = newBlock
	} else {
		s = "New Block NOT added to chain"
		log.Info("BLOCKCHAIN:  GUTS           " + s)
	}

	s = "END:   replaceChain - Replaces chain with the longer one"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// BLOCK *****************************************************************************************************************

// getBlock - Gets a block in the blockchain
func getBlock(id string) blockStruct {

	s := "START: getBlock - Gets a block in the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	var item blockStruct

	// SEARCH DATA FOR blockID
	for _, item := range blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END:   getBlock - Gets a block in the blockchain"
			log.Trace("BLOCKCHAIN:  GUTS   " + s)
			return item
		}
	}

	s = "END:   getBlock - Gets a block in the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return item
}

// calculateBlockHash - Calculates SHA256 hash on a block
func calculateBlockHash(block blockStruct) string {

	s := "START: calculateBlockHash - Calculates SHA256 hash on a block"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	transactionBytes := []byte{}

	for _, item := range block.Transactions {
		jsonBytes, _ := json.Marshal(item)
		fmt.Printf("ffffffffffffffffffffffffffffffffffffffffff%v\n\n", string(jsonBytes))
	}

	record := strconv.Itoa(block.Index) + block.Timestamp + string(transactionBytes) + block.PrevHash + block.Difficulty + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	s = "Calculated block Hash"
	log.Info("BLOCKCHAIN:  GUTS          " + s)

	s = "END:   calculateBlockHash - Calculates SHA256 hash on a block"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return hex.EncodeToString(hashed)

}

// isBlockValid - Checks if block is valid
func isBlockValid(checkBlock, oldBlock blockStruct) bool {

	s := "START: isBlockValid - Checks if block is valid"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	// Check index
	if oldBlock.Index+1 != checkBlock.Index {
		return false
	}

	// Compare the hash matches
	if oldBlock.Hash != checkBlock.PrevHash {
		return false
	}

	// Recalculate Hash to check
	if calculateBlockHash(checkBlock) != checkBlock.Hash {
		return false
	}

	s = "END:   isBlockValid - Checks if block is valid"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return true

}

// LOCKED BLOCK **********************************************************************************************************

// getLockedBlock - Gets the lockedBlock
func getLockedBlock() blockStruct {

	s := "START: getLockedBlock - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	s = "END:   getLockedBlock - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return lockedBlock
}

// appendLockedBlock - Appends the lockedBlock to the blockchain
func appendLockedBlock() blockStruct {

	s := "START: appendLockedBlock - Appends the lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	mutex.Lock()
	blockchain = append(blockchain, lockedBlock)
	mutex.Unlock()

	s = "END:   appendLockedBlock - Appends the lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return lockedBlock

}

// CURRENT BLOCK *********************************************************************************************************

// loadCurrentBlock - Loads the currentBlock
func loadCurrentBlock(message string) {

	s := "START: loadCurrentBlock - Loads the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	// LOAD
	json.Unmarshal([]byte(message), &currentBlock)

	s = "END:   loadCurrentBlock - Loads the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// resetCurrentBlock - Resets the currentBlock
func resetCurrentBlock(transaction string) {

	s := "START: resetCurrentBlock - Resets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	t := time.Now()

	// DO STUFF WITH STRING????????????????????????????????????????????????????
	transaction1 := []transactionStruct{
		{
			ID: "01",
			Inputs: []txInput{
				{
					TXID:          "tbd",
					ReferenceTXID: "33",
					Signature:     "tbd",
				},
			},
			Outputs: []txOutput{
				{
					JeffCoinAddress: "tbd",
					Value:           2,
				},
			},
		},
	}

	currentBlock.Index = 0
	currentBlock.Timestamp = t.String()
	currentBlock.Transactions = transaction1
	currentBlock.PrevHash = currentBlock.Hash
	currentBlock.Difficulty = currentBlock.Difficulty
	currentBlock.Hash = ""
	currentBlock.Nonce = ""

	s = "END:   resetCurrentBlock - Resets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// getCurrentBlock - Gets the currentBlock
func getCurrentBlock() blockStruct {

	s := "START: getCurrentBlock - Gets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	s = "END:   getCurrentBlock - Gets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return currentBlock
}

// addTransactionToCurrentBlock - Adds a transaction to the currentBlock
func addTransactionToCurrentBlock(transaction string) blockStruct {

	s := "START: addTransactionToCurrentBlock - Adds a transaction to the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	// DO STUFF WITH STRING????????????????????????????????????????????????????
	transaction1 := transactionStruct{}

	currentBlock.Transactions = append(currentBlock.Transactions, transaction1)

	s = "END:   addTransactionToCurrentBlock - Adds a transaction to the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return currentBlock

}

// lockCurrentBlock - Moves the currentBlock to the lockedBlock and resets the currentBlock
func lockCurrentBlock(difficulty int) {

	s := "START: lockCurrentBlock - Moves the currentBlock to the lockedBlock and resets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	currentBlock.Hash = calculateBlockHash(currentBlock)
	currentBlock.Difficulty = difficulty

	lockedBlock = currentBlock

	s = "END:   lockCurrentBlock -  Moves the currentBlock to the lockedBlock and resets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}
