// jeffCoin handlers.go

package webserver

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	"github.com/gorilla/mux"
)

type htmlData struct {
	UserName string
}

// GET
// Return entire Blockchain
func indexHandler(res http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles("webserver/index.html")
	checkErr(err)

	htmlTemplateData := htmlData{
		UserName: "John Smith",
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

}

// GET
func showBlockHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	// GET BLOCK ID
	blockID := params["blockID"]

	// GET BLOCK - ask interface
	theblock := blockchain.GetBlock(blockID)

	// RESPOND BLOCK
	// s := "The Block you requested:"
	// respondMessage(s, res)
	js, _ := json.MarshalIndent(theblock, "", "    ")
	s := string(js)
	respondMessage(s, res)

}

// GET
func showChainHandler(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// GET BLOCKCHAIN
	theBlockchain := blockchain.GetBlockchain()

	// RESPOND BLOCKCHAIN
	// s := "The entire Blockchain:"
	// respondMessage(s, res)
	js, _ := json.MarshalIndent(theBlockchain, "", "    ")
	s := string(js)
	log.Println("WEBSERVER:      " + "Blockchain too long, not shown")
	io.WriteString(res, s+"\n")

}

func respondMessage(s string, res http.ResponseWriter) {

	log.Println("WEBSERVER:      " + s)
	io.WriteString(res, s+"\n")

}
