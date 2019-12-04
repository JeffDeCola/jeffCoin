# FUNCTION LIST

## 1. BLOCKCHAIN

### 1.1 BLOCKCHAIN

**[BLOCKCHAIN-INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/blockchain-interface.go)**

* BLOCKCHAIN
  * **GenesisBlockchain()** Creates the blockchain
  * **LoadBlockchain()** Receives the blockchain and the currentBlock
    from a Network Node
    * `SEND-BLOCKCHAIN Request`
  * **GetBlockchain()** Gets the blockchain
* BLOCK
  * **GetBlock()** Gets a block (via Index number) from the blockchain
* LOCKED BLOCK
  * **GetLockedBlock()** Gets the lockedBlock
* CURRENT BLOCK
  * **GetCurrentBlock()** Gets the currentBlock
  * **AddTransactionToCurrentBlock()** Adds a transaction to the currentBlock
* JEFFCOINS
  * **GetAddressBalance()** Gets jeffCoin Address balance
* TRANSACTIONS
  * **TransactionRequest()** Request to Transfer jeffCoins to a jeffCoin Address

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/guts.go)**

* BLOCKCHAIN
  * **loadBlockchain()** Loads the entire blockchain
  * **getBlockchain()** Gets the blockchain
  * **replaceBlockchain()** Replaces chain with the longer one
* BLOCK
  * **getBlock()** Gets a block in the blockchain
  * **calculateBlockHash()** Calculates SHA256 hash on a block
  * **isBlockValid()** Checks if block is valid
* LOCKED BLOCK
  * **getLockedBlock** Gets the lockedBlock
  * **appendLockedBlock()** Appends the lockedBlock to the blockchain
* CURRENT BLOCK
  * **loadCurrentBlock()** Loads the currentBlock
  * **resetCurrentBlock()** Resets the currentBlock
  * **getCurrentBlock()** Gets the currentBlock
  * **addTransactionToCurrentBlock()** Adds a transaction to the currentBlock
  * **lockCurrentBlock()** Moves the currentBlock to the lockedBlock
    and resets the currentBlock
* JEFFCOINS
  * **getAddressBalance()** Gets jeffCoin Address balance

### 1.2 TRANSACTIONS

**[TRANSACTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/transactions.go)**

* TRANSACTIONS
  * **transactionRequest()** Request to Transfer jeffCoins to a jeffCoin Address
* SIGNATURE
  * **verifySignature()** - Verify a ECDSA Digital Signature

## 2. MINER

**[MINER-INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/miner/miner-interface.go)**

* MINING
  * **tbd()** tbd

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/miner/guts.go)**

* MINING
  * **tbd()** tbd

## 3. ROUTINGNODE

### 3.1 NODELIST

**[ROUTINGNODE-INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/routingnode-interface.go)**

* NODELIST
  * **GenesisNodeList()** Creates the nodeList
  * **LoadNodeList()** Receives the nodeList from a Network Node
    * `SEND-NODELIST Request`
  * **GetNodeList()** Gets the nodeList
* NODE
  * **GetNode()** Gets a Node (via Index number) from the nodeList
  * **AppendNewNode()** Appends a New Node to the nodeList
* THIS NODE
  * **LoadThisNode()** Loads thisNode
  * **GetThisNode()** Gets thisNode
  * **AppendThisNode()** Appends thisNode to the nodeList
  * **BroadcastThisNode()** Broadcasts thisNode to the Network
    * `BROADCAST-ADD-NEW-NODE Request`

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/guts.go)**

* NODELIST
  * **loadNodeList()** Loads the entire nodeList
  * **getNodeList()** Gets the nodeList
* NODE
  * **getNode()** Gets a node in the nodeList
  * **appendNewNode()** Appends a node to the nodeList
* THIS NODE
  * **loadThisNode()** Loads thisNode
  * **getThisNode()** Gets thisNode
  * **appendThisNode()** Appends thisNode to the nodeList
  * **checkIfThisNodeinNodeList()** - Check if thisNode is already in the nodeList

### 3.2 REQUESTS AND HANDLERS

**[HANDLERS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/handlers.go)**

* FROM BLOCKCHAIN I/F
  * **SEND-BLOCKCHAIN (SBC)**
    Sends the blockchain & currentBlock to another node
* FROM ROUTINGNODE I/F
  * **BROADCAST-ADD-NEW-NODE (BANN)**
    Adds a node to the nodeList
  * **SEND-NODELIST (SNL)**
    Sends the nodeList to another node
  * **BROADCAST-VERIFIED-BLOCK (BVB)**
    A node verified the next block, get block and verify
  * **BROADCAST-CONSENSUS (BC)**
    51% Consensus reached, get block to add to blockchain
  * **BROADCAST-TRANSACTION-REQUEST (BTR)**
    Request from node to transfer jeffCoins to a jeffCoin address
* FROM WALLET I/F
  * **SEND-ADDRESS-BALANCE (SAB)**
    Sends the jeffCoin balance for a jeffCoin Address
  * **TRANSACTION-REQUEST (TR)**
    Request from Wallet to transfer jeffCoins to a jeffCoin address
* EOF
  * **EOF**
    Close Connection

## 4. WALLET

**[WALLET-INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/wallet/wallet-interface.go)**

* WALLET
  * **GenesisWallet()** Creates the wallet
  * **GetWallet()** Gets the wallet
* KEYS
  * **EncodeKeys()** Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
  * **DecodeKeys()** Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
* JEFFCOINS
  * **GetAddressBalance()** Gets the jeffCoin balance for a jeffCoin Address
    * `SEND-ADDRESS-BALANCE Request`
  * **TransactionRequest()** Request to Transfer jeffCoins to a jeffCoin Address
    * `TRANSACTION-REQUEST Request`
* SIGNATURE
  * **CreateSignature()** - Create a ECDSA Digital Signature

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/wallet/guts.go)**

* WALLET
  * **getWallet()** Gets the wallet
  * **makeWallet()** Creates wallet with Keys and jeffCoin address
* KEYS
  * **generateECDSASKeys()** - Generate privateKeyHex and publicKeyHex
  * **encodeKeys()** - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
  * **decodeKeys()** - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
* JEFFCOIN ADDRESS
  * **generateJeffCoinAddress()** - Creates jeffCoinAddress
  * **hashPublicKey()** - Hashes publicKeyHex
  * **checksumKeyHash()** - Checksums verPublicKeyHash
  * **encodeKeyHash()** - Encodes verPublicKeyHash & checkSum
* SIGNATURE
  * **createSignature()** - Create a ECDSA Digital Signature

## 5. WEBSERVER

### 5.1 GUI

* [192.168.20.100:1234](192.168.20.100:1234/)

### 5.2 REST API

**[API COMMANDS](https://github.com/JeffDeCola/jeffCoin/blob/master/webserver/handlers.go)**

* BLOCKCHAIN
  * /showBlockchain
  * /showBlock/{blockID}
  * /showlockedblock
  * /showcurrentblock
* NODELIST
  * /shownodelist
  * /shownode/{nodeID}
  * /showthisnode
* WALLET
  * /showwallet
  * /showjeffcoinaddress
  * /showaddressbalance
  * /transactionrequest/{destinationaddress}/{value}
