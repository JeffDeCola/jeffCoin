// jeffCoin 3. ROUTINGNODE requests.go

package routingnode

import (
	"bufio"
	"io"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

// REQUESTS **************************************************************************************************************

// HandleRequest handles TCP requests
func HandleRequest(conn net.Conn) {

	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	s := "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: REQ             " + s)
	s = "-S open   SERVER - Opened a connection"
	log.Info("ROUTINGNODE: REQ   " + s)
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: REQ             " + s)

	// READ FROM CONN UTIL EOF
	for {

		s := "Waiting for command: " +
			"SBC, BANN, SNL, BVB, BC, BTR, SAB, TR, EOF"
		returnMessage(s, rw)

		// RCVD - RECEIVED - COMMAND
		cmd, err := rw.ReadString('\n')
		checkErr(err)
		// TRIM CMD
		cmd = strings.Trim(cmd, "\n ")
		s = "-S rcvd   - " + cmd
		log.Info("ROUTINGNODE: REQ   " + s)

		// CHECK FOR EOF
		switch {
		case err == io.EOF:
			s = "Reached EOF"
			log.Info("ROUTINGNODE: REQ             " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ             " + s)
			s = "SERVER - Closing this connection"
			log.Info("ROUTINGNODE: REQ             " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ             " + s)
			return
		case err != nil:
			s = "ERROR reading command. Got: " + cmd
			log.Error("ROUTINGNODE: REQ           " + s)
			return
		}

		// CALL HANDLER
		// ADDNEWBLOCK
		// Otherwise close connection
		switch {
		// FROM BLOCKCHAIN I/F *************************************
		case cmd == "SEND-BLOCKCHAIN" || cmd == "SBC":
			handleSendBlockchain(rw)
		// FROM ROUTINGNODE I/F ************************************
		case cmd == "BROADCAST-ADD-NEW-NODE" || cmd == "BANN":
			handleBroadcastAddNewNode(rw)
		case cmd == "SEND-NODELIST" || cmd == "SNL":
			handleSendNodeList(rw)
		case cmd == "BROADCAST-VERIFIED-BLOCK" || cmd == "BVB":
			handleBroadcastVerifiedBlock(rw)
		case cmd == "BROADCAST-CONSENSUS" || cmd == "BC":
			handleBroadcastConsensus(rw)
		case cmd == "BROADCAST-TRANSACTION-REQUEST" || cmd == "BTR":
			handleBroadcastTransactionRequest(rw)
		// FROM WALLET I/F *****************************************
		case cmd == "SEND-ADDRESS-BALANCE" || cmd == "SAB":
			handleSendAddressBalance(rw)
		case cmd == "TRANSACTION-REQUEST" || cmd == "TR":
			handleTransactionRequest(rw)
		// EOF *****************************************************
		case cmd == "EOF":
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ             " + s)
			s = "-S close  SERVER - Closing this connection"
			log.Info("ROUTINGNODE: REQ   " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ             " + s)
			return
		default:
			s = "Did not get correct command. Received: " + cmd
			log.Warn("ROUTINGNODE: REQ             " + s)
			s = "----------------------------------------------------------------"
			log.Warn("ROUTINGNODE: REQ             " + s)
			s = "-S close  SERVER - Closing this connection"
			log.Warn("ROUTINGNODE: REQ   " + s)
			s = "----------------------------------------------------------------"
			log.Warn("ROUTINGNODE: REQ             " + s)
			return
		}
	}
}

// returnMessage - Returns message
func returnMessage(s string, rw *bufio.ReadWriter) {
	_, err := rw.WriteString("--- " + s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
	log.Info("ROUTINGNODE: REQ   -S sent   " + s)
}
