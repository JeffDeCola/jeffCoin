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
	log.Debug("WALLET:      I/F      " + s)

	theWallet := getWallet()

	s = "END    GetWallet() - Gets the wallet"
	log.Debug("WALLET:      I/F      " + s)

	return theWallet

}

// GenesisWallet - Creates the wallet and writes to file (Keys and jeffCoin Address)
func GenesisWallet(nodeName string, password string) string {

	s := "START  GenesisWallet() - Creates the wallet and writes to file (Keys and jeffCoin Address)"
	log.Debug("WALLET:      I/F      " + s)

	theWallet := makeWallet(nodeName, password)

	s = "Congrats, you created your wallet (-loglevel trace to display)"
	log.Info("ROUTINGNODE: I/F             " + s)
	js, _ := json.MarshalIndent(theWallet, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	s = "END    GenesisWallet() - Creates the wallet and writes to file (Keys and jeffCoin Address)"
	log.Debug("WALLET:      I/F      " + s)

	return theWallet.JeffCoinAddress
}

// ReadWalletFile - Reads the wallet from a file and puts in struct
func ReadWalletFile(nodeName string, password string) string {

	s := "START  ReadWalletFile() - Reads the wallet from a file and puts in struct"
	log.Debug("WALLET:      I/F      " + s)

	theWallet := readWalletFile(nodeName, password)

	s = "Congrats, you loaded your old wallet from a file (-loglevel trace to display)"
	log.Info("ROUTINGNODE: I/F             " + s)
	js, _ := json.MarshalIndent(theWallet, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	s = "END    ReadWalletFile() - Reads the wallet from a file and puts in struct"
	log.Debug("WALLET:      I/F      " + s)

	return theWallet.JeffCoinAddress
}

// KEYS ******************************************************************************************************************

// EncodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func EncodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	s := "START  EncodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Debug("WALLET:      I/F      " + s)

	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	s = "END    EncodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Debug("WALLET:      I/F      " + s)

	return privateKeyHex, publicKeyHex

}

// DecodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
func DecodeKeys(privateKeyHex string, publicKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	s := "START  DecodeKeys() - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Debug("WALLET:      I/F      " + s)

	privateKeyRaw, publicKeyRaw := decodeKeys(privateKeyHex, publicKeyHex)

	s = "END    DecodeKeys() - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Debug("WALLET:      I/F      " + s)

	return privateKeyRaw, publicKeyRaw

}

// JEFFCOINS *************************************************************************************************************

// RequestAddressBalance - Requests the jeffCoin balance for a jeffCoin Address
func RequestAddressBalance(IP string, TCPPort string, jeffCoinAddress string) (string, error) {

	s := "START  RequestAddressBalance() - Requests the jeffCoin balance for a jeffCoin Address"
	log.Debug("WALLET:      I/F      " + s)

	//  CONN - SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("WALLET:      I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)
	s = "-C conn   TCP Connection on " + IP + ":" + TCPPort
	log.Info("WALLET:      I/F   " + s)
	conn, err := net.Dial("tcp", IP+":"+TCPPort)
	checkErr(err)

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Debug("WALLET:      I/F                   " + s)
		return "error", errors.New(s)
	}

	// REQ - SEND-ADDRESS-BALANCE
	s = "-C req    - SEND-ADDRESS-BALANCE"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "SEND-ADDRESS-BALANCE\n")

	// RCV - GET THE RESPONSE (ASKING TO SEND jeffCoin Address)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: ASKING TO SEND jeffCoin Address"
		log.Debug("WALLET:      I/F                   " + s)
		return "error", errors.New(s)
	}

	// SEND - jeffCoinAddress
	s = "-C send   SEND jeffCoinAddress " + jeffCoinAddress
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, jeffCoinAddress+"\n")

	// RCV - GET BALANCE
	theBalance, _ := bufio.NewReader(conn).ReadString('\n')
	theBalance = strings.Trim(theBalance, "--- ")
	theBalance = strings.Trim(theBalance, "\n")
	s = "-C rcv    Message from Network Node: " + theBalance
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get balance"
		log.Debug("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND - THANK YOU
	s = "-C send   - Thank you"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "Thank You\n")

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Debug("WALLET:      I/F              " + s)
		return "error", errors.New(s)
	}

	// REQ - EOF (CLOSE CONNECTION)
	s = "-C req    - EOF (CLOSE CONNECTION)"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)
	s = "CLIENT - Closed a connection"
	log.Info("WALLET:      I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)

	s = "END    RequestAddressBalance() - Requests the jeffCoin balance for a jeffCoin Address"
	log.Debug("WALLET:      I/F      " + s)

	return theBalance, nil

}

