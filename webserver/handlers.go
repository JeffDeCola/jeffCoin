// jeffCoin 5. WEBSERVER handlers.go

package webserver

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"

	"github.com/gorilla/mux"
)

type htmlData struct {
	ToolVersion string
	NodeName    string
	IP          string
	HTTPPort    string
	TCPPort     string
}

func logReceivedAPICommand() {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)
	s = "HTTP SERVER - RECEIVED API REST COMMAND"
	log.Info("WEBSERVER:                 " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)

}

func logDoneAPICommand() {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)
	s = "HTTP SERVER - DONE API REST COMMAND"
	log.Info("WEBSERVER:                 " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)

}

// indexHandler - GET: /
func indexHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)
	s = "HTTP SERVER - DISPLAY MAIN WEBPAGE"
	log.Info("WEBSERVER:                 " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)

	s = "START  indexHandler() - GET: /"
	log.Trace("WEBSERVER:          " + s)

	t, err := template.ParseFiles("webserver/index.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	htmlTemplateData := htmlData{
		NodeName:    thisNode.NodeName,
		ToolVersion: thisNode.ToolVersion,
		IP:          thisNode.IP,
		HTTPPort:    thisNode.HTTPPort,
		TCPPort:     thisNode.TCPPort,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    indexHandler() - GET: /"
	log.Trace("WEBSERVER:          " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY MAIN WEBPAGE"
	log.Info("WEBSERVER:                 " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                 " + s)

}

func respondMessage(s string, res http.ResponseWriter) {

	log.Info("WEBSERVER:                 " + s)
	io.WriteString(res, s+"\n")

}

// BLOCKCHAIN ************************************************************************************************************

// showBlockchainHandler - GET: /showblockchain
func showBlockchainHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBlockchainHandler() - GET: /showblockchain"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET BLOCKCHAIN
	theBlockchain := blockchain.GetBlockchain()

	// RESPOND with blockchain
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "Blockchain too long, not shown")
	io.WriteString(res, s+"\n")

	s = "END    showBlockchainHandler() - GET: /showblockchain"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showBlockHandler - GET: /showblock/{blockID}
func showBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBlockHandler() - GET: /showblock/{blockID}"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET BLOCK ID
	blockID := params["blockID"]

	// GET BLOCK
	theBlock := blockchain.GetBlock(blockID)

	// RESPOND with block
	js, _ := json.MarshalIndent(theBlock, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "Block too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showBlockHandler() - GET: /showblock/{blockID}"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showLockedBlockHandler - GET: /showlockedblock
func showLockedBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showLockedBlockHandler() - GET: /showlockedblock"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET lockedBlock
	theLockedBlock := blockchain.GetLockedBlock()

	// RESPOND with lockedBlock
	js, _ := json.MarshalIndent(theLockedBlock, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "lockedBlock too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showLockedBlockHandler() - GET: /showlockedblock"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showCurrentBlockHandler - GET: /showcurrentblock
func showCurrentBlockHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showCurrentBlockHandler() - GET: /showcurrentblock"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET currentBlock
	theCurrentBlock := blockchain.GetCurrentBlock()

	// RESPOND with currentBlock
	js, _ := json.MarshalIndent(theCurrentBlock, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "currentBlock too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showCurrentBlockHandler() - GET: /showcurrentblock"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// NODELIST **************************************************************************************************************

// showNodeListHandler - GET: /shownodelist
func showNodeListHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showNodeListHandler() - GET: /shownodelist"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET NODELIST
	theNodeList := routingnode.GetNodeList()

	// RESPOND with nodeList
	js, _ := json.MarshalIndent(theNodeList, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "NodeList too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showNodeListHandler() - GET: /shownodelist"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showNodeHandler - GET: /shownode/{nodeID}
func showNodeHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showNodeHandler() - GET: /shownode/{nodeID}"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET NODE ID
	nodeID := params["nodeID"]

	// GET NODE
	theNode := routingnode.GetNode(nodeID)

	// RESPOND with node
	js, _ := json.MarshalIndent(theNode, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "Node too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showNodeHandler() - GET: /shownode/{nodeID}"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showThisNodeHandler - GET: /showthisnode
func showThisNodeHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showThisNodeHandler() - GET: /showthisnode"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET thisNode
	gotThisNode := routingnode.GetThisNode()

	// RESPOND with thisNode
	js, _ := json.MarshalIndent(gotThisNode, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "thisNode too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showThisNodeHandler() - GET: /showthisnode"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// WALLET ****************************************************************************************************************

// showWalletHandler - GET: /showwallet
func showWalletHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showWalletHandler() - GET: /showwallet"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET wallet
	gotWallet := wallet.GetWallet()

	// RESPOND with wallet
	js, _ := json.MarshalIndent(gotWallet, "", "    ")
	s = string(js)
	log.Info("WEBSERVER:                 " + "wallet too long, not shown")
	io.WriteString(res, s+"\n")
	//respondMessage(s, res)

	s = "END    showWalletHandler() - GET: /showwallet"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showJeffCoinAddressHandler - GET: /showjeffcoinaddress
func showJeffCoinAddressHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showJeffCoinAddressHandler() - GET: /showjeffcoinaddress"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET wallet
	gotWallet := wallet.GetWallet()

	// GET jeffCoin Address
	jeffCoinAddress := gotWallet.JeffCoinAddress

	// RESPOND with jeffCoin Address
	s = jeffCoinAddress
	respondMessage(s, res)

	s = "END    showJeffCoinAddressHandler() - GET: /showjeffcoinaddress"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showBalanceHandler - GET: /showaddressbalance
func showBalanceHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showBalanceHandler() - GET: /showaddressbalance"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET nodeIP & nodeTCPPort from thisNode
	thisNode := routingnode.GetThisNode()
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// GET jeffCoinAddress from wallet
	gotWallet := wallet.GetWallet()
	jeffCoinAddress := gotWallet.JeffCoinAddress

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(nodeIP, nodeTCPPort, jeffCoinAddress)
	checkErr(err)

	// RESPOND with Address Balance
	s = gotAddressBalance
	respondMessage(s, res)

	s = "END    showBalanceHandler() - GET: /showaddressbalance"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// showAddressBalanceHandler - GET: /showaddressbalance/{jeffcoinaddress}
func showAddressBalanceHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  showAddressBalanceHandler() - GET: /showaddressbalance/{jeffcoinaddress}"
	log.Trace("WEBSERVER:          " + s)

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
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}

// transactionRequestHandler - GET: /transactionrequest/{destinationaddress}/{value}
func transactionRequestHandler(res http.ResponseWriter, req *http.Request) {

	logReceivedAPICommand()

	s := "START  transactionRequestHandler() - GET: /transactionrequest/{destinationaddress}/{value}"
	log.Trace("WEBSERVER:          " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET ADDRESS & VALUE
	destinationAddress := params["destinationaddress"]
	value := params["value"]

	// GET nodeIP & nodeTCPPort from thisNode
	thisNode := routingnode.GetThisNode()
	nodeIP := thisNode.IP
	nodeTCPPort := thisNode.TCPPort

	// GET wallet
	gotWallet := wallet.GetWallet()

	// GET sourceAddress FROM wallet
	sourceAddress := gotWallet.JeffCoinAddress

	// GET ENCODED KEYS FROM wallet
	privateKeyHex := gotWallet.PrivateKeyHex
	publicKeyHex := gotWallet.PublicKeyHex

	// DECODE KEYS
	privateKeyRaw, _ := wallet.DecodeKeys(privateKeyHex, publicKeyHex)

	// BUILD TRANSACTION REQUEST MESSAGE
	requestMessage := `
        { 
            "sourceAddress": "` + sourceAddress + `",
            "destinationAddress": "` + destinationAddress + `",
            "value" : ` + value + `
        }`

	// Make a long string - Remove /n and whitespace
	requestMessage = strings.Replace(requestMessage, "\n", "", -1)
	requestMessage = strings.Replace(requestMessage, " ", "", -1)

	// SIGN YOUR MESSAGE
	signature := wallet.CreateSignature(privateKeyRaw, requestMessage)

	// SIGNED TRANSACTION REQUEST MESSAGE
	transactionRequestMessageSigned := `
        {
            "requestMessage": ` + requestMessage + `,
            "signature" : "` + signature + `"
        }`
	// Make a long string - Remove /n and whitespace
	transactionRequestMessageSigned = strings.Replace(transactionRequestMessageSigned, "\n", "", -1)
	transactionRequestMessageSigned = strings.Replace(transactionRequestMessageSigned, " ", "", -1)

	// REQUEST TRANSACTION TO SEND COINS
	status, err := wallet.TransactionRequest(nodeIP, nodeTCPPort, transactionRequestMessageSigned)
	checkErr(err)

	// RESPOND with status
	s = status
	respondMessage(s, res)

	s = "END    transactionRequestHandler() - GET: /transactionrequest/{destinationaddress}/{value}"
	log.Trace("WEBSERVER:          " + s)

	logDoneAPICommand()

}
