// jeffCoin 3. ROUTINGNODE handlers.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"strings"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	log "github.com/sirupsen/logrus"
)

// FROM BLOCKCHAIN I/F ***************************************************************************************************

// handleSendBlockchain - REQUEST-BLOCKCHAIN (RBC) - Sends the blockchain and currentBlock to another Node
func handleSendBlockchain(rw *bufio.ReadWriter) {

	s := "START  handleSendBlockchain() - REQUEST-BLOCKCHAIN (RBC) - Sends the blockchain and currentBlock to another Node"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	// RESPOND - SEND BLOCKCHAIN
	sendBlockchain := blockchain.GetBlockchain()
	js, _ := json.Marshal(sendBlockchain)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent blockchain to another node"
	log.Info("ROUTINGNODE: HDLR  -sent   " + s)

	// RESPOND - SEND LockedBlock
	sendLockedBlock := blockchain.GetLockedBlock()
	js, _ = json.Marshal(sendLockedBlock)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent LockedBlock to another node"
	log.Info("ROUTINGNODE: HDLR  -sent   " + s)

	// RESPOND - SEND CurrentBlock
	sendCurrentBlock := blockchain.GetCurrentBlock()
	js, _ = json.Marshal(sendCurrentBlock)
	s = string(js)
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent currentBlock to another node"
	log.Info("ROUTINGNODE: HDLR  -sent   " + s)

	s = "END    handleSendBlockchain() - REQUEST-BLOCKCHAIN (RBC) - Sends the blockchain and currentBlock to another Node"
	log.Trace("ROUTINGNODE: HDLR   " + s)
}

// FROM ROUTINGNODE I/F **************************************************************************************************

// handleBroadcastAddNewNode - BROADCAST-ADD-NEW-NODE (BANN) - Adds a Node to the nodeList
func handleBroadcastAddNewNode(rw *bufio.ReadWriter) {

	s := "START  handleBroadcastAddNewNode() - BROADCAST-ADD-NEW-NODE (BANN) - Adds a Node to the nodeList"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	// RESPOND - SEND NEW NODE
	s = "Please send The New Node so I can append to my nodeList"
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Please send The New Node so I can append to my nodeList"
	log.Info("ROUTINGNODE: HDLR  -sent   " + s)

	// RECEIVING NEW NODE
	messageNewNode, err := rw.ReadString('\n')
	checkErr(err)
	messageNewNode = strings.Trim(messageNewNode, "\n ")
	s = "-rcvd   - NEW NODE: " + messageNewNode
	log.Info("ROUTINGNODE: HDLR  " + s)

	// APPEND
	//newNode := AppendNewNode(messageNewNode)
	_ = AppendNewNode(messageNewNode)
	//js, _ := json.MarshalIndent(newNode, "", "    ")
	//s = "Appended new Node to the NodeList:\n" + string(js)
	//log.Info("ROUTINGNODE: HDLR          " + s)

	// RESPOND - Appended new Node to the NodeList
	s = "Appended new Node to the NodeList"
	_, err = rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Appended new Node to the NodeList"
	log.Info("ROUTINGNODE: HDLR  -sent   " + s)

	s = "END    handleBroadcastAddNewNode() - BROADCAST-ADD-NEW-NODE (BANN) - Adds a Node to the nodeList"
	log.Trace("ROUTINGNODE: HDLR   " + s)

}

