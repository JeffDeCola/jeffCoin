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
	log.Println("ROUTINGNODE:    " + s)
	s = "----------------------------------------------------------------"
	log.Println("ROUTINGNODE:    " + s)

	// READ FROM CONN UTIL EOF
	for {

		s := "Waiting for command: ADDTRANSACTION (AT), SENDBLOCKCHAIN (SB), ADDNEWNODE (NN). SENDNODELIST (SN) or EOF"
		returnMessage(s, rw)

		cmd, err := rw.ReadString('\n')
		// TRIM CMD
		cmd = strings.Trim(cmd, "\n ")

		s = "Received command and working on it: " + cmd
		log.Println("ROUTINGNODE:    " + s)

		// CHECK FOR EOF
		switch {
		case err == io.EOF:
			s = "Reached EOF"
			log.Println("ROUTINGNODE:    " + s)
			s = "Closing this connection"
			log.Println("ROUTINGNODE:    " + s)
			s = "----------------------------------------------------------------"
			log.Println("ROUTINGNODE:    " + s)
			return
		case err != nil:
			s = "ERROR reading command. Got: " + cmd
			log.Println("ROUTINGNODE:    " + s)
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
			log.Println("ROUTINGNODE:    " + s)
			s = "Closing this connection"
			log.Println("ROUTINGNODE:    " + s)
			s = "----------------------------------------------------------------"
			log.Println("ROUTINGNODE:    " + s)
			return
		default:
			s = "Did not get correct command. Received: " + cmd
			log.Println("ROUTINGNODE:    " + s)
			s = "Closing this connection"
			log.Println("ROUTINGNODE:    " + s)
			s = "----------------------------------------------------------------"
			log.Println("ROUTINGNODE:    " + s)
			return
		}
	}
}

func returnMessage(s string, rw *bufio.ReadWriter) {
	log.Println("ROUTINGNODE:    " + s)
	_, err := rw.WriteString("--- " + s + "\n")
	checkErr(err)
	err = rw.Flush()
	checkErr(err)
}
