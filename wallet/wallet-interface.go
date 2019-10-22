// jeffCoin wallet-interface.go

package wallet

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	errors "github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// WALLET ************************************************************************************************************

// GenesisWallet - Creates the Wallet
func GenesisWallet() string {

	s := "START: GenesisWallet - Creates the Wallet"
	log.Trace("WALLET:      I/F    " + s)

	theWallet := makeWallet()

	fmt.Printf("\nCongrats, you created your wallet:\n\n")
	js, _ := json.MarshalIndent(theWallet, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END:   GenesisWallet - Creates the Wallet"
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

// COINS ***************************************************************************************************************

// GetAddressBalance - Gets the coin balance for a jeffCoin Address
func GetAddressBalance(nodeIP string, nodeTCPPort string, jeffCoinAddress string) (string, error) {

	s := "START: GetAddressBalance - Gets the coin balance for a jeffCoin Address"
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
	fmt.Fprintf(conn, "SENDADDRESSBALANCE\n")

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

	s = "END:   GetAddressBalance - Gets the coin balance for a jeffCoin Address"
	log.Trace("WALLET:      I/F    " + s)

	return theBalance, nil

}

// TransactionRequest - Request to Transfer Coins to a jeffCoin Address
func TransactionRequest(nodeIP string, nodeTCPPort string, jeffCoinAddress string, value string) (string, error) {

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
	fmt.Fprintf(conn, "TRANSACTIONREQUEST\n")

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
	fmt.Fprintf(conn, jeffCoinAddress+value+"WHATEVER ELSE will be here???????????????????????\n")

	// GET THE IDONTKNOW
	IDONTKNOW, _ := bufio.NewReader(conn).ReadString('\n')
	IDONTKNOW = strings.Trim(IDONTKNOW, "--- ")
	IDONTKNOW = strings.Trim(IDONTKNOW, "\n")
	s = "Message from Network Node: " + IDONTKNOW
	log.Info("WALLET: I/F                " + s)
	if IDONTKNOW == "ERROR" {
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

	return IDONTKNOW, nil

}
