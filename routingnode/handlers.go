// jeffCoin handlers.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"strings"
	"time"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	log "github.com/sirupsen/logrus"
)

// handleAddTransaction - Add Transaction to CurrentBlock
func handleAddTransaction(rw *bufio.ReadWriter) {

	s := "START: handleAddTransaction - Add Transaction to CurrentBlock"
	log.Println("ROUTINGNODE:RCV     " + s)

	s = "Please enter the Transaction for the latest block"
	log.Println("ROUTINGNODE:RCV     " + s)
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION
	transaction, err := rw.ReadString('\n')
	checkErr(err)
	transaction = strings.Trim(transaction, "\n ")
	s = "Received TRANSACTION: " + transaction
	log.Println("ROUTINGNODE:RCV     " + s)

	// ADD TRANSACTION TO BLOCK
	s = "Sending request to add block to the Blockchain"
	log.Println("ROUTINGNODE:RCV     " + s)
	currentBlock := blockchain.AddTransactionToCurrentBlock(transaction)
	js, _ := json.MarshalIndent(currentBlock, "", "    ")
	s = "Added Transaction to Block:\n" + string(js)
	log.Println("ROUTINGNODE:RCV     " + s)

	s = "END: handleAddTransaction - Add Transaction to CurrentBlock"
	log.Println("ROUTINGNODE:RCV     " + s)
}

// handleSendBlockchain - Send Blockchain, LockedBlock & CurrentBlock to another Node
func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "START: handleSendBlockchain - Send Blockchain, LockedBlock & CurrentBlock to another Node"
	log.Println("ROUTINGNODE:RCV     " + s)

	// SEND ENTIRE BLOCKCHAIN
	sendBlockchain := blockchain.GetBlockchain()
	js, _ := json.Marshal(sendBlockchain)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent Blockchain to another node"
	log.Println("ROUTINGNODE:RCV     " + s)

	// SEND LockedBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent LockedBlock to another node"
	log.Println("ROUTINGNODE:RCV     " + s)

	// SEND CurrentBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)

	s = "END: handleSendBlockchain - Send Blockchain, LockedBlock & CurrentBlock to another Node"
	log.Println("ROUTINGNODE:RCV     " + s)
}

// handleSendNodeList - Send NodeList to another Node
func handleSendNodeList(rw *bufio.ReadWriter) {

	s := "START: handleSendNodeList - Send NodeList to another Node"
	log.Println("ROUTINGNODE:RCV     " + s)

	// SEND NODELIST
	sendNodeList := GetNodeList()
	js, _ := json.Marshal(sendNodeList)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent Nodelist to another node"
	log.Println("ROUTINGNODE:RCV     " + s)

	s = "END: handleSendNodeList - Send NodeList to another Node"
    log.Println("ROUTINGNODE:RCV     " + s)
    
}

// handleAddNewNode - Add Node to NodeList
func handleAddNewNode(rw *bufio.ReadWriter) {

	s := "START: handleNewNode - Add Node to NodeList"
	log.Println("ROUTINGNODE:RCV     " + s)

	t := time.Now()

	newNode := NodeStruct{
		Index:     0,
		Timestamp: t.String(),
		IP:        "adsf?????",
		Port:      "asdf?????",
	}

	NodeList = append(NodeList, newNode)

	s = "END: handleNewNode - Add Node to NodeList"
    log.Println("ROUTINGNODE:RCV     " + s)
    
}
