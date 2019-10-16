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

// getBlockchain - Get the Blockchain
func getBlockchain() BlockchainSlice {

	s := "START: getBlockchain - Get the Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	s = "END: getBlockchain - Get the Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	return Blockchain
}

// getBlock - Get a Block the Blockchain
func getBlock(id string) BlockStruct {

	s := "START: getBlock - Get a Block the Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	var item BlockStruct

	// SEARCH DATA FOR blockID
	for _, item := range Blockchain {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END: getBlock - Get a Block the Blockchain"
			log.Println("BLOCKCHAIN GUTS:    " + s)
			return item
		}
	}

	s = "END: getBlock - Get a Block the Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	return item
}

// loadBlockchain - Loads Blockchain
func loadBlockchain(message string) {

	s := "START: loadBlockchain - Loads Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	// LOAD
	json.Unmarshal([]byte(message), &Blockchain)

	s = "END: loadBlockchain - Loads Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

}

// loadCurrentBlock - Loads CurrentBlock
func loadCurrentBlock(message string) {

	s := "START: loadCurrentBlock - Loads CurrentBlock"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	// LOAD
	json.Unmarshal([]byte(message), &CurrentBlock)

	s = "END: loadCurrentBlock - Loads CurrentBlock"
	log.Println("BLOCKCHAIN GUTS:    " + s)

}

// calculateBlockHash - SHA256 hasing
func calculateBlockHash(block BlockStruct) string {

	s := "START: calculateBlockHash - SHA256 hasing"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	record := strconv.Itoa(block.Index) + block.Timestamp + strings.Join(block.Data, " ") + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	s = "Calculated Block Hash"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	s = "END: calculateBlockHash - SHA256 hasing"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	return hex.EncodeToString(hashed)

}

// isBlockValid - Check that the newBlock is valid
func isBlockValid(checkBlock, oldBlock BlockStruct) bool {

	s := "START: isBlockValid - Check that the newBlock is valid"
	log.Println("BLOCKCHAIN GUTS:    " + s)

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

	s = "END: isBlockValid - Check that the newBlock is valid"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	return true

}

// refreshCurrentBlock - Refresh the CurrentBlock
func refreshCurrentBlock(transaction string) {

	s := "START: refreshCurrentBlock - Refresh the CurrentBlock"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	t := time.Now()

	CurrentBlock.Index = 0
	CurrentBlock.Timestamp = t.String()
	CurrentBlock.Data = append(CurrentBlock.Data, transaction)
	CurrentBlock.PrevHash = CurrentBlock.Hash
	CurrentBlock.Difficulty = CurrentBlock.Difficulty
	CurrentBlock.Hash = ""
	CurrentBlock.Nonce = ""

	s = "END: refreshCurrentBlock - Refresh the CurrentBlock"
	log.Println("BLOCKCHAIN GUTS:    " + s)

}

// addTransactionToCurrentBlock - Add Transaction to CurrentBlock
func addTransactionToCurrentBlock(transaction string) BlockStruct {

	s := "START: addTransactionToCurrentBlock - Add Transaction to CurrentBlock"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	CurrentBlock.Data = append(CurrentBlock.Data, transaction)

	s = "END: addTransactionToCurrentBlock - Add Transaction to CurrentBlock"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	return CurrentBlock

}

// lockCurrentBlock - Move CurrentBlock to LockedBlock (ResetCurrentBlock)
func lockCurrentBlock(difficulty int) {

	s := "START: lockCurrentBlock - Move CurrentBlock to LockedBlock (ResetCurrentBlock)"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	CurrentBlock.Hash = calculateBlockHash(CurrentBlock)
	CurrentBlock.Difficulty = difficulty

	LockedBlock = CurrentBlock

	s = "END: lockCurrentBlock - Move CurrentBlock to LockedBlock (ResetCurrentBlock)"
	log.Println("BLOCKCHAIN GUTS:    " + s)

}

// appendLockedBlock - Append LockedBlock to the Blockchain
func appendLockedBlock() BlockStruct {

	s := "START: appendLockedBlock - Append LockedBlock to the Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	//newBlock.Index = currentBlock.Index + 1
	//newBlock.Timestamp = t.String()
	//newBlock.Data = append(newBlock.Data, data)
	//newBlock.PrevHash = currentBlock.Hash
	//newBlock.Difficulty = currentBlock.Difficulty
	//newBlock.Hash = calculateBlockHash(newBlock)
	//newBlock.Nonce = ""

	mutex.Lock()
	Blockchain = append(Blockchain, LockedBlock)
	mutex.Unlock()

	s = "END: appendLockedBlock - Append LockedBlock to the Blockchain"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	return LockedBlock

}

// replaceChain - Replace a chain with a Longer one
func replaceChain(newBlock BlockchainSlice) {

	s := "START: replaceChain - Replace a chain with a Longer one"
	log.Println("BLOCKCHAIN GUTS:    " + s)

	if len(newBlock) > len(Blockchain) {
		s = "New Block added to chain"
		log.Println("BLOCKCHAIN GUTS:    " + s)
		Blockchain = newBlock
	} else {
		s = "New Block NOT added to chain"
		log.Println("BLOCKCHAIN GUTS:    " + s)
	}

	s = "END: replaceChain - Replace a chain with a Longer one"
	log.Println("BLOCKCHAIN GUTS:    " + s)

}
