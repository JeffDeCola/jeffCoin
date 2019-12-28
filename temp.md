# FUNCTION LIST

## 1. BLOCKCHAIN

### 1.1 BLOCKCHAIN

**[BLOCKCHAIN-INTERFACE FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/blockchain-interface.go)**

* BLOCKCHAIN
  * **GetBlockchain()**
    _Gets the blockchain_
  * **GenesisBlockchain()**
    _Creates the blockchain_
  * **RequestBlockchain()**
    _Requests the blockchain and the pendingBlock from a Network Node_
    * `SEND-BLOCKCHAIN Request`
* BLOCK
  * **GetBlock()**
    _Gets a block (via Index number) from the blockchain_
* LOCKED BLOCK
  * **GetLockedBlock()**
    _Gets the lockedBlock_
  * **AppendLockedBlock()**
    _Appends the lockedBlock to the blockchain_
* PENDING BLOCK
  * **GetPendingBlock()**
    _Gets the pendingBlock_
  * **ResetPendingBlock()**
    _Resets the pendingBlock_
  * **AddTransactionToPendingBlock()**
    _Adds a transaction to the pendingBlock and makes change_
  * **LockPendingBlock()**
    _Moves the pendingBlock to the lockedBlock_
* JEFFCOINS
  * **GetAddressBalance()**
    _Gets the jeffCoin Address balance_
* TRANSACTIONS
  * **ProcessTxRequestMessage()**
    _Request to transfer jeffCoins to a jeffCoin Address_

**[GUTS FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/guts.go)**

* BLOCKCHAIN
  * **getBlockchain()**
    _Gets the blockchain_
  * **loadBlockchain()**
    _Loads the entire blockchain_
  * **replaceBlockchain()**
    _Replaces blockchain with the longer one_
* BLOCK
  * **getBlock()**
    _Gets a block in the blockchain_
  * **calculateBlockHash()**
    _Calculates SHA256 hash on a block_
  * **isBlockValid()**
    _Checks if block is valid_
* LOCKED BLOCK
  * **getLockedBlock()**
    _Gets the lockedBlock_
  * **appendLockedBlock()**
    _Appends the lockedBlock to the blockchain_
* PENDING BLOCK
  * **getPendingBlock()**
    _Gets the pendingBlock_
  * **loadPendingBlock()**
    _Loads the pendingBlock_
  * **resetPendingBlock()**
    _Resets the pendingBlock_
  * **addTransactionToPendingBlock()**
    _Adds a transaction to the pendingBlock and makes change_
  * **lockPendingBlock()**
    _Moves the pendingBlock to the lockedBlock_
* JEFFCOINS
  * **getAddressBalance()**
    _Gets the jeffCoin Address balance_

### 1.2 TRANSACTIONS

**[TRANSACTION FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/transactions.go)**

* TRANSACTIONS
  * **processTxRequestMessage()**
    _Request to transfer jeffCoins to a jeffCoin Address_
* SIGNATURE
  * **verifySignature()**
    _Verifies a ECDSA Digital Signature_
* UNSPENT OUTPUTS
  * **pickUnspentOutputs()**
    _Pick the Unspent Outputs to use and provide change_

## 2. MINER

**[MINER-INTERFACE FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/miner/miner-interface.go)**

* MINING
  * **tbd()**
    _tbd_

**[GUTS FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/miner/guts.go)**

* MINING
  * **tbd()**
    _tbd_

## 3. ROUTINGNODE

### 3.1 NODELIST

**[ROUTINGNODE-INTERFACE FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/routingnode-interface.go)**

* NODELIST
  * **GetNodeList()**
    _Gets the nodeList_
  * **GenesisNodeList()**
    _Creates the nodeList_
  * **RequestsNodeList()**
    _Requests the nodeList from a Network Node_
    * `SEND-NODELIST Request`
* NODE
  * **GetNode()**
    _Gets a Node (via Index number) from the nodeList_
  * **AppendNewNode()**
    _Appends a new Node to the nodeList_
* THIS NODE
  * **GetThisNode()**
    _Gets thisNode_
  * **LoadThisNode()**
    _Loads thisNode_
  * **AppendThisNode()**
    _Appends thisNode to the nodeList_
  * **BroadcastThisNode()**
    _Broadcasts thisNode to the Network_
    * `BROADCAST-ADD-NEW-NODE Request`

**[GUTS FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/guts.go)**

* NODELIST
  * **getNodeList()**
    _Gets the nodeList_
  * **loadNodeList()**
    _Loads the entire nodeList_
* NODE
  * **getNode()**
    _Gets a Node in the nodeList_
  * **appendNewNode()**
    _Appends a new Node to the nodeList_
