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
	sendLockedBlock := blockchain.GetLockedBlock()
	js, _ = json.Marshal(sendLockedBlock)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent LockedBlock to another node"
	log.Info("ROUTINGNODE: RCV           " + s)

	// SEND CurrentBlock
	sendCurrentBlock := blockchain.GetCurrentBlock()
	js, _ = json.Marshal(sendCurrentBlock)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent currentBlock to another node"
	log.Info("ROUTINGNODE: RCV           " + s)

	s = "END:   handleSendBlockchain - Sends the blockchain & currentBlock to another node"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// ROUTING NODE **********************************************************************************************

// handleBroadcastAddNewNode - Broadcast adds a node to the nodeList
func handleBroadcastAddNewNode(rw *bufio.ReadWriter) {

	s := "START: handleBroadcastAddNewNode - Broadcast adds a node to the nodeList"
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

	s = "END:   handleBroadcastAddNewNode - Broadcast adds a node to the nodeList"
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

// handleBroadCastVerifiedBlock - Broadcast a Node Found a Hash
func handleBroadCastVerifiedBlock(rw *bufio.ReadWriter) {

	s := "START: handleBroadCastVerifiedBlock - Broadcast a Node Found a Hash"
    log.Trace("ROUTINGNODE: RCV    " + s)
    
    s = "END:  handleBroadCastVerifiedBlock - Broadcast a Node Found a Hash"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleBroadCastConsensus - Broadcast Transaction Request Message (Signed)
func handleBroadCastConsensus(rw *bufio.ReadWriter) {

	s := "START: handleBroadCastConsensus - Broadcast Transaction Request Message (Signed)"
    log.Trace("ROUTINGNODE: RCV    " + s)
    
    s = "END:  handleBroadCastConsensus - Broadcast Transaction Request Message (Signed)"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleBroadCastTransactionRequest - Broadcast Transaction Request Message (Signed)
func handleBroadCastTransactionRequest(rw *bufio.ReadWriter) {

	s := "START: handleBroadCastTransactionRequest - Broadcast Transaction Request Message (Signed)"
    log.Trace("ROUTINGNODE: RCV    " + s)
    
    s = "END:  handleBroadCastTransactionRequest - Broadcast Transaction Request Message (Signed)"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// WALLET **********************************************************************************************

// handleSendAddressBalance - Gets jeffCoin Address balance
func handleSendAddressBalance(rw *bufio.ReadWriter) {

	s := "START: handleSendAddressBalance - Gets jeffCoin Address balance"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "Please enter the jeffCoinAddress you want the balance for"
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	// WAITING FOR JEFFCOINADDRESS
	jeffCoinAddress, err := rw.ReadString('\n')
	checkErr(err)
	jeffCoinAddress = strings.Trim(jeffCoinAddress, "\n ")
	s = "Received jeffCoinAddress: " + jeffCoinAddress
	log.Info("ROUTINGNODE: RCV           " + s)

	// GET ADDRESS BALANCE
	theBalance := blockchain.GetAddressBalance(jeffCoinAddress)
	s = "The balance for address " + jeffCoinAddress + " is " + theBalance
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	s = "END:   handleSendAddressBalance - Gets jeffCoin Address balance"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleTransactionRequest - Request from Wallet to Transfer Coins to a jeffCoin Address
func handleTransactionRequest(rw *bufio.ReadWriter) {

	s := "START: handleTransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "Please enter the transactionRequestMessageSigned"
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION REQUEST
	transactionRequestMessageSigned, err := rw.ReadString('\n')
	checkErr(err)
	transactionRequestMessageSigned = strings.Trim(transactionRequestMessageSigned, "\n ")
	s = "Received TRANSACTION: " + transactionRequestMessageSigned
	log.Info("ROUTINGNODE: RCV           " + s)

	// BROADCAST TRANSACTION REQUEST TO ALL NODES
	// ???????????????????????????????????????????????

	// TRANSACTION REQUEST
	status := blockchain.TransactionRequest(transactionRequestMessageSigned)
	s = "The Status is: " + status
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	s = "END:   handleTransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleBroadcastTransactionRequest - Request from node to Transfer Coins to a jeffCoin Address
// Same as above except not broadcasting to all nodes
func handleBroadcastTransactionRequest(rw *bufio.ReadWriter) {

	s := "START: handleBroadcastTransactionRequest - Request from node to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "Please enter the transactionRequestMessageSigned"
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	// WAITING FOR TRANSACTION REQUEST
	transactionRequestMessageSigned, err := rw.ReadString('\n')
	checkErr(err)
	transactionRequestMessageSigned = strings.Trim(transactionRequestMessageSigned, "\n ")
	s = "Received TRANSACTION: " + transactionRequestMessageSigned
	log.Info("ROUTINGNODE: RCV           " + s)

	// TRANSACTION REQUEST
	status := blockchain.TransactionRequest(transactionRequestMessageSigned)
	s = "The Status is: " + status
	log.Info("ROUTINGNODE: RCV           " + s)
	returnMessage(s, rw)

	s = "END:   handleBroadcastTransactionRequest - Request from node to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)
}
