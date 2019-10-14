// jeffCoin blockchain.go

package blockchain

// BlockStruct is your block
type BlockStruct struct {
	Index      int      `json:"index"`
	Timestamp  string   `json:"timestamp"`
	Data       []string `json:"data"`
	Hash       string   `json:"hash"`
	PrevHash   string   `json:"prevhash"`
	Difficulty int      `json:"difficulty"`
	Nonce      string   `json:"nonce"`
}

// CurrentBlock - Receiving transactions and not part of chain
var CurrentBlock = BlockStruct{}

// LockedBlock - Going to be added to the chain (No more transactions)
var LockedBlock = BlockStruct{}

// BlockchainSlice is my type
type BlockchainSlice []BlockStruct

// Blockchain is the blockchain
var Blockchain = BlockchainSlice{}
