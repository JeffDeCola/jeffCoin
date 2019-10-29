// jeffCoin requests.go

package routingnode

import (
	"bufio"
	"io"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

// HandleRequest handles TCP requests
func HandleRequest(conn net.Conn) {

	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	s := "Opening a connection"
	log.Info("ROUTINGNODE: REQ           " + s)
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: REQ           " + s)

	// READ FROM CONN UTIL EOF
	for {

		s := "Waiting for command: SENDBLOCKCHAIN (SB), ADDNEWNODE (NN), SENDNODELIST (SN), " +
			"SENDADDRESSBALANCE (SAB), TRANSACTIONREQUEST (TR),  BROADCASTTRANSACTIONREQUEST (BTR) or EOF"
		returnMessage(s, rw)

		cmd, err := rw.ReadString('\n')
		//checkErr(err)
		// TRIM CMD
		cmd = strings.Trim(cmd, "\n ")

		s = "Received command and working on it: " + cmd
		log.Info("ROUTINGNODE: REQ           " + s)

		// CHECK FOR EOF
		switch {
		case err == io.EOF:
			s = "Reached EOF"
			log.Info("ROUTINGNODE: REQ           " + s)
			s = "Closing this connection"
			log.Info("ROUTINGNODE: REQ           " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ           " + s)
			return
		case err != nil:
			s = "ERROR reading command. Got: " + cmd
			log.Info("ROUTINGNODE: REQ           " + s)
			return
		}

		// CALL HANDLER
		// ADDNEWBLOCK
		// Otherwise close connection
		switch {
		// FROM BLOCKCHAIN INTERFACE *******************************
		case cmd == "SENDBLOCKCHAIN" || cmd == "SBC":
			handleSendBlockchain(rw)
		// FROM ROUTINGNODE INTERFACE ******************************
		case cmd == "BCASTADDNEWNODE" || cmd == "BANN":
			handleBroadCastAddNewNode(rw)
		case cmd == "SENDNODELIST" || cmd == "SN":
			handleSendNodeList(rw)
		case cmd == "BCASTVERIFIEDBLOCK" || cmd == "BVB":
			handleBroadCastVerifiedBlock(rw)
		case cmd == "BCASTCONSENSUS" || cmd == "BC":
			handleBroadCastConsensus(rw)
		case cmd == "BCASTTRANSACTIONREQUEST" || cmd == "BTR":
			handleBroadCastTransactionRequest(rw)
		// FROM WALLET INTERFACE ***********************************
		case cmd == "SENDADDRESSBALANCE" || cmd == "SAB":
			handleSendAddressBalance(rw)
		case cmd == "TRANSACTIONREQUEST" || cmd == "TR":
			handleTransactionRequest(rw)

		case cmd == "EOF":
			s = "Received EOF"
			log.Info("ROUTINGNODE: REQ           " + s)
			s = "Closing this connection"
			log.Info("ROUTINGNODE: REQ           " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ           " + s)
			return
		default:
			s = "Did not get correct command. Received: " + cmd
			log.Info("ROUTINGNODE: REQ           " + s)
			s = "Closing this connection"
			log.Info("ROUTINGNODE: REQ           " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ           " + s)
			return
		}
	}
}

func returnMessage(s string, rw *bufio.ReadWriter) {
	log.Info("ROUTINGNODE: REQ           " + s)
	_, err := rw.WriteString("--- " + s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
}
