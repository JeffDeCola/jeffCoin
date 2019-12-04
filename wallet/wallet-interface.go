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

// GenesisWallet - Creates the wallet
func GenesisWallet() string {

	s := "START: GenesisWallet - Creates the wallet"
	log.Trace("WALLET:      I/F    " + s)

	theWallet := makeWallet()

	fmt.Printf("\nCongrats, you created your wallet:\n\n")
	js, _ := json.MarshalIndent(theWallet, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END:   GenesisWallet - Creates the wallet"
	log.Trace("WALLET:      I/F    " + s)

	return theWallet.JeffCoinAddress
}

// GetWallet - Gets the wallet
func GetWallet() walletStruct {

	s := "START: GetWallet - Gets the wallet"
	log.Trace("WALLET:      I/F    " + s)

	theWallet := getWallet()

	s = "END:   GetWallet - Gets the wallet"
	log.Trace("WALLET:      I/F    " + s)

	return theWallet

}

// KEYS ******************************************************************************************************************

// EncodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func EncodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	s := "START: EncodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      I/F    " + s)

	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	s = "END:   EncodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      I/F    " + s)

	return privateKeyHex, publicKeyHex

}

// DecodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
func DecodeKeys(privateKeyHex string, publicKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	s := "START: DecodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Trace("WALLET:      I/F    " + s)

	privateKeyRaw, publicKeyRaw := decodeKeys(privateKeyHex, publicKeyHex)

	s = "END:   DecodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Trace("WALLET:      I/F    " + s)

	return privateKeyRaw, publicKeyRaw

}

// JEFFCOINS *************************************************************************************************************

// GetAddressBalance - Gets the jeffCoin balance for a jeffCoin Address
func GetAddressBalance(nodeIP string, nodeTCPPort string, jeffCoinAddress string) (string, error) {

	s := "START: GetAddressBalance - Gets the jeffCoin balance for a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("WALLET: I/F                " + s)
	if message == "ERROR" {
		s = "ERROR: Could not setup connection"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "SEND-ADDRESS-BALANCE\n")

	// GET THE RESPONSE MESSAGE
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("WALLET: I/F                " + s)
	if message == "ERROR" {
		s = "ERROR: Could not setup connection"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND THE ADDRESS
	fmt.Fprintf(conn, jeffCoinAddress+"\n")

	// GET THE BALANCE
	theBalance, _ := bufio.NewReader(conn).ReadString('\n')
	theBalance = strings.Trim(theBalance, "--- ")
	theBalance = strings.Trim(theBalance, "\n")
	s = "Message from Network Node: " + theBalance
	log.Info("WALLET: I/F                " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END:   GetAddressBalance - Gets the jeffCoin balance for a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	return theBalance, nil

}

// TransactionRequest - Request to Transfer jeffCoins to a jeffCoin Address
func TransactionRequest(nodeIP string, nodeTCPPort string, transactionRequestMessageSigned string) (string, error) {

	s := "START: TransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("WALLET: I/F                " + s)
	if message == "ERROR" {
		s = "ERROR: Could not setup connection"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "TRANSACTION-REQUEST\n")

	// GET THE RESPONSE MESSAGE
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("WALLET: I/F                " + s)
	if message == "ERROR" {
		s = "ERROR: Could not setup connection"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND THE TRANSACTION REQUEST
	fmt.Fprintf(conn, transactionRequestMessageSigned+"\n")

	// GET THE STATUS
	status, _ := bufio.NewReader(conn).ReadString('\n')
	status = strings.Trim(status, "--- ")
	status = strings.Trim(status, "\n")
	s = "Message from Network Node: " + status
	log.Info("WALLET: I/F                " + s)
	if status == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Trace("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END:   TransactionRequest - Request to Transfer Coins to a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	return status, nil

}

// SIGNATURE *************************************************************************************************************

// CreateSignature - Create a ECDSA Digital Signature
func CreateSignature(senderPrivateKeyRaw *ecdsa.PrivateKey, plainText string) string {

	s := "START: CreateSignature - Create a ECDSA Digital Signature"
	log.Trace("WALLET:      I/F    " + s)

	signature := createSignature(senderPrivateKeyRaw, plainText)

	s = "END:   CreateSignature - Create a ECDSA Digital Signature"
	log.Trace("WALLET:      I/F    " + s)

	return signature

}
