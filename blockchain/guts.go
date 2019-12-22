// jeffCoin 1. BLOCKCHAIN guts.go

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

// getBlockchain - Gets the blockchain
func getBlockchain() blockchainSlice {

	s := "START  getBlockchain() - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	s = "END    getBlockchain() - Gets the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return blockchain

}

// loadBlockchain - Loads the entire blockchain
func loadBlockchain(message string) {

	s := "START  loadBlockchain() - Loads the entire blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	// LOAD
	json.Unmarshal([]byte(message), &blockchain)

	s = "END    loadBlockchain() - Loads the entire blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

}

// replaceBlockchain - Replaces blockchain with the longer one
func replaceBlockchain(newBlock blockchainSlice) {

	s := "START  replaceChain() - Replaces blockchain with the longer one"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	if len(newBlock) > len(blockchain) {
		s = "New block added to chain"
		log.Info("BLOCKCHAIN:  GUTS             " + s)
		blockchain = newBlock
	} else {
		s = "New Block NOT added to chain"
		log.Info("BLOCKCHAIN:  GUTS             " + s)
	}

	s = "END    replaceChain() - Replaces blockchain with the longer one"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

}

// BLOCK *****************************************************************************************************************

// getBlock - Gets a block in the blockchain
func getBlock(id string) blockStruct {

	s := "START  getBlock() - Gets a block in the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	var blockItem blockStruct

	// SEARCH DATA FOR blockID
	for _, blockItem := range blockchain {
		if strconv.FormatInt(blockItem.BlockID, 10) == id {
			// RETURN BLOCK
			s = "END    getBlock() - Gets a block in the blockchain"
			log.Trace("BLOCKCHAIN:  GUTS     " + s)
			return blockItem
		}
	}

	s = "END    getBlock() - Gets a block in the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return blockItem
}

// calculateBlockHash - Calculates SHA256 hash on a block
func calculateBlockHash(block blockStruct) string {

	s := "START  calculateBlockHash() - Calculates SHA256 hash on a block"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	// GET ALL THE TRANSACTIONS
	transactionBytes := []byte{}
	for _, transaction := range block.Transactions {
		transactionBytes, _ = json.Marshal(transaction)
	}

	hashMe := strconv.FormatInt(block.BlockID, 10) + block.Timestamp + string(transactionBytes) + block.PrevHash + string(block.Difficulty) + block.Nonce
	hashedByte := sha256.Sum256([]byte(hashMe))
	hashed := hex.EncodeToString(hashedByte[:])

	s = "Calculated block Hash: " + hashed
	log.Info("BLOCKCHAIN:  GUTS            " + s)

	s = "END    calculateBlockHash() - Calculates SHA256 hash on a block"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return hashed

}

// isBlockValid - Checks if block is valid
func isBlockValid(checkBlock, oldBlock blockStruct) bool {

	s := "START  isBlockValid() - Checks if block is valid"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	// Check index
	if oldBlock.BlockID+1 != checkBlock.BlockID {
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

	s = "END    isBlockValid() - Checks if block is valid"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return true

}

// LOCKED BLOCK **********************************************************************************************************

// getLockedBlock - Gets the lockedBlock
func getLockedBlock() blockStruct {

	s := "START  getLockedBlock() - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	s = "END    getLockedBlock() - Gets the lockedBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return lockedBlock
}

// appendLockedBlock - Appends the lockedBlock to the blockchain
func appendLockedBlock() blockStruct {

	s := "START  appendLockedBlock() - Appends the lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	mutex.Lock()
	blockchain = append(blockchain, lockedBlock)
	mutex.Unlock()

	s = "END    appendLockedBlock() - Appends the lockedBlock to the blockchain"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return lockedBlock

}

// CURRENT BLOCK *********************************************************************************************************

// getPendingBlock - Gets the pendingBlock
func getPendingBlock() blockStruct {

	s := "START  getPendingBlock() - Gets the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	s = "END    getPendingBlock() - Gets the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return pendingBlock
}

// loadPendingBlock - Loads the pendingBlock
func loadPendingBlock(message string) {

	s := "START  loadPendingBlock() - Loads the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	// LOAD
	json.Unmarshal([]byte(message), &pendingBlock)

	s = "END    loadPendingBlock() - Loads the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

}

// resetPendingBlock - Resets the pendingBlock
func resetPendingBlock(transaction string) {

	s := "START  resetPendingBlock() - Resets the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	t := time.Now()

	// Place data in struct
	transactionByte := []byte(transaction)
	var theTransactionStruct transactionStruct
	err := json.Unmarshal(transactionByte, &theTransactionStruct)
	checkErr(err)

	// Place data in slice
	var transactionSlice []transactionStruct
	transactionSlice = append(transactionSlice, theTransactionStruct)

	pendingBlock.BlockID = 0
	pendingBlock.Timestamp = t.String()
	pendingBlock.Transactions = transactionSlice
	pendingBlock.PrevHash = pendingBlock.Hash
	pendingBlock.Difficulty = pendingBlock.Difficulty
	pendingBlock.Hash = ""
	pendingBlock.Nonce = ""

	s = "END    resetPendingBlock() - Resets the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

}

// addTransactionToPendingBlock - Adds a transaction to the pendingBlock
func addTransactionToPendingBlock(transaction string) blockStruct {

	s := "START  addTransactionToPendingBlock() - Adds a transaction to the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	// DO STUFF WITH STRING????????????????????????????????????????????????????
	transaction1 := transactionStruct{}

	pendingBlock.Transactions = append(pendingBlock.Transactions, transaction1)

	s = "END    addTransactionToPendingBlock() - Adds a transaction to the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return pendingBlock

}

// lockPendingBlock - Moves the pendingBlock to the lockedBlock and resets the pendingBlock
func lockPendingBlock(difficulty int) {

	s := "START  lockPendingBlock() - Moves the pendingBlock to the lockedBlock and resets the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	pendingBlock.Hash = calculateBlockHash(pendingBlock)
	pendingBlock.Difficulty = difficulty

	lockedBlock = pendingBlock

	s = "END    lockPendingBlock() -  Moves the pendingBlock to the lockedBlock and resets the pendingBlock"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

}

// JEFFCOINS *************************************************************************************************************

// getAddressBalance - Gets the jeffCoin Address balance
func getAddressBalance(jeffCoinAddress string) string {

	s := "START  getAddressBalance() - Gets the jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	// WORK BACKWARDS TO FIND THE TRANSACTION
	// GET BLOCK
	for _, blockItem := range blockchain {
		// GET TRANSACTION
		for _, transactionItem := range blockItem.Transactions {
			fmt.Println(transactionItem)
		}
	}

	balance := "333333"

	if jeffCoinAddress == "1234" {
		balance = "1111111"
	}

	s = "END    getAddressBalance() - Gets the jeffCoin Address balance"
	log.Trace("BLOCKCHAIN:  GUTS     " + s)

	return balance

}
