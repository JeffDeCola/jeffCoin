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
	log.Info("ROUTINGNODE: REQ            " + s)
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: REQ            " + s)

	// READ FROM CONN UTIL EOF
	for {

		s := "Waiting for command: ADDTRANSACTION (AT), SENDBLOCKCHAIN (SB), ADDNEWNODE (NN). SENDNODELIST (SN) or EOF"
		returnMessage(s, rw)

		cmd, err := rw.ReadString('\n')
		//checkErr(err)
		// TRIM CMD
		cmd = strings.Trim(cmd, "\n ")

		s = "Received command and working on it: " + cmd
		log.Info("ROUTINGNODE: REQ            " + s)

		// CHECK FOR EOF
		switch {
		case err == io.EOF:
			s = "Reached EOF"
			log.Info("ROUTINGNODE: REQ            " + s)
			s = "Closing this connection"
			log.Info("ROUTINGNODE: REQ            " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ            " + s)
			return
		case err != nil:
			s = "ERROR reading command. Got: " + cmd
			log.Info("ROUTINGNODE: REQ            " + s)
			return
		}

		// CALL HANDLER
		// ADDNEWBLOCK
		// Otherwise close connection
		switch {
		case cmd == "ADDTRANSACTION" || cmd == "AT":
			handleAddTransaction(rw)
		case cmd == "SENDBLOCKCHAIN" || cmd == "SB":
			handleSendBlockchain(rw)
		case cmd == "ADDNEWNODE" || cmd == "NN":
			handleAddNewNode(rw)
		case cmd == "SENDNODELIST" || cmd == "SN":
			handleSendNodeList(rw)
		case cmd == "EOF":
			s = "Received EOF"
			log.Info("ROUTINGNODE: REQ            " + s)
			s = "Closing this connection"
			log.Info("ROUTINGNODE: REQ            " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ            " + s)
			return
		default:
			s = "Did not get correct command. Received: " + cmd
			log.Info("ROUTINGNODE: REQ            " + s)
			s = "Closing this connection"
			log.Info("ROUTINGNODE: REQ            " + s)
			s = "----------------------------------------------------------------"
			log.Info("ROUTINGNODE: REQ            " + s)
			return
		}
	}
}

func returnMessage(s string, rw *bufio.ReadWriter) {
	log.Info("ROUTINGNODE: REQ            " + s)
	_, err := rw.WriteString("--- " + s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
}
