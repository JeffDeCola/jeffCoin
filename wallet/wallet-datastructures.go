// jeffCoin 4. WALLET wallet-datastructures.go

package wallet

// ***********************************************************************************************************************

// walletStruct is your Wallet
type walletStruct struct {
	PrivateKeyHex   string `json:"privateKeyHex"`
	PublicKeyHex    string `json:"publicKeyHex"`
	JeffCoinAddress string `json:"jeffCoinAddress"`
}

// wallet - The Wallet
var wallet = walletStruct{}
