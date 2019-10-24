// jeffcoin.go

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"
	webserver "github.com/JeffDeCola/jeffCoin/webserver"
)

const (
	toolVersion = "1.0.3"
)

var genesisPtr *bool
var nodeIPPtr, nodeWebPortPtr, nodeTCPPortPtr *string
var networkIPPtr, networkTCPPortPtr *string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// startWebServer - Start the WebServer
func startWebServer(nodeIP string, nodeWebPort string) {

	s := "Web Server listening on " + nodeIP + ":" + nodeWebPort
	log.Debug("jeffCoin                   " + s)

	// CREATE ROUTER
	myRouter := webserver.JeffsRouter()

	// LISTEN ON IP AND PORT
	log.Fatal(http.ListenAndServe(nodeIP+":"+nodeWebPort, myRouter))

}

// startRoutingNode - Start the Routing Node (TCP Server)
func startRoutingNode(nodeIP string, nodeTCPPort string) {

	s := "TCP Server listening on " + nodeIP + ":" + nodeTCPPort
	log.Debug("jeffCoin                   " + s)

	// LISTEN ON IP AND PORT
	server, err := net.Listen("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)
	defer server.Close()

	// CREATE A CONNECTION FOR EACH REQUEST
	// Serve connections concurrently
	for {

		// Wait for a connection request
		conn, err := server.Accept()
		checkErr(err)

		go routingnode.HandleRequest(conn)
	}
}

func init() {

	// SET LOG LEVEL
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.TraceLevel)

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// VERSION FLAG
	versionPtr := flag.Bool("v", false, "prints current version")
	// CREATE FIRST NODE (GENISIS)
	genesisPtr = flag.Bool("genesis", false, "Create your first Node")
	// YOUR IP
	nodeIPPtr = flag.String("ip", "127.0.0.1", "Node IP")
	// YOUR WEB PORT
	nodeWebPortPtr = flag.String("webport", "1234", "Node Web Port")
	// YOUR TCP PORT
	nodeTCPPortPtr = flag.String("tcpport", "3333", "Node TCP Port")
	// NETWORK NODE IP
	networkIPPtr = flag.String("netip", "192.169.20.100", "Network IP")
	// NETWORK NODE TCP PORT
	networkTCPPortPtr = flag.String("netport", "3333", "Network TCP Port")
	// Parse the flags
	flag.Parse()

	// CHECK VERSION
	if *versionPtr {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func main() {

	fmt.Printf("\nSTARTING...\n")
	fmt.Println("Press return to exit\n")

	// START WEBSERVER
	go startWebServer(*nodeIPPtr, *nodeWebPortPtr)

	// START ROUTING NODE (TCP SERVER)
	go startRoutingNode(*nodeIPPtr, *nodeTCPPortPtr)

	// GIVE IT A SECOND
	time.Sleep(2 * time.Second)

	// LOAD thisNode
	routingnode.LoadThisNode(*nodeIPPtr, *nodeTCPPortPtr)

	// GENESIS wallet
	JeffCoinAddress := wallet.GenesisWallet()

	// IS THIS GENESIS?
	if *genesisPtr {

		// GENESIS blockchain
		firstTransaction := `
        {
            "ID": 0,
            "inputs": [
                {
                    "txID": 0,
                    "referenceTXID": -1,
                    "signature": ""
                }
            ],
            "outputs": [
                {
                    "jeffCoinAddress": "` + JeffCoinAddress + `",
                    "value": 1000
                }
            ]
        }`
		difficulty := 10
		blockchain.GenesisBlockchain(firstTransaction, difficulty)

		// GENESIS nodeList
		routingnode.GenesisNodeList()

	} else {

		// LOAD nodeList FROM THE NETWORK
		err := routingnode.LoadNodeList(*networkIPPtr, *networkTCPPortPtr)
		checkErr(err)

		// BROADCAST thisNode TO THE NETWORK
		err = routingnode.BroadcastThisNode()
		checkErr(err)

		// APPEND thisNode to nodeList
		routingnode.AppendThisNode()

		// LOAD blockchain FROM THE NETWORK
		err = blockchain.LoadBlockchain(*networkIPPtr, *networkTCPPortPtr)
		checkErr(err)

	}

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("\n...DONE\n")
}
