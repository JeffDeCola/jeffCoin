// jeffCoin guts.go

package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// calculateBlockHash - SHA256 hasing
func calculateBlockHash(block BlockStruct) string {

	s := "START: calculateBlockHash - SHA256 hasing"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	record := strconv.Itoa(block.Index) + block.Timestamp + strings.Join(block.Data, " ") + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	log.Println("GUTS:           Calculated Block Hash")

	s = "END: calculateBlockHash - SHA256 hasing"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	return hex.EncodeToString(hashed)

}

// isBlockValid - Check that the newBlock is valid
func isBlockValid(checkBlock, oldBlock BlockStruct) bool {

	s := "START: isBlockValid - Check that the newBlock is valid"
	log.Println("BLOCKCHAIN GUTS:     " + s)

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
	log.Println("BLOCKCHAIN GUTS:     " + s)

    return true
    
}

// addLockedBlock - Add LockedBlock to the Blockchain
func addLockedBlock(currentBlock BlockStruct, data string) BlockStruct {

	s := "START: addLockedBlock - Add a LockedBlock to the Blockchain"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	var newBlock BlockStruct

	t := time.Now()

	newBlock.Index = currentBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = append(newBlock.Data, data)
	newBlock.PrevHash = currentBlock.Hash
	newBlock.Difficulty = currentBlock.Difficulty
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Nonce = ""

	s := "END: addLockedBlock - Add a LockedBlock to the Blockchain"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	return newBlock

}

// addTransaction - Add Transaction to CurrentBlock
func addTransaction(block BlockStruct, transaction string) BlockStruct {

	s := "START: addTransaction - Add Transaction to a block"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	var updateBlock BlockStruct

	updateBlock.Index = block.Index
	updateBlock.Timestamp = block.Timestamp
	updateBlock.Data = append(block.Data, transaction)
	updateBlock.PrevHash = block.Hash
	updateBlock.Difficulty = block.Difficulty
	updateBlock.Hash = block.Hash
	updateBlock.Nonce = block.Nonce

	s = "END: addTransaction - Add Transaction to a block"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	return updateBlock

}

// replaceChain - Replace a chain with a Longer one
func replaceChain(newBlock BlockchainSlice) {

	s := "START: replaceChain - Replace a chain with a Longer one"
	log.Println("BLOCKCHAIN GUTS:     " + s)

	if len(newBlock) > len(Blockchain) {
		s = "New Block added to chain"
		log.Println("BLOCKCHAIN GUTS:     " + s)
		Blockchain = newBlock
	} else {
		s = "New Block NOT added to chain"
		log.Println("BLOCKCHAIN GUTS:     " + s)
	}

	s = "END: replaceChain - Replace a chain with a Longer one"
	log.Println("BLOCKCHAIN GUTS:     " + s)

}
