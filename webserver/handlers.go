// jeffCoin 5. WEBSERVER handlers.go

package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"

	"github.com/gorilla/mux"
)

// SOME GLOBAL FUNCTIONS **************************************************************************************************

func logReceivedAPICommand() {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - RECEIVED API REST COMMAND"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

func logDoneAPICommand() {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DONE API REST COMMAND"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// checkAuthentication - Get session token from User's cookie and compare to stored one
func checkAuthentication(req *http.Request) bool {

	s := "START  checkAuthentication() - Get session token from User's cookie and compare to stored one"
	log.Debug("WEBSERVER:            " + s)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// GET SESSION TOKEN FROM USER
	cookie, err := req.Cookie("jeffCoin_session_token_" + thisNode.NodeName)
	if err != nil {

		s = "Cookie not present"
		log.Warn("WEBSERVER:                   " + s)

		s = "END    checkAuthentication() - Get session token from User's cookie and compare to stored one"
		log.Debug("WEBSERVER:            " + s)

		return false
	}
	sessionToken := cookie.Value

	s = "Session Cookie from User: " + sessionToken
	log.Info("WEBSERVER:                   " + s)

	s = "Session Cookie from storage: " + sessionTokenString
	log.Info("WEBSERVER:                   " + s)

	// CHECK THAT WE HAVE THE RIGHT SESSION TOKEN
	if sessionToken == sessionTokenString {

		s = "YOUR SESSION COOKIE MATCHES STORED COOKIE"
		log.Info("WEBSERVER:                   " + s)

		s = "END    checkAuthentication() - Get session token from User's cookie and compare to stored one"
		log.Debug("WEBSERVER:            " + s)

		return true
	}

	s = "END    checkAuthentication() - Get session token from User's cookie and compare to stored one"
	log.Debug("WEBSERVER:            " + s)

	return false

}

func respondMessage(s string, res http.ResponseWriter) {

	log.Info("WEBSERVER:                   " + s)
	io.WriteString(res, s+"\n")

}

func checkIfFounder() bool {

	// GET publicKeyHex FROM FIRST block IN blockchain
	firstBlock := blockchain.GetBlock("0")
	publicKeyHexFromFirstBlock := firstBlock.Transactions[0].Outputs[0].OutPubKey

	// GET publicKeyHex FROM THIS WALLET
	gotWallet := wallet.GetWallet()
	publicKeyHexFromWallet := gotWallet.PublicKeyHex

	if publicKeyHexFromFirstBlock == publicKeyHexFromWallet {
		return true
	}

	return false

}

func checkIfWalletOnly() bool {

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	return thisNode.WalletOnly
}

// BLOCKCHAIN ************************************************************************************************************

// showBlockchainHandler - GET: /showblockchain
func showBlockchainHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBlockchainHandler() - GET: /showblockchain"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - blockchain can not be shown"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// GET BLOCKCHAIN
		theBlockchain := blockchain.GetBlockchain()

		// RESPOND with blockchain
		js, _ := json.MarshalIndent(theBlockchain, "", "    ")
		s = string(js)
		log.Info("WEBSERVER:                   " + "Blockchain too long, not shown")
		io.WriteString(res, s+"\n")

	}

	s = "END    showBlockchainHandler() - GET: /showblockchain"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// showBlockHandler - GET: /showblock/{blockID}
func showBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBlockHandler() - GET: /showblock/{blockID}"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - block can not be shown"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// GET BLOCK ID
		blockID := params["blockID"]

		// GET BLOCK
		theBlock := blockchain.GetBlock(blockID)

		// RESPOND with block
		js, _ := json.MarshalIndent(theBlock, "", "    ")
		s = string(js)
		log.Info("WEBSERVER:                   " + "Block too long, not shown")
		io.WriteString(res, s+"\n")
		//respondMessage(s, res)
	}

	s = "END    showBlockHandler() - GET: /showblock/{blockID}"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// LOCKED BLOCK  *********************************************************************************************************

// showLockedBlockHandler - GET: /showlockedblock
func showLockedBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showLockedBlockHandler() - GET: /showlockedblock"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - lockedBlock can not be shown"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// GET lockedBlock
		theLockedBlock := blockchain.GetLockedBlock()

		// RESPOND with lockedBlock
		js, _ := json.MarshalIndent(theLockedBlock, "", "    ")
		s = string(js)
		log.Info("WEBSERVER:                   " + "lockedBlock too long, not shown")
		io.WriteString(res, s+"\n")
		//respondMessage(s, res)

	}

	s = "END    showLockedBlockHandler() - GET: /showlockedblock"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// appendLockedBlockHandler - GET: /appendlockedblock
func appendLockedBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  appendLockedBlockHandler() - GET: /appendlockedblock"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - can not append lockedBlock"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// ONLY FOUNDERS CAN MAKE THESE CHANGES
		if checkIfFounder() {

			// GET BlockID FROM LOCKED BLOCK
			theLockedBlock := blockchain.GetLockedBlock()

			// ADD lockedBlock TO THE blockchain
			s = "ADD lockedBlock TO THE blockchain. Adding block number " + fmt.Sprint(theLockedBlock.BlockID)
			log.Info("WEBSERVER:                   " + s)
			blockchain.AppendLockedBlock()

			// GET THE NEW BLOCK
			theNewBlock := blockchain.GetBlock(strconv.FormatInt(theLockedBlock.BlockID, 10))

			// RESPOND with New Block in the blockchain
			js, _ := json.MarshalIndent(theNewBlock, "", "    ")
			s = string(js)
			log.Info("WEBSERVER:                   " + "New Block too long, not shown")
			io.WriteString(res, s+"\n")
			//respondMessage(s, res)

		} else {

			// ONLY FOUNDER GENESIS CAN MAKE THESE CHANGES
			s = "Sorry, only Founders can make these changes"
			respondMessage(s, res)

		}

	}

	s = "END    appendLockedBlockHandler() - GET: /appendlockedblock"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// PENDING BLOCK  ********************************************************************************************************

// showPendingBlockHandler - GET: /showpendingblock
func showPendingBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showPendingBlockHandler() - GET: /showpendingblock"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - pendingBlock can not be shown"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// GET pendingBlock
		thePendingBlock := blockchain.GetPendingBlock()

		// RESPOND with pendingBlock
		js, _ := json.MarshalIndent(thePendingBlock, "", "    ")
		s = string(js)
		log.Info("WEBSERVER:                   " + "pendingBlock too long, not shown")
		io.WriteString(res, s+"\n")
		//respondMessage(s, res)

	}

	s = "END    showPendingBlockHandler() - GET: /showpendingblock"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// resetPendingBlockHandler - GET: /resetpendingblock
func resetPendingBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  resetPendingBlockHandler() - GET: /resetpendingblock"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - can not reset pendingBlock"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// ONLY FOUNDERS CAN MAKE THESE CHANGES
		if checkIfFounder() {

			// RESET pendingBlock
			s = "RESET pendingBlock"
			log.Info("WEBSERVER:                   " + s)
			blockchain.ResetPendingBlock()

			// GET NEW pendingBlock
			thePendingBlock := blockchain.GetPendingBlock()

			// RESPOND with pendingBlock
			js, _ := json.MarshalIndent(thePendingBlock, "", "    ")
			s = string(js)
			log.Info("WEBSERVER:                   " + "pendingBlock too long, not shown")
			io.WriteString(res, s+"\n")
			//respondMessage(s, res)

		} else {

			// ONLY FOUNDER GENESIS CAN MAKE THESE CHANGES
			s = "Sorry, only Founders can make these changes"
			respondMessage(s, res)

		}

	}
	s = "END    resetPendingBlockHandler() - GET: /resetpendingblock"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// lockPendingBlockHandler - GET: /lockpendingblock
func lockPendingBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  lockPendingBlockHandler() - GET: /lockpendingblock"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - can not move pendingBlock to lockedBlock"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// ONLY FOUNDERS CAN MAKE THESE CHANGES
		if checkIfFounder() {

			// GET DIFFICULTY FROM LAST LOCKED BLOCK
			theLockedBlock := blockchain.GetLockedBlock()

			// MOVE pendingBlock to the lockedBlock
			s = "MOVE pendingBlock to the lockedBlock with difficulty " + strconv.Itoa(theLockedBlock.Difficulty)
			log.Info("WEBSERVER:                   " + s)
			blockchain.LockPendingBlock(theLockedBlock.Difficulty)

			// GET THE NEW LOCKED BLOCK
			theLockedBlock = blockchain.GetLockedBlock()

			// RESPOND with lockedBlock
			js, _ := json.MarshalIndent(theLockedBlock, "", "    ")
			s = string(js)
			log.Info("WEBSERVER:                   " + "lockedBlock too long, not shown")
			io.WriteString(res, s+"\n")
			//respondMessage(s, res)

		} else {

			// ONLY FOUNDER GENESIS CAN MAKE THESE CHANGES
			s = "Sorry, only Founders can make these changes"
			respondMessage(s, res)

		}

	}
	s = "END    lockPendingBlockHandler() - GET: /lockpendingblock"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// NODELIST **************************************************************************************************************

