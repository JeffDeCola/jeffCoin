// jeffCoin handlers.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"strings"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	log "github.com/sirupsen/logrus"
)

// BLOCKCHAIN **********************************************************************************************

// handleSendBlockchain - Sends the blockchain & currentBlock to another node
func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "START: handleSendBlockchain - Sends the blockchain & currentBlock to another node"
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
	log.Info("ROUTINGNODE: RCV           " + s)

	// SEND LockedBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent LockedBlock to another node"
	log.Info("ROUTINGNODE: RCV           " + s)

	// SEND CurrentBlock
	sendBlockchain = blockchain.GetBlockchain()
	js, _ = json.Marshal(sendBlockchain)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)

	s = "END:   handleSendBlockchain - Sends the blockchain & currentBlock to another node"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// ROUTING NODE **********************************************************************************************

// handleAddNewNode - Adds a node to the nodeList
func handleAddNewNode(rw *bufio.ReadWriter) {

	s := "START: handleAddNewNode - Adds a node to the nodeList"
	log.Trace("ROUTINGNODE: RCV    " + s)

	// RESPOND - SEND NEW NODE
	s = "Sent The New Node"
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent The New Node"
	log.Info("ROUTINGNODE: RCV           " + s)

	// RECEIVE THE NEW NODE
	messageNewNode, err := rw.ReadString('\n')
	checkErr(err)
	// TRIM CMD
	messageNewNode = strings.Trim(messageNewNode, "\n ")

	// APPEND
	newNode := AppendNewNode(messageNewNode)

	js, _ := json.MarshalIndent(newNode, "", "    ")
	s = "Appended new Node to the NodeList:\n" + string(js)
	log.Info("ROUTINGNODE: RCV           " + s)

	// RESPOND - THANK YOU
	s = "Thank you"
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Thank you"
	log.Info("ROUTINGNODE: RCV           " + s)

	s = "END:   handleAddNewNode - Adds a node to the nodeList"
	log.Trace("ROUTINGNODE: RCV    " + s)

}

// handleSendNodeList - Sends the nodeList to another node
func handleSendNodeList(rw *bufio.ReadWriter) {

	s := "START: handleSendNodeList - Sends the nodeList to another node"
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

	s = "END:   handleSendNodeList - Sends the nodeList to another node"
	log.Trace("ROUTINGNODE: RCV    " + s)

}

// WALLET **********************************************************************************************

// handleSendAddressBalance - Gets jeffCoin Address balance
func handleSendAddressBalance(rw *bufio.ReadWriter) {

	s := "START: handleSendAddressBalance - Gets jeffCoin Address balance"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "Please enter the jeffCoinAddress you want the balance for"
	log.Info("ROUTINGNODE: RCV    " + s)
	returnMessage(s, rw)

	// WAITING FOR JEFFCOINADDRESS
	jeffCoinAddress, err := rw.ReadString('\n')
	checkErr(err)
	jeffCoinAddress = strings.Trim(jeffCoinAddress, "\n ")
	s = "Received jeffCoinAddress: " + jeffCoinAddress
	log.Info("ROUTINGNODE: RCV    " + s)

	// GET ADDRESS BALANCE
	theBalance := blockchain.GetAddressBalance(jeffCoinAddress)
	s = "The balance for address " + jeffCoinAddress + " is " + theBalance
	log.Info("ROUTINGNODE: RCV    " + s)
	returnMessage(s, rw)

	s = "END:   handleSendAddressBalance - Gets jeffCoin Address balance"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleTransactionRequest - Request to Transfer Coins to a jeffCoin Address
func handleTransactionRequest(rw *bufio.ReadWriter) {

	s := "START: handleTransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "Please enter the Address and value????????????????????????????????"
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION REQUEST
	transactionRequest, err := rw.ReadString('\n')
	checkErr(err)
	transactionRequest = strings.Trim(transactionRequest, "\n ")
	s = "Received TRANSACTION: " + transactionRequest
	log.Info("ROUTINGNODE: RCV           " + s)

	// TRANSACTION REQUEST
	IDONTKNOW := blockchain.TransactionRequest(transactionRequest)
	js, _ := json.MarshalIndent(IDONTKNOW, "", "    ")
	s = "I DONTKNOW ?????" + string(js)
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	s = "END:   handleTransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)
}
