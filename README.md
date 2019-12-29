# jeffCoin

```text
*** THE REPO IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/jeffCoin)](https://goreportcard.com/report/github.com/JeffDeCola/jeffCoin)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/jeffCoin?status.svg)](https://godoc.org/github.com/JeffDeCola/jeffCoin)
[![Maintainability](https://api.codeclimate.com/v1/badges/0c7cf619a01dd65fc06b/maintainability)](https://codeclimate.com/github/JeffDeCola/jeffCoin/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/jeffCoin/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/jeffCoin/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A cryptocurrency (transaction based data) built on decentralized
multi-node P2P open Network using a sha256 Proof of Work (PoW) blockchain
with a REST JSON API and a TCP Server to communicate between
the Nodes over IP._

Or more simply, **a distributed decentralized public ledger.**

To dive right in, head down to [RUN](https://github.com/JeffDeCola/jeffCoin#run).

Table of Contents,

* [IMPORTANT](https://github.com/JeffDeCola/jeffCoin#important)
* [PREREQUISITES](https://github.com/JeffDeCola/jeffCoin#prerequisites)
* [OVERVIEW](https://github.com/JeffDeCola/jeffCoin#overview)
* [SOFTWARE ARCHITECTURE](https://github.com/JeffDeCola/jeffCoin#software-architecture)
* [RUN](https://github.com/JeffDeCola/jeffCoin#run)
  * [GENESIS NODE](https://github.com/JeffDeCola/jeffCoin#genesis-node)
  * [ADDING NEW NODES](https://github.com/JeffDeCola/jeffCoin#adding-new-nodes)
  * [WEBSERVER & REST API](https://github.com/JeffDeCola/jeffCoin#webserver--rest-api)
  * [SWITCHES (REFERENCE)](https://github.com/JeffDeCola/jeffCoin#switches-reference)
  * [JUST WALLET (OPTIONAL)](https://github.com/JeffDeCola/jeffCoin#just-wallet-optional)
  * [CONNECT USING TCP (OPTIONAL)](https://github.com/JeffDeCola/jeffCoin#connect-using-tcp-optional)
  * [TEST MOCK TRANSACTIONS (OPTIONAL)](https://github.com/JeffDeCola/jeffCoin#test-mock-transactions-optional)
  * [RUN ON GOOGLE COMPUTE ENGINE (GCE) (OPTIONAL)](https://github.com/JeffDeCola/jeffCoin#run-on-google-compute-engine-gce-optional)

The explanation of the software architecture is
[architecture.md](https://github.com/JeffDeCola/jeffCoin/blob/master/architecture.md).

This project was built from some of my other projects,

* The **BLOCKCHAIN** is built from my
  [single-node-blockchain-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST)
  * The **BLOCKCHAIN TRANSACTIONS** is built from my
    [bitcoin-ledger](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger)
  * The ecdsa signature verification from my
    [ecdsa-digital-signature](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/asymmetric-cryptography/ecdsa-digital-signature)
* The **ROUTINGNODE** (TCP Server) is built from my
  [simple-tcp-ip-server](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server)
* The **WALLET** for generating keys and creating the jeffCoin address
  is built from my
  [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/create-bitcoin-address-from-ecdsa-publickey)
* The **WEBSERVER** (GUI & REST JSON API) is built from my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
* Other projects I used are my
  [errors](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/errors),
  [logrus](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/logrus),
  and
  [flag](https://github.com/JeffDeCola/my-go-examples/tree/master/packages/flag)
projects.

Documentation and reference,

* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of inspiration
  [here](https://github.com/nosequeldeebee/blockchain-tutorial)
* jeffCoin
  [architecture.md](https://github.com/JeffDeCola/jeffCoin/blob/master/architecture.md).

[GitHub Webpage](https://jeffdecola.github.io/jeffCoin/)

## IMPORTANT

Your private keys are kept in `/wallet`.  The .gitignore
file does ignore them, but just be aware were they live.
There are a few mock keys there used for testing.

## PREREQUISITES

```bash
go get -v -u github.com/btcsuite/btcutil/base58
go get -v -u golang.org/x/crypto/ripemd160
go get -u -v github.com/gorilla/mux
go get -u -v github.com/sirupsen/logrus
go get -u -v github.com/pkg/errors
```

## OVERVIEW

`jeffCoin` (JEFC) is my interpretation of a transaction based (ledger) using a blockchain.
This is a work in progress I feel can be used as a foundation to
build bigger and better things.

Coins (a.k.a jeffCoins) are minted as follows,

* A grand total of **1,000,000 jeffCoins**
* The blockchain will not store jeffCoins but **addies** which are
  1/1000 of a jeffCoin (.001 JEFF)
* The founders wallet will start with **100,000 jeffCoins (100,000,000 addies)**
  (10% of all jeffCoins)
* Rewards **1 jeffCoin (1000 addies) every 10 minutes**
  _(144 jeffCoins/day or 52,560 jeffCoins/year)_
* Will take **17.12 years to mint all the jeffCoins**
  _(900,000/52,560 = 17.12)_

jeffCoin uses the following technology,

* Written in golang
* Implements a blockchain using a sha256 hash
* A decentralized multi-node P2P architecture maintaining a Network of Nodes
* A Webserver with both a GUI and a REST API
* A TCP Server for inter-node communication
* ECDSA Private & Public Key generation
* Creates a jeffCoin Address from the ECDSA Public Key _(Just like bitcoin)_
* ECDSA Digital Signature Verification
* Mining uses Proof of Work (PoW)
* Transaction as stored using the unspent transaction output model

What jeffCoin does not have,

* No database, so if the entire Network dies, the chain dies
* Rigorous testing of all corner cases

The following illustration shows the network of nodes,

![IMAGE - the-jeffcoin-network - IMAGE](docs/pics/the-jeffcoin-network.jpg)

## SOFTWARE ARCHITECTURE

This readme got too big so I moved the software explanation to
[architecture.md](https://github.com/JeffDeCola/jeffCoin/blob/master/architecture.md).

## RUN

If this is you first time running, you need to create the first Node (Genesis Node).
You only do this once. You can set the log level (info, debug, trace)
to cut down on the amount of logging.

### GENESIS NODE

```bash
go run jeffCoin.go \
       -loglevel debug \
       -genesis \
       -nodename Founders \
       -ip 127.0.0.1 \
       -httpport 2000 \
       -tcpport 3000
