// jeffCoin wallet.go

package wallet

// walletStruct is your Wallet
type walletStruct struct {
	PrivateKeyHex   string `json:"privateKeyHex"`
	PublicKeyHex    string `json:"publicKeyHex"`
	JeffCoinAddress string `json:"jeffCoinAddress"`
}

// Wallet - The Wallet
var wallet = walletStruct{}
