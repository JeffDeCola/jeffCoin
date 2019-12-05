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
	wallet "github.com/JeffDeCola/jeffCoin/wallet"
	webserver "github.com/JeffDeCola/jeffCoin/webserver"
)

const (
	toolVersion = "1.2.0"
)

var genesisPtr *bool
var nodeNamePtr, nodeIPPtr, nodeHTTPPortPtr, nodeTCPPortPtr *string
var networkIPPtr, networkTCPPortPtr *string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
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
	// CREATE FIRST NODE (GENISIS)
	genesisPtr = flag.Bool("genesis", false, "Create your first Node")
	// LOG LEVEL
	logLevelPtr := flag.String("loglevel", "trace", "LogLevel (trace or info)")
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

	s := "START WEBSERVER (HTTP SERVER)"
	log.Info("MAIN:                      " + s)
	go webserver.StartHTTPServer(*nodeIPPtr, *nodeHTTPPortPtr)

	s = "START ROUTING NODE (TCP SERVER)"
	log.Info("MAIN:                      " + s)
	go routingnode.StartRoutingNode(*nodeIPPtr, *nodeTCPPortPtr)

	// GIVE IT A SECOND
	time.Sleep(1 * time.Second)

	s = "LOAD thisNode"
	log.Info("MAIN:                      " + s)
	routingnode.LoadThisNode(*nodeIPPtr, *nodeHTTPPortPtr, *nodeTCPPortPtr, *nodeNamePtr, toolVersion)
	time.Sleep(100000 * time.Minute)

	s = "GENESIS wallet"
	log.Info("MAIN:                      " + s)
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
		difficulty := 8
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

	// KICK OFF MASTER CONTROL
	go masterFlowControl(*genesisPtr)

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Printf("\n...DONE\n")
}
