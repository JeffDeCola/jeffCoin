// jeffCoin handlers.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"strings"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	log "github.com/sirupsen/logrus"
)

// handleAddTransaction - Add Transaction to CurrentBlock
func handleAddTransaction(rw *bufio.ReadWriter) {

	s := "START: handleAddTransaction - Add Transaction to CurrentBlock"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "Please enter the Transaction for the latest block"
	log.Trace("ROUTINGNODE: RCV    " + s)
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION
	transaction, err := rw.ReadString('\n')
	checkErr(err)
	transaction = strings.Trim(transaction, "\n ")
	s = "Received TRANSACTION: " + transaction
	log.Trace("ROUTINGNODE: RCV    " + s)

	// ADD TRANSACTION TO BLOCK
	s = "Sending request to add block to the Blockchain"
	log.Trace("ROUTINGNODE: RCV    " + s)
	currentBlock := blockchain.AddTransactionToCurrentBlock(transaction)
	js, _ := json.MarshalIndent(currentBlock, "", "    ")
	s = "Added Transaction to Block:\n" + string(js)
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "END:   handleAddTransaction - Add Transaction to CurrentBlock"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleSendBlockchain - Send Blockchain, LockedBlock & CurrentBlock to another Node
func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "START: handleSendBlockchain - Send Blockchain, LockedBlock & CurrentBlock to another Node"
	log.Trace("ROUTINGNODE: RCV    " + s)

	// SEND ENTIRE BLOCKCHAIN
	sendBlockchain := blockchain.GetBlockchain()
	js, _ := json.Marshal(sendBlockchain)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent Blockchain to another node"
	log.Trace("ROUTINGNODE: RCV           " + s)

	// SEND LockedBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent LockedBlock to another node"
	log.Trace("ROUTINGNODE: RCV           " + s)

	// SEND CurrentBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)

	s = "END:   handleSendBlockchain - Send Blockchain, LockedBlock & CurrentBlock to another Node"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleSendNodeList - Send NodeList to another Node
func handleSendNodeList(rw *bufio.ReadWriter) {

	s := "START: handleSendNodeList - Send NodeList to another Node"
	log.Trace("ROUTINGNODE: RCV    " + s)

	// SEND NODELIST
	sendNodeList := GetNodeList()
	js, _ := json.Marshal(sendNodeList)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent Nodelist to another node"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "END:   handleSendNodeList - Send NodeList to another Node"
	log.Trace("ROUTINGNODE: RCV    " + s)

}

// handleAddNewNode - Add Node to NodeList
func handleAddNewNode(rw *bufio.ReadWriter) {

	s := "START: handleAddNewNode - Add Node to NodeList"
	log.Trace("ROUTINGNODE: RCV    " + s)

	// RESPOND - SEND NEW NODE
	s = "Sent The New Node"
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent The New Node"
	log.Trace("ROUTINGNODE: RCV           " + s)

	// RECEIVE THE NEW NODE
	messageNewNode, err := rw.ReadString('\n')
	checkErr(err)
	// TRIM CMD
	messageNewNode = strings.Trim(messageNewNode, "\n ")

	// APPEND
	newNode := AppendNewNode(messageNewNode)

	js, _ := json.MarshalIndent(newNode, "", "    ")
	s = "Appended new Node to the NodeList:\n" + string(js)
	log.Trace("ROUTINGNODE: RCV           " + s)

	// RESPOND - THANK YOU
	s = "Thank you"
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Thank you"
	log.Info("ROUTINGNODE: RCV           " + s)

	s = "END:   handleAddNewNode - Add Node to NodeList"
	log.Trace("ROUTINGNODE: RCV    " + s)

}