```

This will created the first Node (the Founders node) in the Network.
It will also create a wallet and save the credentials in `/wallet`.
But having one node is boring so create more.

### ADDING NEW NODES

To hook up to the Network.  You need the IP of any
working Network Node. If you have the above running
on `127.0.0.1:3000`, adding a second Node
"Jeff" in your network could look like,

```bash
go run jeffCoin.go \
       -loglevel debug \
       -nodename Jeff \
       -ip 127.0.0.1 \
       -httpport 2001 \
       -tcpport 3001 \
       -netip 127.0.0.1 \
       -netport 3000
```

Might as well add a third Node,

```bash
go run jeffCoin.go \
       -loglevel debug \
       -nodename Matt \
       -ip 127.0.0.1 \
       -httpport 2002 \
       -tcpport 3002 \
       -netip 127.0.0.1 \
       -netport 3000
```

Each node has it's own wallet, so now you can send jeffCoins/Value.
To do this, use the webserver and API interface.

### WEBSERVER & REST API

The GUI for the three nodes you just created are,

[127.0.0.1:2000](http://127.0.0.1:2000/)
**/**
[127.0.0.1:2001](http://127.0.0.1:2001/)
**/**
[127.0.0.1:2002](http://127.0.0.1:2002/)

The main page will list the various API commands.
For example, to show a particular block,

[127.0.0.1:2000/showblock/0](http://127.0.0.1:2000/showblock/0)

### SWITCHES (REFERENCE)

  `-h` prints the following,
  
* `-gce`
  Is this Node on GCE
* `-genesis`
  Create your first Node
* `-httpport` _string_
  Node Web Port (default "2001")
* `-ip` string
  Node IP (default "127.0.0.1")
* `-loglevel` _string_
  LogLevel (info, debug or trace) (default "info")
* `-netip` _string_
  Network IP (default "127.0.0.1")
* `-netport` _string_
  Network TCP Port (default "3000")
* `-nodename` _string_
  Node Name (default "Jeff")
* `-tcpport` _string_
  Node TCP Port (default "3001")
* `-test`
  Loads the blockchain with test data (SEE BELOW)
* `-v`
  prints current version
* -`wallet`
  Just the wallet and gui/api

### JUST WALLET (OPTIONAL)

If you want to just have the wallet, use the `-wallet` switch.
You will not be part of the network since there is no blockchain
or miner.  But you can always restart and become part of the node.

You will need to hook up to a node, so the following could work,

```bash
go run jeffCoin.go \
       -loglevel debug \
       -nodename Jeff \
       -ip 127.0.0.1 \
       -httpport 2001 \
       -tcpport 3001 \
       -netip 127.0.0.1 \
       -netport 3000
       -wallet
