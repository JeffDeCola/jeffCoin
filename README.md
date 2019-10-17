# jeffCoin

```text
*** THE REPO IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/jeffCoin)](https://goreportcard.com/report/github.com/JeffDeCola/jeffCoin)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/jeffCoin?status.svg)](https://godoc.org/github.com/JeffDeCola/jeffCoin)
[![Maintainability](https://api.codeclimate.com/v1/badges/0c7cf619a01dd65fc06b/maintainability)](https://codeclimate.com/github/JeffDeCola/jeffCoin/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/jeffCoin/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/jeffCoin/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A cryptocurrency (transaction based data) built on a multi node
decentralized P2P open network using a sha256 Proof of Work blockchain
with a REST JSON API and a TCP Server to communicate between
the nodes over IP._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/jeffCoin#overview)
* [1. BLOCKCHAIN](https://github.com/JeffDeCola/jeffCoin#1-blockchain)
* [2. MINER](https://github.com/JeffDeCola/jeffCoin#2-miner)
* [3. ROUTING NODE](https://github.com/JeffDeCola/jeffCoin#3-routing-node)
* [4. WALLET](https://github.com/JeffDeCola/jeffCoin#4-wallet)
* [5. WEBSERVER](https://github.com/JeffDeCola/jeffCoin#5-webserver)
* [RUN](https://github.com/JeffDeCola/jeffCoin#run)
  * [GENESIS NODE](https://github.com/JeffDeCola/jeffCoin#genesis-node)
  * [NEW NODE](https://github.com/JeffDeCola/jeffCoin#new-node)
  * [WEBSERVER AND API](https://github.com/JeffDeCola/jeffCoin#webserver-and-api)
  * [ROUTING NODE](https://github.com/JeffDeCola/jeffCoin#routing-node)

Documentation and reference,

* The Blockchain is built from my
  [single-node-blockchain-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST)
* The Webserver is built from my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
* The Routing Node (TCP Server) is built from my
  [simple-tcp-ip-server](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server)
* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of inspiration from
  [here](https://github.com/nosequeldeebee/blockchain-tutorial)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This code is broken up into four main parts,

* [1. BLOCKCHAIN](https://github.com/JeffDeCola/jeffCoin/tree/master/blockchain)
  The Blockchain code
  * **Guts**
    The guts deal directly with the blockchain
  * **Blockchain Interface**
    The interface to the blockchain
* [2. MINER](https://github.com/JeffDeCola/jeffCoin/tree/master/miner)
  To mine the cryptocurrency
* [3. ROUTING NODE (TCP Server)](https://github.com/JeffDeCola/jeffCoin/tree/master/routingnode)
  To communicate between the P2P nodes (network)
  * **Requests & Handlers**
    The TCP Server
  * **Guts**
    The guts deal directly with the NodeList
  * **RoutingNode Interface**
    The interface to the routingnode
* [4. WALLET](https://github.com/JeffDeCola/jeffCoin/tree/master/wallet)
  To hold the cryptocurrency

I also added a WebServer for a GUI and an API,

* [5. WEBSERVER](https://github.com/JeffDeCola/jeffCoin/tree/master/webserver)
  The API and GUI

jeffCoin will,

* Blockchain using sha256 hash
* Allow multi node with open P2P Architecture
* Maintain a network of Nodes
* View the Entire Blockchain via a web GUI
* View the Entire Nodelist via a web GUI

This illustration may help,

![IMAGE - jeffCoin-overview - IMAGE](docs/pics/jeffCoin-overview.jpg)

## 1. BLOCKCHAIN

The blockchain section is the core of the entire design.

A block in the blockchain made from following struct,

```go
type blockStruct struct {
    Index      int      `json:"index"`
    Timestamp  string   `json:"timestamp"`
    Data       []string `json:"data"`
    Hash       string   `json:"hash"`
    PrevHash   string   `json:"prevhash"`
    Difficulty int      `json:"difficulty"`
    Nonce      string   `json:"nonce"`
}
```

The states of a block are,

* **currentBlock** Receiving transactions and not part of blockchain
* **lockedBlock** Going to be added to the blockchain.
  This one is the one to be mined
* **Part of Chain** These are already in the **blockchain**

Functions in Blockchain Interface,

* BLOCKCHAIN
  * **GenesisBlockchain()** Creates the blockchain
  * **LoadBlockchain()** Receives the blockchain and the currentBlock
    from a Network Node
    * **SENDBLOCKCHAIN** Request
  * **GetBlockchain()** Gets the blockchain
* BLOCK  
  * **GetBlock()** Gets a block (via Index number) from the blockchain
* LOCKED BLOCK
  * **GetLockedBlock** Gets the lockedBlock
* CURRENT BLOCK
  * **GetCurrentBlock** Gets the currentBlock
  * **AddTransactionToCurrentBlock()** Adds a transaction to the currentBlock

The guts will deal directly with the blockchain,

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
  * **getCurrentBlock** Gets the currentBlock
  * **addTransactionToCurrentBlock()** Adds a transaction to the currentBlock
  * **lockCurrentBlock()** Moves the currentBlock to the lockedBlock
    and resets the currentBlock

## 2. MINER

## 3. ROUTING NODE

The Routing Node has two main parts, the nodeList
and the ability to handling Node Requests (TCP Server).

A node in the nodelist is the following struct,

```go
type nodeStruct struct {
    Index     int    `json:"index"`
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
    IP        string `json:"ip"`
    Port      string `json:"port"`
}
```

Functions in RoutingNode Interface,

* NODELIST
  * **GenesisNodeList()** Creates the nodeList
  * **LoadNodeList()** Receives the nodeList from a Network Node
    * **SENDNODELIST** Request
  * **GetNodeList()** Gets the nodeList
* NODE
  * **GetNode()** Gets a Node (via Index number) from the nodeList
  * **AppendNewNode()** Appends a New Node to the nodeList  
* THIS NODE
  * **LoadThisNode()** Loads thisNode
  * **GetThisNode()** Gets thisNode  
  * **AppendThisNode()** Appends thisNode to the nodeList  
  * **BroadcastThisNode()** Broadcasts thisNode to the Network
    * **ADDNEWNODE** Request

The guts will deal directly with the nodeList,

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
  * **checkIfThisNodeinNodeList** - Check if thisNode is already in the nodeList

Handling TCP Server Requests,

* **ADDTRANSACTION (AT)** Add Transaction to CurrentBlock
* **SENDBLOCKCHAIN (SB)** Send Blockchain, LockedBlock &
  CurrentBlock to another Node
* **ADDNEWNODE (NN)** Add Node to nodeList
* **SENDNODELIST (GN)** Send NodeList to another Node
* **EOF**

## 4. WALLET

## 5. WEBSERVER

The user GUI,

[192.168.20.100:1234/](http://localhost:1234/)

The API commands,

* /showBlockchain
* /showBlock/{blockID}
* /showlockedblock
* /showcurrentblock
* /shownodelist
* /shownode/{nodeID}
* /showthisnode

## RUN

If this is you first time running, you need to create the first node.
You only do this once.

### GENESIS NODE

```bash
go run jeffCoin.go \
       -genesis \
       -ip 192.168.20.100 \
       -webport 1234 \
       -tcpport 3334