* THIS NODE
  * **getThisNode()**
    _Gets thisNode_
  * **loadThisNode()**
    _Loads thisNode_
  * **appendThisNode()**
    _Appends thisNode to the nodeList_
  * **checkIfThisNodeinNodeList()**
    _Check if thisNode is already in the nodeList_

### 3.2 TCP REQUESTS & HANDLERS

**[HANDLERS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/handlers.go)**

* FROM BLOCKCHAIN I/F
  * **handleSendBlockchain()**
    _SEND-BLOCKCHAIN (**SBC**)- Sends the blockchain and pendingBlock to another Node_
* FROM ROUTINGNODE I/F
  * **handleBroadcastAddNewNode()**
    _BROADCAST-ADD-NEW-NODE (**BANN**) - Adds a Node to the nodeList_
  * **handleSendNodeList()**
    _SEND-NODELIST (SNL) - Sends the nodeList to another Node_
  * **handleBroadcastVerifiedBlock()**
    _BROADCAST-VERIFIED-BLOCK (**BVB**) - A Node verified the next block,
    get block and verify_
  * **handleBroadcastConsensus()**
    _BROADCAST-CONSENSUS (**BC**) - 51% Consensus reached, get block to add to blockchain_
  * **handleBroadcastTransactionRequest()**
    _BROADCAST-TRANSACTION-REQUEST (**BTR**) - Request from a Node
    to transfer jeffCoins to a jeffCoin Address_
* FROM WALLET I/F
  * **handleSendAddressBalance()**
    _SEND-ADDRESS-BALANCE (**SAB**) - Sends the jeffCoin balance
    for a jeffCoin Address_
  * **handleTransactionRequest()**
    _TRANSACTION-REQUEST (**TR**) - Request from Wallet to transfer jeffCoins
    to a jeffCoin Address_
* EOF
  * **EOF**
    _Close Connection_

## 4. WALLET

**[WALLET-INTERFACE FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/wallet/wallet-interface.go)**

* WALLET
  * **GetWallet()**
    _Gets the wallet_
  * **GenesisWallet()**
    _Creates the wallet and writes to file (Keys and jeffCoin Address)_
  * **ReadWalletFile()**
    _Reads the wallet from a file_
* KEYS
  * **EncodeKeys()**
    _Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex_
  * **DecodeKeys()**
    _Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw_
* JEFFCOINS
  * **RequestAddressBalance()**
    _Requests the jeffCoin balance for a jeffCoin Address_
    * `SEND-ADDRESS-BALANCE Request`
  * **TransactionRequest()**
    _Request to transfer jeffCoins to a jeffCoin Address_
    * `TRANSACTION-REQUEST Request`
* SIGNATURE
  * **CreateSignature()**
    _Creates a ECDSA Digital Signature_

**[GUTS FUNCTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/wallet/guts.go)**

* WALLET
  * **getWallet()**
    _Gets the wallet_
  * **makeWallet()**
    _Creates the wallet and writes to file (Keys and jeffCoin Address)_
  * **readWalletFile()**
    _Reads the wallet from a file_
* KEYS
  * **generateECDSASKeys()**
    _Generate privateKeyHex and publicKeyHex_
  * **encodeKeys()**
    _Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex_
  * **decodeKeys()**
    _Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw_
* JEFFCOIN ADDRESS
  * **generateJeffCoinAddress()**
    _Creates a jeffCoin Address_
  * **hashPublicKey()**
    _Hashes publicKeyHex_
  * **checksumKeyHash()**
    _Checksums verPublicKeyHash_
  * **encodeKeyHash()**
    _Encodes verPublicKeyHash & checkSum_
* SIGNATURE
  * **createSignature()**
    _Creates a ECDSA Digital Signature_

## 5. WEBSERVER

### 5.2 REST API

**[API COMMANDS](https://github.com/JeffDeCola/jeffCoin/blob/master/webserver/handlers.go)**

* BLOCKCHAIN
  * **/showBlockchain**
  * **/showBlock/{blockID}**
* LOCKED BLOCK  
  * **/showlockedblock**
  * **/appendlockedblock**
* PENDING BLOCK
  * **/showpendingblock**
  * **/resetpendingblock**
  * **/lockpendingblock**
* NODELIST
  * **/shownodelist**
  * **/shownode/{nodeID}**
  * **/showthisnode**
* WALLET (THIS NODE)
  * **/showwallet**
  * **/showjeffcoinaddress**
  * **/showbalance**
  * **/transactionrequest/{destinationaddress1,destinationaddress2,...}/{value1,value2,...}**
* WALLET (OTHER)
  * **/showaddressbalance/{jeffcoinaddress}**
