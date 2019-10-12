// jeffCoin handlers.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"strings"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	log "github.com/sirupsen/logrus"
)

func handleAddBlock(rw *bufio.ReadWriter) {

	s := "Please enter the first Transaction for the new block"
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION
	transaction, err := rw.ReadString('\n')
	checkErr(err)
	transaction = strings.Trim(transaction, "\n ")
	s = "Received TRANSACTION: " + transaction
	returnMessage(s, rw)

	// MAKE A NEW BLOCK
	// ADD NEW BLOCK TO CHAIN
	s = "Sending request to add block to the Blockchain"
	returnMessage(s, rw)
	newBlock := blockchain.AddBlockToChain(transaction)
	js, _ := json.MarshalIndent(newBlock, "", "    ")
	s = "Added block to Blockchain:\n" + string(js)
	returnMessage(s, rw)

}

func handleAddTransaction(rw *bufio.ReadWriter) {

	s := "Please enter the Transaction for the latest block"
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION
	transaction, err := rw.ReadString('\n')
	checkErr(err)
	transaction = strings.Trim(transaction, "\n ")
	s = "Received TRANSACTION: " + transaction
	returnMessage(s, rw)

	// ADD TRANSACTION TO BLOCK
	s = "Sending request to add block to the Blockchain"
	returnMessage(s, rw)
	updatedBlock := blockchain.AddTransactionToBlock(transaction)
	js, _ := json.MarshalIndent(updatedBlock, "", "    ")
	s = "Added Transaction to Block:\n" + string(js)
	returnMessage(s, rw)

}

func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "Going to Send entire Blockchain to another Node"
	log.Println("ROUTINGNODE:    " + s)

	// SEND ENTIRE BLOCKCHAIN
	sendBlockchain := blockchain.GetBlockchain()
	js, _ := json.Marshal(sendBlockchain)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)

	s = "Sent Entire Blockchain to another node"
	log.Println("ROUTINGNODE:    " + s)
}
