// jeffcoin.go

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	testdata "github.com/JeffDeCola/jeffCoin/testdata"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"
	webserver "github.com/JeffDeCola/jeffCoin/webserver"
)

const (
	toolVersion = "1.3.1"
)

var genesisPtr, testPtr, gcePtr *bool
var nodeNamePtr, nodeIPPtr, nodeHTTPPortPtr, nodeTCPPortPtr *string
var networkIPPtr, networkTCPPortPtr *string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func genesisNode(jeffCoinAddress string) {

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
                    "jeffCoinAddress": "` + jeffCoinAddress + `",
                    "value": 100000
                }
            ]
        }`
	difficulty := 8
	s := "GENESIS blockchain"
	log.Info("MAIN:                        " + s)
	blockchain.GenesisBlockchain(firstTransaction, difficulty)

	// GENESIS nodeList
	s = "GENESIS nodeList"
	log.Info("MAIN:                        " + s)
	routingnode.GenesisNodeList()

}

func newNode() {

	// REQUEST nodeList FROM THE NETWORK
	s := "REQUEST nodeList FROM THE NETWORK"
	log.Info("MAIN:                        " + s)
	err := routingnode.RequestNodeList(*networkIPPtr, *networkTCPPortPtr)
	checkErr(err)

	// BROADCAST thisNode TO THE NETWORK
	s = "BROADCAST thisNode TO THE NETWORK"
	log.Info("MAIN:                        " + s)
	err = routingnode.BroadcastThisNode()
	checkErr(err)

	// APPEND thisNode to nodeList
	s = "APPEND thisNode to nodeList"
	log.Info("MAIN:                        " + s)
	routingnode.AppendThisNode()

	// REQUEST blockchain FROM THE NETWORK
	s = "REQUEST blockchain FROM THE NETWORK"
	log.Info("MAIN:                        " + s)
	err = blockchain.RequestBlockchain(*networkIPPtr, *networkTCPPortPtr)
	checkErr(err)

}

// masterFlowControl - Controls the mining of blocks and addint to the chain
func masterFlowControl(genesis bool) {

	// START IT UP
	if genesis {

	}

}

func init() {

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// VERSION FLAG
	versionPtr := flag.Bool("v", false, "prints current version")
	// TEST FLAG
	testPtr = flag.Bool("test", false, "Loads the blockchain with test data")
	// CREATE FIRST NODE (GENISIS)
	genesisPtr = flag.Bool("genesis", false, "Create your first Node")
	// LOG LEVEL
	logLevelPtr := flag.String("loglevel", "trace", "LogLevel (trace or info)")
	// GCE
	gcePtr = flag.Bool("gce", false, "Is this Node on GCE")
	// YOUR IP
	nodeIPPtr = flag.String("ip", "127.0.0.1", "Node IP")
	// YOUR WEB PORT
	nodeHTTPPortPtr = flag.String("httpport", "1234", "Node Web Port")
	// YOUR TCP PORT
	nodeTCPPortPtr = flag.String("tcpport", "3333", "Node TCP Port")
	// NETWORK NODE IP
	networkIPPtr = flag.String("netip", "192.169.20.100", "Network IP")
	// NETWORK NODE TCP PORT
	networkTCPPortPtr = flag.String("netport", "3333", "Network TCP Port")
	// NODE NAME
	nodeNamePtr = flag.String("nodename", "Monkey", "Node Name")
	// Parse the flags
	flag.Parse()

	// SET LOG LEVEL
	if *logLevelPtr == "trace" {
		log.SetLevel(log.TraceLevel)
	} else if *logLevelPtr == "info" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.Error("Error setting log levels")
		os.Exit(0)
	}

	// CHECK VERSION
	if *versionPtr {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func main() {

	fmt.Printf("\nSTART...\n")
	fmt.Printf("Press return to exit\n\n")

	// START WEBSERVER (HTTP SERVER)
	s := "START WEBSERVER (HTTP SERVER)"
	log.Info("MAIN:                        " + s)
	if *gcePtr {
		go webserver.StartHTTPServer("0.0.0.0", *nodeHTTPPortPtr)
	} else {
		go webserver.StartHTTPServer(*nodeIPPtr, *nodeHTTPPortPtr)
	}

	// GIVE IT A SECOND
	time.Sleep(1 * time.Second)

	// START ROUTING NODE (TCP SERVER)
	s = "START ROUTING NODE (TCP SERVER)"
	log.Info("MAIN:                        " + s)
	if *gcePtr {
		go routingnode.StartRoutingNode("0.0.0.0", *nodeTCPPortPtr)
	} else {
		go routingnode.StartRoutingNode(*nodeIPPtr, *nodeTCPPortPtr)
	}

	// GIVE IT A SECOND
	time.Sleep(1 * time.Second)

	// LOAD thisNode
	s = "LOAD thisNode"
	log.Info("MAIN:                        " + s)
	routingnode.LoadThisNode(*nodeIPPtr, *nodeHTTPPortPtr, *nodeTCPPortPtr, *nodeNamePtr, toolVersion)

	var jeffCoinAddress string

	// DO YOU ALREADY HAVE A WALLET
	if _, err := os.Stat("wallet/" + *nodeNamePtr + "-wallet.json"); err == nil {
		// READ existing wallet from a file (Keys and jeffCoin Address)
		s = "READ existing wallet from a file (Keys and jeffCoin Address)"
		log.Info("MAIN:                        " + s)
		jeffCoinAddress = wallet.ReadWalletFile(*nodeNamePtr)
	} else {
		// GENESIS wallet - Creates the wallet and write to file (Keys and jeffCoin Address)
		s = "GENESIS wallet - Creates the wallet and write to file (Keys and jeffCoin Address)"
		log.Info("MAIN:                        " + s)
		jeffCoinAddress = wallet.GenesisWallet(*nodeNamePtr)
	}

	// CREATE GENESIS NODE OR A NEW NODE
	if *genesisPtr {
		genesisNode(jeffCoinAddress)
	} else {
		newNode()
	}

	// LOAD BLOCKCHAIN WITH TEST DATA
	if *testPtr {
		testdata.LoadTestDatatoBlockchain()
	}

	// KICK OFF MASTER CONTROL
	s = "KICK OFF MASTER CONTROL"
	log.Info("MAIN:                        " + s)
	go masterFlowControl(*genesisPtr)

	// PRESS RETURN TO EXIT
	s = "PRESS RETURN TO EXIT"
	log.Info("MAIN:                        " + s)
	fmt.Scanln()
	fmt.Printf("\n...DONE\n")
}
