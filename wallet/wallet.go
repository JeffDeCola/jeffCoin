// jeffCoin wallet.go

package wallet

// walletStruct is your Wallet
type walletStruct struct {
	PrivateKey []byte `json:"privateKey"`
	PublicKey  []byte `json:"publicKey"`
	Address    string `json:"address"`
}

// Wallet - The Wallet
var wallet = walletStruct{}