```

### NEW NODES

Then all other nodes, you do something like this to hook
up to the network.  You need the ip of any working network node,

```bash
go run jeffCoin.go \
       -ip 192.168.20.100 \
       -webport 1235 \
       -tcpport 3335 \
       -netip 192.168.20.100 \
       -netport 3334
```

```bash
go run jeffCoin.go \
       -ip 192.168.20.100 \
       -webport 1236 \
       -tcpport 3336 \
       -netip 192.168.20.100 \
       -netport 3335
```

```bash
go run jeffCoin.go \
       -ip 192.168.20.100 \
       -webport 1237 \
       -tcpport 3337 \
       -netip 192.168.20.100 \
       -netport 3336
```

This will,

* Add your Node to the node list
  * That list will be updated with the network nodes
* Receive the blockchain

### WEBSERVER AND API

The user GUI,

[192.168.20.100:1234/](http://localhost:1234/)

You could also use curl from the command line,

```go
curl 192.168.20.100:1234
```

Show a Particular Block,

[192.168.20.100:1234//showblock/0](http://192.168.20.100:1234/showblock/0)

Show the Chain,

[192.168.20.100:1234//showchain](http://192.168.20.100:1234/showchain)

### ROUTING NODE

Since no security has been setup yet, you can open a connection,

```txt
netcat -q -1 192.168.20.100 3334
```

Now add a transaction,

```txt
ADDTRANSACTION or AT
```

## UPDATE GITHUB WEBPAGE USING CONCOURSE (OPTIONAL)

For fun, I use concourse to update
[jeffCoin GitHub Webpage](https://jeffdecola.github.io/jeffCoin/)
and alert me of the changes via repo status and slack.

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/jeffCoin/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - jeffCoin concourse ci pipeline - IMAGE](docs/pics/jeffCoin-pipeline.jpg)

The `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [readme-github-pages.sh](https://github.com/JeffDeCola/jeffCoin/tree/master/ci/scripts/readme-github-pages.sh).

The concourse `resources types` are,

* `jeffCoin` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
