// my-go-examples multi-node-blockchain-with-REST-and-tcp-ip.go

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/my-go-examples/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/blockchain"
	tcpserver "github.com/JeffDeCola/my-go-examples/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/tcpserver"
	webserver "github.com/JeffDeCola/my-go-examples/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/webserver"
)

const (
	ip          = "127.0.0.1"
	webPort     = "1234"
	tcpPort     = "3333"
	toolVersion = "1.0.1"
)

var createBlockchainPtr *bool

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func startWebServer() {

	// CREATE ROUTER
	myRouter := webserver.JeffsRouter()

	// LISTEN ON IP AND PORT
	fmt.Printf("Webserver listening on %s:%s\n\n", ip, webPort)
	log.Fatal(http.ListenAndServe(ip+":"+webPort, myRouter))

}

func startTCPServer() {

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

		go tcpserver.HandleRequest(conn)
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
	createBlockchainPtr = flag.Bool("b", false, "Create a blockchain")
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

	// CREATE BLOCKCHAIN
	if *createBlockchainPtr {
		firstTransaction := "create chain"
		difficulty := 1
		blockchain.CreateBlockchain(firstTransaction, difficulty)
	} else {
		// GET BLOCKCHAIN
		firstTransaction := "get chain"
		difficulty := 1
		blockchain.CreateBlockchain(firstTransaction, difficulty)
	}

	// START WEBSERVER
	go startWebServer()

	// START TCP SERVER
	go startTCPServer()

	// PRESS RETURN TO EXIT
	fmt.Println("\nPress return to exit\n")
	fmt.Scanln()
	fmt.Println("\n...DONE\n")
}
