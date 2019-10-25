package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// proofOfWorkStruct - Used to get POW
// The targetHash works with difficulty to get number of leading 0s you want
type proofOfWorkStruct struct {
	block      *blockStruct
	targetHash *big.Int
}

// blockStruct is your block
type blockStruct struct {
	Index        int                 `json:"index"`
	Timestamp    string              `json:"timestamp"`
	Transactions []transactionStruct `json:"transactions"`
	Hash         string              `json:"hash"`
	PrevHash     string              `json:"prevhash"`
	Difficulty   int                 `json:"difficulty"`
	Nonce        string              `json:"nonce"`
}

var lockedBlock = blockStruct{}

// transactionStruct is your transaction - To be place in block
type transactionStruct struct {
	ID      int64            `json:"ID"`
	Inputs  []txInputStruct  `json:"inputs"`
	Outputs []txOutputStruct `json:"outputs"`
}

// txInput
type txInputStruct struct {
	TXID          int64  `json:"txID"`
	ReferenceTXID int64  `json:"referenceTXID"`
	Signature     string `json:"signature"`
}

// txOutput - This is where the money is stored
type txOutputStruct struct {
	Address string `json:"jeffCoinAddress"`
	Value   int64  `json:"value"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// MINING ************************************************************************************************************

func mineBlock() {

	// GET THE BLOCK TO MINE - lockedBlock
	// Already have lockedBlock

	// LOAD POW STRUCT (With Block and targetHash)
	targetHash := big.NewInt(1)
	difficulty := lockedBlock.Difficulty
	targetHash.Lsh(targetHash, uint(256-difficulty))
	fmt.Printf("The target hash will have %d leading 0s\n", difficulty)

	// powStruct := &proofOfWorkStruct{&lockedBlock, targetHash}

	// MINE - FIND THE NONCE
	foundNonce, hashHex := FindTheNonce(&lockedBlock, targetHash)

	fmt.Printf("The nonce is %d\n", foundNonce)
	fmt.Printf("The hash is %v   %v\n", hashHex, len(hashHex))

	fmt.Printf("The binary is\n")
	// JUST THE BINARY OF THE BEGINING
	//hashByte, _ := hex.DecodeString(hashHex)
	for _, v := range hashHex {
		ui, _ := strconv.ParseUint(string(v), 16, 64)
		format := fmt.Sprintf("%%0%db", 4)
		bin := fmt.Sprintf(format, ui)
		//fmt.Printf("%d, %c, %s\n", i, v, bin)
		fmt.Printf("%s", bin)
	}
	fmt.Println("")

	// VALIDATE FOR FUN
	data := prepareData(&lockedBlock, targetHash, foundNonce)
	hashBytew := sha256.Sum256([]byte(data))

	fmt.Printf("The nonce is %v\n", foundNonce)
	fmt.Printf("The hash is %v\n", hex.EncodeToString(hashBytew[:]))

}

// FindTheNonce - Find the nonce so you have the correct number of leading 0s (difficulty)
func FindTheNonce(block *blockStruct, targetHash *big.Int) (int, string) {

	var hashInt big.Int
	var hashByte [32]byte
	nonce := 0

	// START AT nonce 0 and increment until you find the correct number of leading 0s
	for nonce < math.MaxInt64 {

		// GET DATA WITH INSERTED NONCE & HASH
		data := prepareData(block, targetHash, nonce)
		hashByte = sha256.Sum256([]byte(data))

		// fmt.Printf("\n The nonce is: %v", nonce)
		// fmt.Printf("\n The hash is: %x", hash)

		// CHECK IF HASH HAS targetHash Leading 0 bits
		hashInt.SetBytes(hashByte[:])
		if hashInt.Cmp(targetHash) == -1 {
			break
		} else {
			nonce++
		}
	}

	hashHex := hex.EncodeToString(hashByte[:])

	return nonce, hashHex

}

func prepareData(block *blockStruct, targetHash *big.Int, nonce int) string {

	// GET ALL THE TRANSACTIONS
	transactionBytes := []byte{}
	for _, transaction := range block.Transactions {
		transactionBytes, _ = json.Marshal(transaction)
	}

	// DATA YOU WANT TO HASH
	data := strconv.Itoa(block.Index) + block.Timestamp + string(transactionBytes) + block.PrevHash + string(block.Difficulty) + strconv.Itoa(nonce)

	return data
}

func main() {

	t := time.Now()

	JeffCoinAddress := "99dgh9999"

	firstTransaction := `
    {
        "ID": 0,
        "inputs": [
            {
                "txID": 0,
                "referenceTXID": -1,
                "signature": ""
            }
        ],
        "outputs": [
            {
                "jeffCoinAddress": "` + JeffCoinAddress + `",
                "value": 1000
            }
        ]
    }`

	// Place data in struct
	transactionByte := []byte(firstTransaction)
	var theTransactionStruct transactionStruct
	err := json.Unmarshal(transactionByte, &theTransactionStruct)
	checkErr(err)

	// Place data in slice
	var transactionSlice []transactionStruct
	transactionSlice = append(transactionSlice, theTransactionStruct)

	lockedBlock.Index = 0
	lockedBlock.Timestamp = t.String()
	lockedBlock.Transactions = transactionSlice
	lockedBlock.PrevHash = lockedBlock.Hash
	lockedBlock.Difficulty = 20
	lockedBlock.Hash = ""
	lockedBlock.Nonce = ""

	// js, _ := json.MarshalIndent(lockedBlock, "", "    ")
	// s := string(js)
	// fmt.Printf("The lockBlock is:\n%v", s)

	mineBlock()

}
