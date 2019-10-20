// jeffCoin wallet-interface.go

package wallet

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// WALLET ************************************************************************************************************

// GenesisWallet - Creates the Wallet
func GenesisWallet() {

	s := "START: GenesisWallet - Creates the Wallet"
	log.Trace("WALLET:      I/F    " + s)

	theWallet := makeWallet()

	fmt.Printf("\nCongrats, you created your wallet:\n\n")
	js, _ := json.MarshalIndent(theWallet, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END:   GenesisWallet - Creates the Wallet"
	log.Trace("WALLET:      I/F    " + s)

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