// TransactionRequest - Request to transfer jeffCoins to a jeffCoin Address
func TransactionRequest(nodeIP string, nodeTCPPort string, transactionRequestMessageSigned string) (string, error) {

	s := "START  TransactionRequest() - Request to transfer Coins to a jeffCoin Address"
	log.Debug("WALLET:      I/F      " + s)

	//  CONN - SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("WALLET:      I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)
	s = "-C conn   TCP Connection on " + nodeIP + ":" + nodeTCPPort
	log.Info("WALLET:      I/F   " + s)
	conn, err := net.Dial("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Debug("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// REQ - TRANSACTION-REQUEST
	s = "-C req    - TRANSACTION-REQUEST"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "TRANSACTION-REQUEST\n")

	// RCV - GET THE RESPONSE MESSAGE (???????????????????????)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: ?????????????????????????????"
		log.Debug("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND - transactionRequestMessageSigned
	s = "-C send   SEND transactionRequestMessageSigned (-loglevel trace to display)"
	log.Info("WALLET:      I/F   " + s)
	log.Trace("\n\n" + transactionRequestMessageSigned + "\n\n")
	fmt.Fprintf(conn, transactionRequestMessageSigned+"\n")

	// RCV - GET THE STATUS
	status, _ := bufio.NewReader(conn).ReadString('\n')
	status = strings.Trim(status, "--- ")
	status = strings.Trim(status, "\n")
	s = "-C rcv    Message from Network Node: " + status
	log.Info("WALLET:      I/F   " + s)

	if status == "ERROR" {
		s = "ERROR: Could not get the status from node"
		log.Debug("WALLET: I/F                 " + s)
		return "error", errors.New(s)
	}

	// SEND - THANK YOU
	s = "-C send   - Thank you"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "Thank You\n")

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("WALLET:      I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Debug("WALLET:      I/F              " + s)
		return "error", errors.New(s)
	}

	// REQ - EOF (CLOSE CONNECTION)
	s = "-C req    - EOF (CLOSE CONNECTION)"
	log.Info("WALLET:      I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)
	s = "CLIENT - Closed a connection"
	log.Info("WALLET:      I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("WALLET:      I/F             " + s)

	s = "END    TransactionRequest() - Request to transfer Coins to a jeffCoin Address"
	log.Debug("WALLET:      I/F      " + s)

	return status, nil

}

// SIGNATURE *************************************************************************************************************

// CreateSignature - Creates a ECDSA Digital Signature
func CreateSignature(senderPrivateKeyHex string, plainText string) string {

	s := "START  CreateSignature() - Creates a ECDSA Digital Signature"
	log.Debug("WALLET:      I/F      " + s)

	signature := createSignature(senderPrivateKeyHex, plainText)

	s = "END    CreateSignature() - Creates a ECDSA Digital Signature"
	log.Debug("WALLET:      I/F      " + s)

	return signature

}

// ENCRYPT/DECRYPT TEXT **************************************************************************************************

// EncryptAES - AES-256 GCM (Galois/Counter Mode) mode encryption
func EncryptAES(keyByte []byte, plaintext string, additionalData string) string {

	s := "START  EncryptAES() - AES-256 GCM (Galois/Counter Mode) mode encryption"
	log.Debug("WALLET:      I/F      " + s)

	// ENCRYPT
	cipherText := encryptAES(keyByte, plaintext, additionalData)

	s = "END    EncryptAES() - AES-256 GCM (Galois/Counter Mode) mode encryption"
	log.Debug("WALLET:      I/F      " + s)

	// RETURN HEX
	return cipherText

}

// DecryptAES - AES-256 GCM (Galois/Counter Mode) mode decryption
func DecryptAES(keyByte []byte, cipherText string, additionalData string) string {

	s := "START  DecryptAES() - AES-256 GCM (Galois/Counter Mode) mode decryption"
	log.Debug("WALLET:      I/F      " + s)

	plainText := decryptAES(keyByte, cipherText, additionalData)

	s = "END    DecryptAES() - AES-256 GCM (Galois/Counter Mode) mode decryption"
	log.Debug("WALLET:      I/F      " + s)

	// RETURN STRING
	return plainText

}
