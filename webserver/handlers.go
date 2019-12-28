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

type htmlIndexData struct {
	NodeName     string
	PublicKeyHex string
	Balance      string
	ToolVersion  string
	IP           string
	HTTPPort     string
	TCPPort      string
}

type htmlAPIData struct {
	NodeName string
}

type htmlSendData struct {
	NodeName     string
	PublicKeyHex string
	Balance      string
}

type htmlConfirmData struct {
	NodeName                  string
	PublicKeyHex              string
	DestinationAddressComma   string
	DestinationAddressNewLine string
	ValueComma                string
	ValueNewLine              string
}

type htmlTransactionRequestData struct {
	Status string
}

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

// HTML PAGES *************************************************************************************************************

// indexHandler - GET: /
func indexHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY MAIN WEBPAGE"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  indexHandler() - GET: /"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/index.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// GET WALLET
	theWallet := wallet.GetWallet()

	// GET nodeIP & nodeTCPPort from thisNode
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// GET address PublicKeyHex from wallet
	addressPublicKeyHex := theWallet.PublicKeyHex

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(nodeIP, nodeTCPPort, addressPublicKeyHex)
	gotAddressBalance = strings.Trim(gotAddressBalance, "\"")
	checkErr(err)
	gotAddressBalanceInt, err := strconv.ParseFloat(gotAddressBalance, 64)
	checkErr(err)
	balance := gotAddressBalanceInt / float64(1000)
	gotAddressBalance = strconv.FormatFloat(balance, 'f', 3, 64)

	htmlTemplateData := htmlIndexData{
		NodeName:     thisNode.NodeName,
		PublicKeyHex: addressPublicKeyHex,
		Balance:      gotAddressBalance,
		ToolVersion:  thisNode.ToolVersion,
		IP:           thisNode.IP,
		HTTPPort:     thisNode.HTTPPort,
		TCPPort:      thisNode.TCPPort,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    indexHandler() - GET: /"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY MAIN WEBPAGE"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// apiHandler - GET: /api
func apiHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  apiHandler() - GET: /api"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/api.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	htmlTemplateData := htmlAPIData{
		NodeName: thisNode.NodeName,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    apiHandler() - GET: /api"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// sendHandler - GET: /send
func sendHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  sendHandler() - GET: /send"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/send.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// GET WALLET
	theWallet := wallet.GetWallet()

	// GET nodeIP & nodeTCPPort from thisNode
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// GET address PublicKeyHex from wallet
	addressPublicKeyHex := theWallet.PublicKeyHex

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(nodeIP, nodeTCPPort, addressPublicKeyHex)
	gotAddressBalance = strings.Trim(gotAddressBalance, "\"")
	checkErr(err)
	gotAddressBalanceInt, err := strconv.ParseFloat(gotAddressBalance, 64)
	checkErr(err)
	balance := gotAddressBalanceInt / float64(1000)
	gotAddressBalance = strconv.FormatFloat(balance, 'f', 3, 64)

	htmlTemplateData := htmlSendData{
		NodeName:     thisNode.NodeName,
		PublicKeyHex: addressPublicKeyHex,
		Balance:      gotAddressBalance,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    sendHandler() - GET: /send"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// confirmHandler - GET: /confirm
func confirmHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY CONFIRM SEND"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  confirmHandler() - GET: /confirm"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/confirm.html")
	checkErr(err)

	// GET THE PARAMATERS SENT VIA POST FORM
	// Parses the request body
	// It may or may not have a comma
	req.ParseForm()
	destinationAddressComma := req.Form.Get("DestinationAddress")
	valueComma := req.Form.Get("Value")

	// ADD NEWLINE TO COMMAS (Just added whitespace but would like to figure this out)
	destinationAddressNewline := strings.Replace(destinationAddressComma, ",", ", ", -1)
	valueNewLine := strings.Replace(valueComma, ",", ", ", -1)
	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// GET WALLET
	theWallet := wallet.GetWallet()

	htmlTemplateData := htmlConfirmData{
		NodeName:                  thisNode.NodeName,
		PublicKeyHex:              theWallet.PublicKeyHex,
		DestinationAddressComma:   destinationAddressComma,
		DestinationAddressNewLine: destinationAddressNewline,
		ValueComma:                valueComma,
		ValueNewLine:              valueNewLine,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    confirmHandler() - GET: /confirm"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE CONFIRM SEND"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// BLOCKCHAIN ************************************************************************************************************

// showBlockchainHandler - GET: /showblockchain
func showBlockchainHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBlockchainHandler() - GET: /showblockchain"
	log.Debug("WEBSERVER:            " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET BLOCKCHAIN
	theBlockchain := blockchain.GetBlockchain()

	// RESPOND with blockchain
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "Blockchain too long, not shown")
	io.WriteString(res, s+"\n")

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

	// GET lockedBlock
	theLockedBlock := blockchain.GetLockedBlock()

	// RESPOND with lockedBlock
	js, _ := json.MarshalIndent(theLockedBlock, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "lockedBlock too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

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

	// GET pendingBlock
	thePendingBlock := blockchain.GetPendingBlock()

	// RESPOND with pendingBlock
	js, _ := json.MarshalIndent(thePendingBlock, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "pendingBlock too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

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

	// GET NODELIST
	theNodeList := routingnode.GetNodeList()

	// RESPOND with nodeList
	js, _ := json.MarshalIndent(theNodeList, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "NodeList too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

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

	// GET NODE
	theNode := routingnode.GetNode(nodeID)

	// RESPOND with node
	js, _ := json.MarshalIndent(theNode, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                   " + "Node too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

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

	// GET nodeIP & nodeTCPPort from thisNode
	thisNode := routingnode.GetThisNode()
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// GET jeffCoinAddress from wallet
	gotWallet := wallet.GetWallet()
	jeffCoinAddress := gotWallet.PublicKeyHex

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(nodeIP, nodeTCPPort, jeffCoinAddress)
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

	// ------------------------------------
	// GET nodeIP & nodeTCPPort from thisNode
	thisNode := routingnode.GetThisNode()
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

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
	// ------------------------------------------------
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
		status, err = wallet.TransactionRequest(nodeIP, nodeTCPPort, txRequestMessageSigned)
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

	// GET nodeIP & nodeTCPPort from thisNode
	thisNode := routingnode.GetThisNode()
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(nodeIP, nodeTCPPort, jeffCoinAddress)
	checkErr(err)

	// RESPOND with Address Balance
	s = gotAddressBalance
	respondMessage(s, res)

	s = "END    showAddressBalanceHandler() - GET: /showaddressbalance/{jeffcoinaddress}"
	log.Debug("WEBSERVER:            " + s)

	logDoneAPICommand()

}
