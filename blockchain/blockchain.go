// jeffCoin blockchain.go

package blockchain

// blockStruct is your block
type blockStruct struct {
	Index      int      `json:"index"`
	Timestamp  string   `json:"timestamp"`
	Data       []string `json:"data"`
	Hash       string   `json:"hash"`
	PrevHash   string   `json:"prevhash"`
	Difficulty int      `json:"difficulty"`
	Nonce      string   `json:"nonce"`
}

// currentBlock - Receiving transactions and not part of chain
var currentBlock = blockStruct{}

// lockedBlock - Going to be added to the chain (No more transactions)
var lockedBlock = blockStruct{}

// BlockchainSlice is my type
type blockchainSlice []blockStruct

// Blockchain is the blockchain
var blockchain = blockchainSlice{}