// showNodeListHandler - GET: /shownodelist
func showNodeListHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showNodeListHandler() - GET: /shownodelist"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - nodeList can not be shown"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// GET NODELIST
		theNodeList := routingnode.GetNodeList()

		// RESPOND with nodeList
		js, _ := json.MarshalIndent(theNodeList, "", "    ")
		s = string(js)
		log.Info("WEBSERVER:                   " + "NodeList too long, not shown")
		io.WriteString(res, s+"\n")
		//respondMessage(s, res)

	}

	s = "END    showNodeListHandler() - GET: /shownodelist"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// showNodeHandler - GET: /shownode/{nodeID}
func showNodeHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showNodeHandler() - GET: /shownode/{nodeID}"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET NODE ID
	nodeID := params["nodeID"]

	// CHECK IF WALLET ONLY
	if checkIfWalletOnly() {

		s = "WALLET ONLY - node can not be shown"
		log.Info("WEBSERVER:                   " + s)
		io.WriteString(res, s+"\n")

	} else {

		// GET NODE
		theNode := routingnode.GetNode(nodeID)

		// RESPOND with node
		js, _ := json.MarshalIndent(theNode, "", "    ")
		s = string(js)
		log.Info("WEBSERVER:                   " + "Node too long, not shown")
		io.WriteString(res, s+"\n")
		//respondMessage(s, res)

	}

	s = "END    showNodeHandler() - GET: /shownode/{nodeID}"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// showThisNodeHandler - GET: /showthisnode
func showThisNodeHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showThisNodeHandler() - GET: /showthisnode"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET thisNode
	gotThisNode := routingnode.GetThisNode()

	// RESPOND with thisNode
	js, _ := json.MarshalIndent(gotThisNode, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "thisNode too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showThisNodeHandler() - GET: /showthisnode"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// WALLET (THIS NODE) ****************************************************************************************************

// showWalletHandler - GET: /showwallet
func showWalletHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showWalletHandler() - GET: /showwallet"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET wallet
	gotWallet := wallet.GetWallet()

	// RESPOND with wallet
	js, _ := json.MarshalIndent(gotWallet, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "wallet too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showWalletHandler() - GET: /showwallet"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// showJeffCoinAddressHandler - GET: /showjeffcoinaddress
func showJeffCoinAddressHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showJeffCoinAddressHandler() - GET: /showjeffcoinaddress"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET wallet
	gotWallet := wallet.GetWallet()

	// GET jeffCoin Address
	jeffCoinAddress := gotWallet.JeffCoinAddress

	// RESPOND with jeffCoin Address
	s = "\njeffCoinAddress (Not Used) = " + jeffCoinAddress
	respondMessage(s, res)

	// Get PubKeyHex
	theWallet := wallet.GetWallet()
	s = "\nPublicKeyHex = " + theWallet.PublicKeyHex

	// RESPOND with PubKeyHex
	respondMessage(s, res)

	s = "END    showJeffCoinAddressHandler() - GET: /showjeffcoinaddress"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// showBalanceHandler - GET: /showaddressbalance
func showBalanceHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBalanceHandler() - GET: /showaddressbalance"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// CHECK IF WALLET ONLY
	// GET IP & TCPPort from thisNode
	var IP, TCPPort string
	if checkIfWalletOnly() {

		IP = thisNode.NetworkIP
		TCPPort = thisNode.NetworkTCPPort
		s = "WALLET ONLY - Use network IP:Port to get balance"
		log.Info("WEBSERVER:                   " + s)

	} else {

		IP = thisNode.NodeIP
		TCPPort = thisNode.NodeTCPPort

	}

	// GET jeffCoinAddress from wallet
	gotWallet := wallet.GetWallet()
	jeffCoinAddress := gotWallet.PublicKeyHex

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(IP, TCPPort, jeffCoinAddress)
	checkErr(err)

	// RESPOND with Address Balance
	s = gotAddressBalance
	respondMessage(s, res)

	s = "END    showBalanceHandler() - GET: /showaddressbalance"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// transactionRequestHandler - GET: /transactionrequest/{destinationaddress}/{value}
func transactionRequestHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  transactionRequestHandler() - GET: /transactionrequest/{destinationaddress}/{value}"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/transactionrequest.html")
	checkErr(err)

	//res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// ------------------------------------
	// GET ADDRESS & VALUE (COMMA DELIMITED)
	destinationAddressComma := params["destinationaddress"]
	valueComma := params["value"]

	// -----------------------------------
	// GET IP & TCPPort
	// GET THIS NODE
	thisNode := routingnode.GetThisNode()
	// CHECK IF WALLET ONLY
	// GET IP & TCPPort from thisNode
	var IP, TCPPort string
	if checkIfWalletOnly() {
		IP = thisNode.NetworkIP
		TCPPort = thisNode.NetworkTCPPort
		s = "WALLET ONLY - Use network IP:Port for transaction request"
		log.Info("WEBSERVER:                   " + s)
	} else {
		IP = thisNode.NodeIP
		TCPPort = thisNode.NodeTCPPort
	}

	// ------------------------------------
	// GET sourceAddress FROM wallet
	gotWallet := wallet.GetWallet()
	sourceAddress := gotWallet.PublicKeyHex
	// GET ENCODED KEYS FROM wallet
	privateKeyHex := gotWallet.PrivateKeyHex
	// publicKeyHex := gotWallet.PublicKeyHex
	// DECODE KEYS
	// privateKeyRaw, _ := wallet.DecodeKeys(privateKeyHex, publicKeyHex)

	// ------------------------------------
	// BUILD TRANSACTION REQUEST MESSAGE
	// The destinationAddress and Value are comma delimiter
	// Clean/remove the comma into a space
	destinationAddressSpace := strings.Replace(destinationAddressComma, ",", " ", -1)
	valueSpace := strings.Replace(valueComma, ",", " ", -1)
	// Convert 'cleaned' comma separated string to slice
	destinationAddressSlice := strings.Fields(destinationAddressSpace)
	valueSlice := strings.Fields(valueSpace)

	// Check that both slices have the same amount
	// -----------------------------------------
	var status string
	if cap(destinationAddressSlice) == cap(valueSlice) {
		// Build destinations first
		destinations := "["
		for i := range destinationAddressSlice {
			destinations = destinations + `
            {
                "destinationAddress": "` + destinationAddressSlice[i] + `",
                "value": ` + valueSlice[i] + `
            }`
			// Add comma between destinations
			if i != (len(destinationAddressSlice) - 1) {
				destinations = destinations + ","
			}
		}
		destinations = destinations + "]"
		// Build tx request message
		txRequestMessage := `
        { 
            "sourceAddress": "` + sourceAddress + `",
            "destinations": ` + destinations + `
        }`
		// Make a long string - Remove /n and whitespace
		txRequestMessage = strings.Replace(txRequestMessage, "\n", "", -1)
		txRequestMessage = strings.Replace(txRequestMessage, " ", "", -1)

		// ------------------------------------
		// SIGN YOUR MESSAGE
		signature := wallet.CreateSignature(privateKeyHex, txRequestMessage)

		// ------------------------------------
		// BUILD SIGNED TRANSACTION REQUEST MESSAGE
		txRequestMessageSigned := `
        {
            "txRequestMessage": ` + txRequestMessage + `,
            "signature" : "` + signature + `"
        }`
		// Make a long string - Remove /n and whitespace
		txRequestMessageSigned = strings.Replace(txRequestMessageSigned, "\n", "", -1)
		txRequestMessageSigned = strings.Replace(txRequestMessageSigned, " ", "", -1)

		// ------------------------------------
		// REQUEST TRANSACTION
		status, err = wallet.TransactionRequest(IP, TCPPort, txRequestMessageSigned)
		checkErr(err)
	} else {
		status = "Field are incorrect. Try again."
	}

	htmlTemplateData := htmlTransactionRequestData{
		Status: status,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    transactionRequestHandler() - GET: /transactionrequest/{destinationaddress}/{value}"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}

// WALLET (OTHER) ********************************************************************************************************

// showAddressBalanceHandler - GET: /showaddressbalance/{jeffcoinaddress}
func showAddressBalanceHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showAddressBalanceHandler() - GET: /showaddressbalance/{jeffcoinaddress}"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET jeffCoin Address
	jeffCoinAddress := params["jeffCoinAddress"]

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// CHECK IF WALLET ONLY
	// GET IP & TCPPort from thisNode
	var IP, TCPPort string
	if checkIfWalletOnly() {

		IP = thisNode.NetworkIP
		TCPPort = thisNode.NetworkTCPPort
		s = "WALLET ONLY - Use network IP:Port to get balance"
		log.Info("WEBSERVER:                   " + s)

	} else {

		IP = thisNode.NodeIP
		TCPPort = thisNode.NodeTCPPort

	}

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(IP, TCPPort, jeffCoinAddress)
	checkErr(err)

	// RESPOND with Address Balance
	s = gotAddressBalance
	respondMessage(s, res)

	s = "END    showAddressBalanceHandler() - GET: /showaddressbalance/{jeffcoinaddress}"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}
