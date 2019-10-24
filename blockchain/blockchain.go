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
	ID      int64            `json:"ID"`
	Inputs  []txInputStruct  `json:"inputs"`
	Outputs []txOutputStruct `json:"outputs"`
}

// transaction is the transaction
var transactions = transactionStruct{}

// txInput
type txInputStruct struct {
	TXID          int64  `json:"txID"`
	ReferenceTXID int64  `json:"referenceTXID"`
	Signature     string `json:"signature"`
}

// txOutput - This is where the money is stored
type txOutputStruct struct {
	JeffCoinAddress string `json:"jeffCoinAddress"`
	Value           int64  `json:"value"`
}

// TRANSACTION REQUESTS *************************************************************************************

// transactionRequestMessageStruct is your transaction request
type transactionRequestMessageStruct struct {
	RequestMessage requestMessageStruct `json:"requestMessage"`
	Signature      string               `json:"signature"`
}

// requestMessageStruct
type requestMessageStruct struct {
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	Value              int64  `json:"value"`
}
