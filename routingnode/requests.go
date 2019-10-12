// jeffCoin requests.go

package routingnode

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

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

		s := "Waiting for command: ADDNEWBLOCK (AB), ADDTRANSACTION (AT), SENDBLOCKCHAIN (SB) or EOF"
		returnMessage(s, rw)

		cmd, err := rw.ReadString('\n')
		// TRIM CMD
		cmd = strings.Trim(cmd, "\n ")

		s = "Received command and working on it: " + cmd
		returnMessage(s, rw)

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
		case cmd == "ADDBLOCK" || cmd == "AB":
			handleAddBlock(rw)
		case cmd == "ADDTRANSACTION" || cmd == "AT":
			handleAddTransaction(rw)
		case cmd == "SENDBLOCKCHAIN" || cmd == "SB":
			handleSendBlockchain(rw)
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