// handleSendNodeList - SEND-NODELIST (SNL) - Sends the nodeList to another Node
func handleSendNodeList(rw *bufio.ReadWriter) {

	s := "START  handleSendNodeList() - SEND-NODELIST (SNL) - Sends the nodeList to another Node"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	// RESPOND - SEND NODELIST
	sendNodeList := GetNodeList()
	js, _ := json.Marshal(sendNodeList)
	s = string(js)
	_, err := rw.WriteString(s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	s = "Sent Nodelist to another node"
	log.Info("ROUTINGNODE: HDLR  -sent   " + s)

	s = "END    handleSendNodeList() - SEND-NODELIST (SNL) - Sends the nodeList to another Node"
	log.Trace("ROUTINGNODE: HDLR   " + s)

}

// handleBroadcastVerifiedBlock - BROADCAST-VERIFIED-BLOCK (BVB) - A Node verified the next block, get block and verify
func handleBroadcastVerifiedBlock(rw *bufio.ReadWriter) {

	s := "START  handleBroadcastVerifiedBlock() - BROADCAST-VERIFIED-BLOCK (BVB) - A Node verified the next block, get block and verify"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	s = "END   handleBroadcastVerifiedBlock() - BROADCAST-VERIFIED-BLOCK (BVB) - A Node verified the next block, get block and verify"
	log.Trace("ROUTINGNODE: HDLR   " + s)
}

// handleBroadcastConsensus - BROADCAST-CONSENSUS (BC) - 51% Consensus reached, get block to add to blockchain
func handleBroadcastConsensus(rw *bufio.ReadWriter) {

	s := "START  handleBroadcastConsensus() - BROADCAST-CONSENSUS (BC) - 51% Consensus reached, get block to add to blockchain"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	s = "END   handleBroadcastConsensus() - BROADCAST-CONSENSUS (BC) - 51% Consensus reached, get block to add to blockchain"
	log.Trace("ROUTINGNODE: HDLR   " + s)
}

// handleBroadcastTransactionRequest - BROADCAST-TRANSACTION-REQUEST (BTR) - Request from a Node to transfer jeffCoins to a jeffCoin Address
func handleBroadcastTransactionRequest(rw *bufio.ReadWriter) {

	s := "START  handleBroadcastTransactionRequest() - BROADCAST-TRANSACTION-REQUEST (BTR) - Request from a Node to transfer jeffCoins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	s = "Please enter the transactionRequestMessageSigned"
	log.Info("ROUTINGNODE: HDLR          " + s)
	returnMessage(s, rw)

	// RECEIVING - TRANSACTION REQUEST
	transactionRequestMessageSigned, err := rw.ReadString('\n')
	checkErr(err)
	transactionRequestMessageSigned = strings.Trim(transactionRequestMessageSigned, "\n ")
	s = "-rcvd   - TRANSACTION: " + transactionRequestMessageSigned
	log.Info("ROUTINGNODE: HDLR  " + s)

	// TRANSACTION REQUEST
	status := blockchain.TransactionRequest(transactionRequestMessageSigned)
	s = "The Status is: " + status
	log.Info("ROUTINGNODE: HDLR          " + s)
	returnMessage(s, rw)

	s = "END   handleBroadcastTransactionRequest() - BROADCAST-TRANSACTION-REQUEST (BTR) - Request from a Node to transfer jeffCoins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: HDLR   " + s)
}

// FROM WALLET I/F *******************************************************************************************************

// handleSendAddressBalance - SEND-ADDRESS-BALANCE (SAB) - Sends the jeffCoin balance for a jeffCoin Address
func handleSendAddressBalance(rw *bufio.ReadWriter) {

	s := "START  handleSendAddressBalance() - SEND-ADDRESS-BALANCE (SAB) - Sends the jeffCoin balance for a jeffCoin Address"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	s = "Please enter the jeffCoinAddress you want the balance for"
	log.Info("ROUTINGNODE: HDLR          " + s)
	returnMessage(s, rw)

	// RECEIVING JEFFCOINADDRESS
	jeffCoinAddress, err := rw.ReadString('\n')
	checkErr(err)
	jeffCoinAddress = strings.Trim(jeffCoinAddress, "\n ")
	s = "-rcvd   - jeffCoinAddress: " + jeffCoinAddress
	log.Info("ROUTINGNODE: HDLR  " + s)

	// GET ADDRESS BALANCE
	theBalance := blockchain.GetAddressBalance(jeffCoinAddress)
	s = "The balance for address " + jeffCoinAddress + " is " + theBalance
	log.Info("ROUTINGNODE: HDLR          " + s)
	returnMessage(s, rw)

	s = "END    handleSendAddressBalance() - SEND-ADDRESS-BALANCE (SAB) - Sends the jeffCoin balance for a jeffCoin Address"
	log.Trace("ROUTINGNODE: HDLR   " + s)
}

// handleTransactionRequest - TRANSACTION-REQUEST (TR) - Request from Wallet to transfer jeffCoins to a jeffCoin Address
func handleTransactionRequest(rw *bufio.ReadWriter) {

	s := "START  handleTransactionRequest() - TRANSACTION-REQUEST (TR) - Request from Wallet to Transfer jeffCoins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: HDLR   " + s)

	s = "Please enter the transactionRequestMessageSigned"
	log.Info("ROUTINGNODE: HDLR          " + s)
	returnMessage(s, rw)

	// RECEIVING TRANSACTION REQUEST
	transactionRequestMessageSigned, err := rw.ReadString('\n')
	checkErr(err)
	transactionRequestMessageSigned = strings.Trim(transactionRequestMessageSigned, "\n ")
	s = "-rcvd   - TRANSACTION: " + transactionRequestMessageSigned
	log.Info("ROUTINGNODE: HDLR  " + s)

	// BROADCAST TRANSACTION REQUEST TO ALL NODES
	// ???????????????????????????????????????????????

	// TRANSACTION REQUEST
	status := blockchain.TransactionRequest(transactionRequestMessageSigned)
	s = "The Status is: " + status
	log.Info("ROUTINGNODE: HDLR          " + s)
	returnMessage(s, rw)

	s = "END    handleTransactionRequest() - TRANSACTION-REQUEST (TR) - Request from Wallet to Transfer jeffCoins to a jeffCoin Address"
	log.Trace("ROUTINGNODE: HDLR   " + s)
}
