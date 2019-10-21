// jeffCoin blockchain.go

package blockchain

// BLOCK & BLOCKCHAIN **************************************************************************************

// blockStruct is your block
type blockStruct struct {
	Index        int                 `json:"index"`
	Timestamp    string              `json:"timestamp"`
	Transactions []transactionStruct `json:"transactions"`
	Hash         string              `json:"hash"`
	PrevHash     string              `json:"prevhash"`
	Difficulty   int                 `json:"difficulty"`
	Nonce        string              `json:"nonce"`
}

// currentBlock - Receiving transactions and not part of chain
var currentBlock = blockStruct{}

// lockedBlock - Going to be added to the chain (No more transactions)
var lockedBlock = blockStruct{}

// BlockchainSlice is my type
type blockchainSlice []blockStruct

// Blockchain is the blockchain
var blockchain = blockchainSlice{}

// TRANSACTIONS **************************************************************************************

// transactionStruct is your transaction - To be place in block
type transactionStruct struct {
	ID    string     `json:"ID"`
	Inputs  []txInput  `json:"inputs"`
	Outputs []txOutput `json:"outputs"`
}

// transaction is the transaction
var transactions = transactionStruct{}

// txInput
type txInput struct {
    TXID           string `json:"txID"`
	ReferenceTXID string `json:"referenceTXID"`
	Signature     string `json:"signature"`
}

// txOutput - This is where the money is stored
type txOutput struct {
	JeffCoinAddress string `json:"jeffCoinAddress"`
	Value           int64  `json:"value"`
}
