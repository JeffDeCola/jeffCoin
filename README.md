# multi-node-blockchain-with-REST-and-tcp-ip example

_A cryptocurrency (transaction based data) built on a multi node P2P open network
using a sha256 blockchain with a REST JSON API and a tcp server to communicate
between the nodes over ip._

Table of Contents,

* tbd

Documentation and reference,

* The blockchain is built from my
  [single-node-blockchain-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST)
* The webserver is built from my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
* The tcp server is built from my
  [simple-tcp-ip-server](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server)
* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of inspiration from
  [here](https://github.com/nosequeldeebee/blockchain-tutorial)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This code is broken up into four separate parts,

* [BLOCKCHAIN](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/blockchain)
  The Blockchain code
  * **Guts**
    The guts deal directly with the blockchain
  * **Blockchain Interface**
    The interface to the blockchain
* [MINER](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/miner)
  To mine the cryptocurrency
* [ROUTING NODE (TCP Server)](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/routing-node)
  To communicate between the P2P nodes
* [WALLET](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/wallet)
  To hold the cryptocurrency

I also added a webserver for a GUI,

* [WEBSERVER](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/multi-node-blockchain-with-REST-and-tcp-ip/webserver)
  The API and GUI

This example is pretty ambitious and will,

* Allow multi node with P2P Architecture
* Create a Blockchain
* Hash a Block
* View the Entire Blockchain via a web GUI
* View a specific Block via a REST API
* Add a Block (with your data) to the chain via a REST API

This illustration may help,

![IMAGE - multi-node-blockchain-with-REST-and-tcp-ip - IMAGE](https://github.com/JeffDeCola/my-go-examples/blob/master/docs/pics/multi-node-blockchain-with-REST-and-tcp-ip.jpg)

## BLOCKCHAIN

```go
type BlockStruct struct {
    Index      int      `json:"index"`
    Timestamp  string   `json:"timestamp"`
    Data       []string `json:"data"`
    Hash       string   `json:"hash"`
    PrevHash   string   `json:"prevhash"`
    Difficulty int      `json:"difficulty"`
    Nonce      string   `json:"nonce"`
}
```

## MINER

## ROUTING NODE (TCP Server)

## WALLET

## WEBSERVER

## RUN

```bash
go run multi-node-blockchain-with-REST-and-tcp-ip.go
```

### WEBSERVER AND API

#### GET (View the entire Blockchain)

Then you can goto the webpage to see your first block,

[localhost:1234/](http://localhost:1234/)

You could also use curl from the command line,

```go
curl localhost:1234
```

#### GET (Show a Particular Block)

[localhost:1234//showblock/0](http://localhost:1234/showblock/0)

```go
curl localhost:1234/showblock/0
```

#### POST (Add a Block)

```bash
curl -H "Content-Type: application/json" \
     -X POST \
     -d '{"data":"Add this data for new block"}' \
     localhost:1234/addblock
```

Check,

[localhost:1234/](http://localhost:1234/)

### TCP SERVER

Open a connection,

```txt
netcat -q -1 localhost 3333
```

Now add a block or transaction,

```txt
ADDBLOCK or AB
ADDTRANSACTION or AT
```