```

Note the node name is the wallet name.

### CONNECT USING TCP (OPTIONAL)

You can also bypass the REST API and just open a connection to the TCP server itself,

```txt
netcat -q -1 127.0.0.1 3000
```

And request commands such as,

```txt
--- Waiting for command: SBC, BANN, SNL, BVB, BC, BTR, SAB, TR, EOF
SNL
[...nodeList...]
thank you
```

Notice you will need to handshake it with a `thank you` at the end.

There is a complete list of commands up above in
[TCP REQUESTS & HANDLERS](https://github.com/JeffDeCola/jeffCoin#32-tcp-requests--handlers).

### TEST MOCK TRANSACTIONS (OPTIONAL)

If you add the `-test` switch you will run some mock transactions from mock wallets.
Those wallets are located in `/wallets` and just used for testing.

You must use the MockFounders nodename,

```bash
go run jeffCoin.go \
       -loglevel debug \
       -genesis \
       -nodename MockFounders \
       -ip 127.0.0.1 \
       -httpport 2000 \
       -tcpport 3000 \
       -test
```

These transactions are the same I used in my
[bitcoin-ledger](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/bitcoin-ledger)
example.

So your blockchain and pendingBlock should look similar to
[blockchain-output.txt](https://github.com/JeffDeCola/my-go-examples/blob/master/blockchain/bitcoin-ledger/blockchain-output.txt).

And the balances in the blockchain should be,

```txt
The balance for MockFounders PubKey (Address) is 99657000
The balance for MockJeffs PubKey (Address) is 42500
The balance for MockMatts PubKey (Address) is 265000
The balance for MockJills PubKey (Address) is 35000
The balance for MockCoinVaults PubKey (Address) is 500
```

Remember, the pendingBlock is pending, so it's not part of this calculation.
Transaction do not have value if they are not part of the blockchain.

### RUN ON GOOGLE COMPUTE ENGINE (GCE) (OPTIONAL)

Make sure your create a firewall rule and have your instance use
it as a network tag,

```bash
gcloud compute firewall-rules create jeffs-firewall-settings-rule \
    --action allow \
    --rules tcp:1234,tcp:3334 \
    --priority 1000 \
    --source-ranges 0.0.0.0/0 \
    --target-tags "jeffs-firewall-settings" \
    --description "Jeffs firewall rules"
```

The IP `0.0.0.0` gets forwarded to your external IP, hence I added a
`-gce switch` to deal with this,

```bash
go run jeffCoin.go \
       -gce \
       -loglevel debug \
       -genesis \
       -nodename Founders \
       -ip 35.203.189.193 \
       -httpport 1234 \
       -tcpport 3334
```

Add another node (not at gce) with,

```bash
go run jeffCoin.go \
       -loglevel debug \
       -nodename Jeff \
       -ip 192.168.20.100 \
       -httpport 1235 \
       -tcpport 3335 \
       -netip 35.203.189.193 \
       -netport 3334
```

I have a gce build example
[here](https://github.com/JeffDeCola/my-packer-image-builds#jeffs-gce-ubuntu-1904-xxxx).

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
