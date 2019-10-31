// jeffCoin handlers.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"strings"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	log "github.com/sirupsen/logrus"
)

// FROM BLOCKCHAIN I/F **********************************************************************************************

// handleSendBlockchain (SBC) - Sends the blockchain & currentBlock to another node
func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "START: handleSendBlockchain (SBC) - Sends the blockchain & currentBlock to another node"
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

	s = "END:   handleSendBlockchain (SBC) - Sends the blockchain & currentBlock to another node"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// FROM ROUTINGNODE I/F **********************************************************************************************

// handleBroadcastAddNewNode (BANN) - Adds a node to the nodeList
func handleBroadcastAddNewNode(rw *bufio.ReadWriter) {

	s := "START: handleBroadcastAddNewNode (BANN) - Adds a node to the nodeList"
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

	s = "END:   handleBroadcastAddNewNode (BANN) - Adds a node to the nodeList"
	log.Trace("ROUTINGNODE: RCV    " + s)

}

// handleSendNodeList (SNL) - Sends the nodeList to another node
func handleSendNodeList(rw *bufio.ReadWriter) {

	s := "START: handleSendNodeList (SNL) - Sends the nodeList to another node"
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
	log.Trace("ROUTINGNODE: RCV           " + s)

	s = "END:   handleSendNodeList (SNL) - Sends the nodeList to another node"
	log.Trace("ROUTINGNODE: RCV    " + s)

}

// handleBroadcastVerifiedBlock (BVB) - A node verified the next block, get block and verify
func handleBroadcastVerifiedBlock(rw *bufio.ReadWriter) {

	s := "START: handleBroadcastVerifiedBlock (BVB) - A node verified the next block, get block and verify"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "END:  handleBroadcastVerifiedBlock (BVB) - A node verified the next block, get block and verify"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleBroadcastConsensus (BC) - 51% Consensus reached, get block to add to blockchain
func handleBroadcastConsensus(rw *bufio.ReadWriter) {

	s := "START: handleBroadcastConsensus (BC) - 51% Consensus reached, get block to add to blockchain"
	log.Trace("ROUTINGNODE: RCV    " + s)

	s = "END:  handleBroadcastConsensus (BC) - 51% Consensus reached, get block to add to blockchain"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleBroadcastTransactionRequest (BTR) - Request from node to transfer coins to a jeffCoin address
func handleBroadcastTransactionRequest(rw *bufio.ReadWriter) {

	s := "START: handleBroadcastTransactionRequest (BTR) - Request from node to transfer coins to a jeffCoin address"
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

	s = "END:  handleBroadcastTransactionRequest (BTR) - Request from node to transfer coins to a jeffCoin address"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// FROM WALLET I/F **********************************************************************************************

// handleSendAddressBalance (SAB) - Sends the coin balance for a jeffCoin Address
func handleSendAddressBalance(rw *bufio.ReadWriter) {

	s := "START: handleSendAddressBalance (SAB) - Sends the coin balance for a jeffCoin Address"
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

	s = "END:   handleSendAddressBalance (SAB) - Sends the coin balance for a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)
}

// handleTransactionRequest (TR) - Request from Wallet to Transfer Coins to a jeffCoin Address
func handleTransactionRequest(rw *bufio.ReadWriter) {

	s := "START: handleTransactionRequest (TR) - Request from Wallet to Transfer Coins to a jeffCoin Address"
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

	s = "END:   handleTransactionRequest (TR) - Request from Wallet to Transfer Coins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: RCV    " + s)
}
