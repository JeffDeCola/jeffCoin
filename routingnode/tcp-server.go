// jeffCoin 3. ROUTINGNODE tcp-server.go

package routingnode

import (
	"net"

	log "github.com/sirupsen/logrus"
)

// TCP SERVER ************************************************************************************************************

// StartRoutingNode - Start the Routing Node (TCP Server)
func StartRoutingNode(nodeIP string, nodeTCPPort string) {

	s := "START  StartRoutingNode() - Start the Routing Node (TCP Server)"
	log.Trace("ROUTINGNODE: SERVER " + s)

	s = "TCP  Server listening on " + nodeIP + ":" + nodeTCPPort
	log.Info("ROUTINGNODE: SERVER        " + s)

	// LISTEN ON IP AND PORT
	server, err := net.Listen("tcp", nodeIP+":"+nodeTCPPort)
	checkErr(err)
	defer server.Close()

	// CREATE A CONNECTION FOR EACH REQUEST
	// Serve connections concurrently
	for {

		// Wait for a connection request
		conn, err := server.Accept()
		checkErr(err)

		go HandleRequest(conn)
	}

}
