// jeffCoin guts.go

package miner

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

    blockchain "github.com/JeffDeCola/jeffCoin/blockchain"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// MINING ************************************************************************************************************

func mine() {

    // GET THE LOCKED BLOCK
    theLockedblock := blockchain.GetLockedBlock() 
    
    // LOAD STRUCT
    targetHash := big.NewInt(1)
	targetHash.Lsh(target, uint(256-targetBits))
    pow := &ProofOfWork{thelockedBlock, targetHash}
    

    // MINE - FIND THE NONCE AND HASH
    nonce, hash := pow.Run()

    fmt.Println("The nonce is", nonce)
    fmt.Println("The hash is", hash)
    
}


func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

    return nonce, hash[:]
    
}


