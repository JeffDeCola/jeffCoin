// jeffCoin handlers.go

package webserver

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"

	"github.com/gorilla/mux"
)

type htmlData struct {
	UserName string
	IP       string
	Port     string
}

// indexHandler - GET: /
func indexHandler(res http.ResponseWriter, req *http.Request) {

	s := "START: indexHandler - GET: /"
	log.Println("WEBSERVER:      " + s)

	t, err := template.ParseFiles("webserver/index.html")
	checkErr(err)

	// GET NODE
	thisNode := routingnode.GetNode("0")

	htmlTemplateData := htmlData{
		UserName: "John Smith",
		IP:       thisNode.IP,
		Port:     thisNode.Port,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END: indexHandler - GET: /"
	log.Println("WEBSERVER:      " + s)

}

// showBlockchainHandler - GET: /showblockchain
func showBlockchainHandler(res http.ResponseWriter, req *http.Request) {

	s := "START: showBlockchainHandler - GET: /showblockchain"
	log.Println("WEBSERVER:      " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET BLOCKCHAIN
	theBlockchain := blockchain.GetBlockchain()

	// RESPOND BLOCKCHAIN
	// s := "The entire Blockchain:"
	// respondMessage(s, res)
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	s = string(js)
	log.Println("WEBSERVER:      " + "Blockchain too long, not shown")
	io.WriteString(res, s+"\n")

	s = "END: showBlockchainHandler - GET: /showblockchain"
	log.Println("WEBSERVER:      " + s)

}

// showBlockHandler - GET: /showblock/{blockID}
func showBlockHandler(res http.ResponseWriter, req *http.Request) {

	s := "START: showBlockHandler - GET: /showblock/{blockID}"
	log.Println("WEBSERVER:      " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET BLOCK ID
	blockID := params["blockID"]

	// GET BLOCK
	theBlock := blockchain.GetBlock(blockID)

	// RESPOND BLOCK
	// s := "The Block you requested:"
	// respondMessage(s, res)
	js, _ := json.MarshalIndent(theBlock, "", "    ")
	s = string(js)
	respondMessage(s, res)

	s = "END: showBlockHandler - GET: /showblock/{blockID}"
	log.Println("WEBSERVER:      " + s)
}

// showNodeListHandler - GET: /shownodelist
func showNodeListHandler(res http.ResponseWriter, req *http.Request) {

	s := "START: showNodeListHandler - GET: /shownodelist"
	log.Println("WEBSERVER:      " + s)

	res.Header().Set("Content-Type", "application/json")

	// GET NODELIST
	theNodeList := routingnode.GetNodeList()

	// RESPOND NODELIST
	// s := "The entire Blockchain:"
	// respondMessage(s, res)
	js, _ := json.MarshalIndent(theNodeList, "", "    ")
	s = string(js)
	log.Println("WEBSERVER:      " + s)
	io.WriteString(res, s+"\n")

	s = "END: showNodeListHandler - GET: /shownodelist"
	log.Println("WEBSERVER:      " + s)

}

// showNodeHandler - GET: /shownode/{nodeID}
func showNodeHandler(res http.ResponseWriter, req *http.Request) {

	s := "START: showNodeHandler - GET: /shownode/{nodeID}"
	log.Println("WEBSERVER:      " + s)

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET NODE ID
	nodeID := params["nodeID"]

	// GET NODE
	theNode := routingnode.GetNode(nodeID)

	// RESPOND BLOCK
	// s := "The Block you requested:"
	// respondMessage(s, res)
	js, _ := json.MarshalIndent(theNode, "", "    ")
	s = string(js)
	respondMessage(s, res)

	s = "END: showNodeHandler - GET: /shownode/{nodeID}"
	log.Println("WEBSERVER:      " + s)
}

func respondMessage(s string, res http.ResponseWriter) {

	log.Println("WEBSERVER:      " + s)
	io.WriteString(res, s+"\n")

}
