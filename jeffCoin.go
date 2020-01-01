// jeffcoin.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"

	blockchain "github.com/JeffDeCola/jeffCoin/blockchain"
	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	testblockchain "github.com/JeffDeCola/jeffCoin/testblockchain"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"
	webserver "github.com/JeffDeCola/jeffCoin/webserver"
)

const (
	toolVersion = "1.3.3"
)

var genesisPtr, testblockchainPtr, gcePtr, walletPtr *bool
var nodeNamePtr, nodeIPPtr, nodeHTTPPortPtr, nodeTCPPortPtr *string
var networkIPPtr, networkTCPPortPtr *string
var passwordPtr *string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func genesisNode(jeffCoinAddress string) {

	genesisBlockDataString := `
    {
        "blockID": 0,
        "timestamp": "",
        "transactions": [
            {
                "txID": 0,
                "inputs": [
                    {
                        "refTxID": -1,
                        "inPubKey": "",
                        "signature": "Hello World, Welcome to the first transaction and block of jeffCoin"
                    }
                ],
                "outputs": [
                    {
                        "outPubKey": "` + jeffCoinAddress + `",
                        "value": 100000000
                    }
                ]
            }
        ],
        "hash": "",
        "prevhash": "",
        "difficulty": 8,
        "nonce": ""
    }`

	// GENESIS blockchain
	s := "GENESIS blockchain"
	log.Info("MAIN:                        " + s)
	blockchain.GenesisBlockchain(genesisBlockDataString)

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

// ContextHook ...
type ContextHook struct{}

// Levels ...
func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire ...
func (hook ContextHook) Fire(entry *log.Entry) error {
	if pc, file, line, ok := runtime.Caller(10); ok {
		funcName := runtime.FuncForPC(pc).Name()

		entry.Data["source"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
	}

	return nil
}

func init() {

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

	// GCE
	gcePtr = flag.Bool("gce", false, "Is this Node on GCE")
	// CREATE FIRST NODE (GENISIS)
	genesisPtr = flag.Bool("genesis", false, "Create your first Node")
	// LOG LEVEL
	logLevelPtr := flag.String("loglevel", "info", "LogLevel (info, debug or trace)")
	// NETWORK NODE IP
	networkIPPtr = flag.String("networkip", "127.0.0.1", "Network IP")
	// NETWORK NODE TCP PORT
	networkTCPPortPtr = flag.String("networktcpport", "3000", "Network TCP Port")
	// YOUR NODE WEB PORT
	nodeHTTPPortPtr = flag.String("nodehttpport", "2001", "Node Web Port")
	// YOUR NODE IP
	nodeIPPtr = flag.String("nodeip", "127.0.0.1", "Node IP")
	// NODE NAME
	nodeNamePtr = flag.String("nodename", "Jeff", "Node Name")
	// YOUR NODE TCP PORT
	nodeTCPPortPtr = flag.String("nodetcpport", "3001", "Node TCP Port")
	// PASSWORD
	passwordPtr = flag.String("password", "yourpassword", "Set/Reset your Password")
	// TEST FLAG
	testblockchainPtr = flag.Bool("testblockchain", false, "Loads the blockchain with test data")
	// VERSION FLAG
	versionPtr := flag.Bool("v", false, "prints current version")
	// VERSION FLAG
	walletPtr = flag.Bool("wallet", false, "Only the wallet and webserver (GUI/API)")

	// Parse the flags
	flag.Parse()

	// SET LOG LEVEL
	if *logLevelPtr == "trace" {
		log.SetLevel(log.TraceLevel)
	} else if *logLevelPtr == "debug" {
		log.SetLevel(log.DebugLevel)
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

	// ERASE -------------------------------------
	test := false
	if test {
		type erasewalletStruct struct {
			PrivateKeyHex   string `json:"privateKeyHex"`
			PublicKeyHex    string `json:"publicKeyHex"`
			JeffCoinAddress string `json:"jeffCoinAddress"`
		}

		var walletStruct erasewalletStruct

		// READ WALLET STRUCT TO JSON FILE
		filename := "wallet/" + *nodeNamePtr + "-wallet.json"
		filedata, _ := ioutil.ReadFile(filename)
		_ = json.Unmarshal([]byte(filedata), &walletStruct)
		s := "Read wallet from " + filename
		log.Info("WALLET:      GUTS            " + s)

		js, _ := json.MarshalIndent(walletStruct, "", "    ")
		log.Info("\n\n" + string(js) + "\n\n")

		// ENCRYPT privateKeyHex using key
		keyText := "myverystrongpasswordo32bitlength"
		keyByte := []byte(keyText)
		additionalData := "Jeff's additional data for authorization"

		privateKeyHexEncrypted := wallet.EncryptAES(keyByte, walletStruct.PrivateKeyHex, additionalData)
		publicKeyHex := walletStruct.PublicKeyHex
		jeffCoinAddressHex := walletStruct.JeffCoinAddress

		// ENCRYPTED STRUCT FOR FILE
		walletStructEncrypted := erasewalletStruct{privateKeyHexEncrypted, publicKeyHex, jeffCoinAddressHex}

		js, _ = json.MarshalIndent(walletStructEncrypted, "", "    ")
		log.Info("\n\n" + string(js) + "\n\n")

		// WRITE WALLET STRUCT TO JSON FILE
		filedata, _ = json.MarshalIndent(walletStructEncrypted, "", " ")
		filename = "wallet/" + *nodeNamePtr + "-wallet.json"
		_ = ioutil.WriteFile(filename, filedata, 0644)
		s = "Wrote wallet to " + filename
		log.Info("WALLET:      GUTS            " + s)

		os.Exit(0)
	}
	// ERASE -------------------------------------

	// CHECK IF SET PASSWORD
	// If default, see if you already have a password
	if *passwordPtr == "yourpassword" {
		// Check if you have a password in file
		if _, err := os.Stat("credentials/" + *nodeNamePtr + "-password.json"); err == nil {
			// READ existing password from a file and put in struct
			s := "READ existing password from a file and put in struct"
			log.Info("MAIN:                        " + s)
			_ = webserver.ReadPasswordFile(*nodeNamePtr)
		} else {
			fmt.Println("Please set a password")
			os.Exit(0)
		}
	} else {
		// Writes the password to file (AES-256 encryption)
		s := "Writes the password to file (AES-256 encryption)"
		log.Info("MAIN:                        " + s)
		_ = webserver.WritePasswordFile(*nodeNamePtr, *passwordPtr)
	}

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
	// Do not start if wallet only
	if !*walletPtr {
		s = "START ROUTING NODE (TCP SERVER)"
		log.Info("MAIN:                        " + s)
		if *gcePtr {
			go routingnode.StartRoutingNode("0.0.0.0", *nodeTCPPortPtr)
		} else {
			go routingnode.StartRoutingNode(*nodeIPPtr, *nodeTCPPortPtr)
		}
	}

	// GIVE IT A SECOND
	time.Sleep(1 * time.Second)

	// LOAD thisNode
	s = "LOAD thisNode"
	log.Info("MAIN:                        " + s)
	myNode := routingnode.NodeStruct{
		NodeName:       *nodeNamePtr,
		NodeIP:         *nodeIPPtr,
		NodeHTTPPort:   *nodeHTTPPortPtr,
		NodeTCPPort:    *nodeTCPPortPtr,
		NetworkIP:      *networkIPPtr,
		NetworkTCPPort: *networkTCPPortPtr,
		TestBlockChain: *testblockchainPtr,
		WalletOnly:     *walletPtr,
		ToolVersion:    toolVersion,
	}
	myNode.LoadThisNode()

	// DO YOU ALREADY HAVE A WALLET
	var jeffCoinAddress string
	if _, err := os.Stat("wallet/" + *nodeNamePtr + "-wallet.json"); err == nil {
		// READ existing wallet from a file (Keys and jeffCoin Address) and put in struct
		s = "READ existing wallet from a file (Keys and jeffCoin Address) and put in struct"
		log.Info("MAIN:                        " + s)
		jeffCoinAddress = wallet.ReadWalletFile(*nodeNamePtr)
	} else {
		// GENESIS wallet - Creates the wallet and write to file (Keys and jeffCoin Address)
		s = "GENESIS wallet - Creates the wallet and write to file (Keys and jeffCoin Address)"
		log.Info("MAIN:                        " + s)
		jeffCoinAddress = wallet.GenesisWallet(*nodeNamePtr)
	}
	s = "Not using jeffCoinAddress " + jeffCoinAddress
	log.Info("MAIN:                        " + s)

	// CREATE GENESIS NODE OR A NEW NODE
	// Do not start if wallet only
	if !*walletPtr {
		if *genesisPtr {
			theWallet := wallet.GetWallet()
			genesisNode(theWallet.PublicKeyHex)
		} else {
			newNode()
		}
	}

	// LOAD BLOCKCHAIN WITH TEST DATA
	if *testblockchainPtr {
		testblockchain.LoadTestDatatoBlockchain()
		// SHOW BALANCES
		balance := blockchain.GetAddressBalance(testblockchain.MockFoundersPubKey)
		log.Info("\n\nThe balance for " + testblockchain.MockFoundersPubKey[0:40] + "... (Address) is " + balance + "\n\n")
		balance = blockchain.GetAddressBalance(testblockchain.MockJeffPubKey)
		log.Info("\n\nThe balance for " + testblockchain.MockJeffPubKey[0:40] + "... (Address) is " + balance + "\n\n")
		balance = blockchain.GetAddressBalance(testblockchain.MockMattPubKey)
		log.Info("\n\nThe balance for " + testblockchain.MockMattPubKey[0:40] + "... (Address) is " + balance + "\n\n")
		balance = blockchain.GetAddressBalance(testblockchain.MockJillPubKey)
		log.Info("\n\nThe balance for " + testblockchain.MockJillPubKey[0:40] + "... (Address) is " + balance + "\n\n")
		balance = blockchain.GetAddressBalance(testblockchain.MockCoinVaultPubKey)
		log.Info("\n\nThe balance for " + testblockchain.MockCoinVaultPubKey[0:40] + "... (Address) is " + balance + "\n\n")
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
