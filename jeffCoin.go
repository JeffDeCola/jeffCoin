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
	webserver "github.com/JeffDeCola/jeffCoin/webserver"
)

const (
	toolVersion = "1.0.1"
)

var genesisBlockchainPtr *bool
var yourIPPtr, yourWebPortPtr, yourTCPPortPtr *string
var networkIPPtr, networkTCPPortPtr *string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func startWebServer(ip string, webPort string) {

	// CREATE ROUTER
	myRouter := webserver.JeffsRouter()

	// LISTEN ON IP AND PORT
	fmt.Printf("Webserver listening on %s:%s\n\n", ip, webPort)
	log.Fatal(http.ListenAndServe(ip+":"+webPort, myRouter))

}

func startRoutingNode(ip string, tcpPort string) {

	// LISTEN ON IP AND PORT
	fmt.Printf("\nTCP Server listening on %s:%s\n\n", ip, tcpPort)
	server, err := net.Listen("tcp", ip+":"+tcpPort)
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
	// CREATEBLOCKCHAIN
	genesisBlockchainPtr = flag.Bool("genesis", false, "Create a blockchain")
	// YOUR IP
	yourIPPtr = flag.String("ip", "127.0.0.1", "Your IP")
	// YOUR WEB PORT
	yourWebPortPtr = flag.String("wp", "1234", "Your Web Port")
	// YOUR TCP PORT
	yourTCPPortPtr = flag.String("tp", "3333", "Your TCP Port")
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
	go startWebServer(*yourIPPtr, *yourWebPortPtr)

	// START ROUTING NODE (TCP SERVER)
	go startRoutingNode(*yourIPPtr, *yourTCPPortPtr)

	time.Sleep(2 * time.Second)

	// GENESIS BLOCKCHAIN
	if *genesisBlockchainPtr {
		firstTransaction := "create chain"
		difficulty := 1
		blockchain.GenesisBlockchain(firstTransaction, difficulty)
	} else {
		// LOAD BLOCKCHAIN FROM ANOTHER NODE
		err := blockchain.LoadBlockchain(*networkIPPtr, *networkTCPPortPtr)
		checkErr(err)
	}

	// PRESS RETURN TO EXIT
	fmt.Scanln()
	fmt.Println("\n...DONE\n")
}
