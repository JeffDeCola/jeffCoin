// jeffCoin guts.go

package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

// getBlockchain - Get the blockchain
func getBlockchain() blockchainSlice {

	s := "START: getBlockchain - Get the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	s = "END:   getBlockchain - Get the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return blockchain
}

// getBlock - Get a block the blockchain
func getBlock(id string) blockStruct {

	s := "START: getBlock - Get a block the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	var item blockStruct

	// SEARCH DATA FOR blockID
	for _, item := range blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END:   getBlock - Get a block the blockchain"
			log.Trace("BLOCKCHAIN:  GUTS   " + s)
			return item
		}
	}

	s = "END:   getBlock - Get a block the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return item
}

// getLockedBlock - Get the lockedBlock
func getLockedBlock() blockStruct {

	s := "START: getLockedBlock - Get the lockedBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	s = "END:   getLockedBlock - Get the lockedBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return lockedBlock
}

// getCurrentBlock - Get the currentBlock
func getCurrentBlock() blockStruct {

	s := "START: getCurrentBlock - Get the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	s = "END:   getCurrentBlock - Get the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return currentBlock
}

// loadBlockchain - Loads blockchain
func loadBlockchain(message string) {

	s := "START: loadBlockchain - Loads blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	// LOAD
	json.Unmarshal([]byte(message), &blockchain)

	s = "END:   loadBlockchain - Loads blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// loadCurrentBlock - Loads currentBlock
func loadCurrentBlock(message string) {

	s := "START: loadCurrentBlock - Loads currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	// LOAD
	json.Unmarshal([]byte(message), &currentBlock)

	s = "END:   loadCurrentBlock - Loads currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// addTransactionToCurrentBlock - Add Transaction to CurrentBlock
func addTransactionToCurrentBlock(transaction string) blockStruct {

	s := "START: addTransactionToCurrentBlock - Add Transaction to currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	currentBlock.Data = append(currentBlock.Data, transaction)

	s = "END:   addTransactionToCurrentBlock - Add Transaction to currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return currentBlock

}

// calculateBlockHash - SHA256 hasing
func calculateBlockHash(block blockStruct) string {

	s := "START: calculateBlockHash - SHA256 hasing"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	record := strconv.Itoa(block.Index) + block.Timestamp + strings.Join(block.Data, " ") + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	s = "Calculated block Hash"
	log.Info("BLOCKCHAIN:  GUTS          " + s)

	s = "END:   calculateBlockHash - SHA256 hasing"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return hex.EncodeToString(hashed)

}

// isBlockValid - Check that the newBlock is valid
func isBlockValid(checkBlock, oldBlock blockStruct) bool {

	s := "START: isBlockValid - Check that the newBlock is valid"
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

	s = "END:   isBlockValid - Check that the newBlock is valid"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return true

}

// resetCurrentBlock - Resets the currentBlock
func resetCurrentBlock(transaction string) {

	s := "START: resetCurrentBlock - Resets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	t := time.Now()

	currentBlock.Index = 0
	currentBlock.Timestamp = t.String()
	currentBlock.Data = []string{transaction}
	currentBlock.PrevHash = currentBlock.Hash
	currentBlock.Difficulty = currentBlock.Difficulty
	currentBlock.Hash = ""
	currentBlock.Nonce = ""

	s = "END:   resetCurrentBlock - Resets the currentBlock"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// lockCurrentBlock - Move currentBlock to lockedBlock (ResetCurrentBlock)
func lockCurrentBlock(difficulty int) {

	s := "START: lockCurrentBlock - Move currentBlock to lockedBlock (ResetCurrentBlock)"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	currentBlock.Hash = calculateBlockHash(currentBlock)
	currentBlock.Difficulty = difficulty

	lockedBlock = currentBlock

	s = "END:   lockCurrentBlock - Move currentBlock to lockedBlock (ResetCurrentBlock)"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}

// appendLockedBlock - Append lockedBlock to the Blockchain
func appendLockedBlock() blockStruct {

	s := "START: appendLockedBlock - Append lockedBlock to the Blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	//newBlock.Index = currentBlock.Index + 1
	//newBlock.Timestamp = t.String()
	//newBlock.Data = append(newBlock.Data, data)
	//newBlock.PrevHash = currentBlock.Hash
	//newBlock.Difficulty = currentBlock.Difficulty
	//newBlock.Hash = calculateBlockHash(newBlock)
	//newBlock.Nonce = ""

	mutex.Lock()
	blockchain = append(blockchain, lockedBlock)
	mutex.Unlock()

	s = "END:   appendLockedBlock - Append lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	return lockedBlock

}

// replaceChain - Replace a chain with a Longer one
func replaceChain(newBlock blockchainSlice) {

	s := "START: replaceChain - Replace a chain with a Longer one"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

	if len(newBlock) > len(blockchain) {
		s = "New block added to chain"
		log.Info("BLOCKCHAIN:  GUTS           " + s)
		blockchain = newBlock
	} else {
		s = "New Block NOT added to chain"
		log.Info("BLOCKCHAIN:  GUTS           " + s)
	}

	s = "END:   replaceChain - Replace a chain with a Longer one"
	log.Trace("BLOCKCHAIN:  GUTS   " + s)

}
