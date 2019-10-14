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

	s := "STARTING - Add Block BlockChain"
	log.Println("ROUTINGNODE:RCV " + s)

	s = "Please enter the first Transaction for the new block"
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

	s = "DONE - Added Block BlockChain"
	log.Println("ROUTINGNODE:RCV " + s)
}

// Add Transaction to Current Block
func handleAddTransaction(rw *bufio.ReadWriter) {

	s := "STARTING - Add Transaction to Current Block"
	log.Println("ROUTINGNODE:RCV " + s)

	s = "Please enter the Transaction for the latest block"
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION
	transaction, err := rw.ReadString('\n')
	checkErr(err)
	transaction = strings.Trim(transaction, "\n ")
	s = "Received TRANSACTION: " + transaction
	log.Println("ROUTINGNODE:RCV " + s)

	// ADD TRANSACTION TO BLOCK
	s = "Sending request to add block to the Blockchain"
	log.Println("ROUTINGNODE:RCV " + s)
	updatedBlock := blockchain.AddTransactionToBlock(transaction)
	js, _ := json.MarshalIndent(updatedBlock, "", "    ")
	s = "Added Transaction to Block:\n" + string(js)
	log.Println("ROUTINGNODE:RCV " + s)

	s = "DONE - Added Transaction to Current Block"
	log.Println("ROUTINGNODE:RCV " + s)
}

// Send Blockchain, LockedBlock & CurrentBlock to another Node
func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "STARTING - Send Blockchain, LockedBlock & CurrentBlock to another node"
	log.Println("ROUTINGNODE:RCV " + s)

	// SEND ENTIRE BLOCKCHAIN
	sendBlockchain := blockchain.GetBlockchain()
	js, _ := json.Marshal(sendBlockchain)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent Blockchain to another node"
	log.Println("ROUTINGNODE:RCV " + s)

	// SEND LockedBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent LockedBlock to another node"
	log.Println("ROUTINGNODE:RCV " + s)

	// SEND CurrentBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent CurrentBlock to another node"
	log.Println("ROUTINGNODE:RCV " + s)

	s = "DONE - Sent Blockchain, LockedBlock & CurrentBlock to another node"
	log.Println("ROUTINGNODE:RCV " + s)
}

// Add Node to Node List
func handleNewNode(rw *bufio.ReadWriter) {

	s := "STARTING - Adding Node to Node List"
	log.Println("ROUTINGNODE:RCV " + s)

    t := time.Now()

	newNode = NodesStruct{
		Index:      0,
        Timestamp:  t.String(),
        IPs:        ,
        Ports:      ,
    }

    Nodes = append(Nodes, newNode)

	s = "DONE - Added Node to Node List"
	log.Println("ROUTINGNODE:RCV " + s)
}
