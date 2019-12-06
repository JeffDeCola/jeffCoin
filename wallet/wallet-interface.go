// jeffCoin 4. WALLET wallet-interface.go

package wallet

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	errors "github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// WALLET ****************************************************************************************************************

// GetWallet - Gets the wallet
func GetWallet() walletStruct {

	s := "START  GetWallet() - Gets the wallet"
	log.Trace("WALLET:      I/F    " + s)

	theWallet := getWallet()

	s = "END    GetWallet() - Gets the wallet"
	log.Trace("WALLET:      I/F    " + s)

	return theWallet

}

// GenesisWallet - Creates the wallet (Keys and jeffCoin Address)
func GenesisWallet() string {

	s := "START  GenesisWallet() - Creates the wallet (Keys and jeffCoin Address)"
	log.Trace("WALLET:      I/F    " + s)

	theWallet := makeWallet()

	fmt.Printf("\nCongrats, you created your wallet:\n\n")
	js, _ := json.MarshalIndent(theWallet, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END    GenesisWallet() - Creates the wallet (Keys and jeffCoin Address)"
	log.Trace("WALLET:      I/F    " + s)

	return theWallet.JeffCoinAddress
}

// KEYS ******************************************************************************************************************

// EncodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func EncodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	s := "START  EncodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      I/F    " + s)

	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	s = "END    EncodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      I/F    " + s)

	return privateKeyHex, publicKeyHex

}

// DecodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
func DecodeKeys(privateKeyHex string, publicKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	s := "START  DecodeKeys() - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Trace("WALLET:      I/F    " + s)

	privateKeyRaw, publicKeyRaw := decodeKeys(privateKeyHex, publicKeyHex)

	s = "END    DecodeKeys() - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Trace("WALLET:      I/F    " + s)

	return privateKeyRaw, publicKeyRaw

}

// JEFFCOINS *************************************************************************************************************

// RequestAddressBalance - Requests the jeffCoin balance for a jeffCoin Address
func RequestAddressBalance(nodeIP string, nodeTCPPort string, jeffCoinAddress string) (string, error) {

	s := "START  RequestAddressBalance() - Requests the jeffCoin balance for a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	// SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("WALLET:      I/F           " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)
	s = "-conn   TCP Connection on " + nodeIP + ":" + nodeTCPPort
	log.Info("WALLET:      I/F   " + s)
	conn, err := net.Dial("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("WALLET:      I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND-ADDRESS-BALANCE
	s = "-req    - SEND-ADDRESS-BALANCE"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "SEND-ADDRESS-BALANCE\n")

	// GET THE RESPONSE (ASKING TO SEND jeffCoin Address)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: ASKING TO SEND jeffCoin Address"
		log.Trace("WALLET:      I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND jeffCoinAddress
	s = "-send   SEND jeffCoinAddress " + jeffCoinAddress
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, jeffCoinAddress+"\n")

	// GET BALANCE
	theBalance, _ := bufio.NewReader(conn).ReadString('\n')
	theBalance = strings.Trim(theBalance, "--- ")
	theBalance = strings.Trim(theBalance, "\n")
	s = "-rcv    Message from Network Node: " + theBalance
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get balance"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("WALLET:      I/F            " + s)
		return "error", errors.New(s)
	}

	// EOF (CLOSE CONNECTION)
	s = "-req    - EOF (CLOSE CONNECTION)"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)
	s = "CLIENT - Closed a connection"
	log.Info("WALLET:      I/F           " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)

	s = "END    RequestAddressBalance() - Requests the jeffCoin balance for a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	return theBalance, nil

}

// TransactionRequest - Request to transfer jeffCoins to a jeffCoin Address
func TransactionRequest(nodeIP string, nodeTCPPort string, transactionRequestMessageSigned string) (string, error) {

	s := "START  TransactionRequest() - Request to transfer Coins to a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	// SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("WALLET:      I/F           " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)
	s = "-conn   TCP Connection on " + nodeIP + ":" + nodeTCPPort
	log.Info("WALLET:      I/F   " + s)
	conn, err := net.Dial("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// TRANSACTION-REQUEST
	s = "-req    - TRANSACTION-REQUEST"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "TRANSACTION-REQUEST\n")

	// GET THE RESPONSE MESSAGE (???????????????????????)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: ?????????????????????????????"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND transactionRequestMessageSigned
	s = "-send   SEND transactionRequestMessageSigned " + transactionRequestMessageSigned
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, transactionRequestMessageSigned+"\n")

	// GET THE STATUS
	status, _ := bufio.NewReader(conn).ReadString('\n')
	status = strings.Trim(status, "--- ")
	status = strings.Trim(status, "\n")
	s = "-rcv    Message from Network Node: " + status
	log.Info("WALLET:      I/F   " + s)
	if status == "ERROR" {
		s = "ERROR: Could not get the status from node"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("WALLET:      I/F            " + s)
		return "error", errors.New(s)
	}

	// EOF (CLOSE CONNECTION)
	s = "-req    - EOF (CLOSE CONNECTION)"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)
	s = "CLIENT - Closed a connection"
	log.Info("WALLET:      I/F           " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F           " + s)

	s = "END    TransactionRequest() - Request to transfer Coins to a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	return status, nil

}

// SIGNATURE *************************************************************************************************************

// CreateSignature - Creates a ECDSA Digital Signature
func CreateSignature(senderPrivateKeyRaw *ecdsa.PrivateKey, plainText string) string {

	s := "START  CreateSignature() - Creates a ECDSA Digital Signature"
	log.Trace("WALLET:      I/F    " + s)

	signature := createSignature(senderPrivateKeyRaw, plainText)

	s = "END    CreateSignature() - Creates a ECDSA Digital Signature"
	log.Trace("WALLET:      I/F    " + s)

	return signature

}
