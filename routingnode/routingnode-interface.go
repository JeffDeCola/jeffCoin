// jeffCoin 3. ROUTINGNODE routingnode-interface.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

// NODELIST **************************************************************************************************************

// GetNodeList -  Gets the nodeList
func GetNodeList() nodeSlice {

	s := "START  GetNodeList() -  Gets the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	theNodeList := getNodeList()

	s = "END    GetNodeList() -  Gets the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	return theNodeList

}

// GenesisNodeList - Creates the nodeList
func GenesisNodeList() {

	s := "START  GenesisNodeList() - Creates the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	thisNode := appendThisNode()

	fmt.Printf("\nCongrats, your first Node in your nodeList is:\n\n")
	js, _ := json.MarshalIndent(thisNode, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END    GenesisNodeList() - Creates the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

}

// RequestNodeList -  Requests the nodeList from a Network Node
func RequestNodeList(networkIP string, networkTCPPort string) error {

	s := "START  RequestNodeList() -  Requests the nodeList from a Network Node"
	log.Trace("ROUTINGNODE: I/F      " + s)

	// CONN - SETUP THE CONNECTION
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: I/F             " + s)
	s = "CLIENT - Requesting a connection"
	log.Info("ROUTINGNODE: I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: I/F             " + s)
	s = "-C conn   TCP Connection on " + networkIP + ":" + networkTCPPort
	log.Info("ROUTINGNODE: I/F   " + s)
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("ROUTINGNODE: I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("ROUTINGNODE: I/F              " + s)
		return errors.New(s)
	}

	// REQ - SEND-NODELIST
	s = "-C req    - SEND-NODELIST"
	log.Info("ROUTINGNODE: I/F   " + s)
	fmt.Fprintf(conn, "SEND-NODELIST\n")

	// RCV - GET THE nodeList
	messageNodeList, _ := bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + messageNodeList
	log.Info("ROUTINGNODE: I/F   " + s)
	if messageNodeList == "ERROR" {
		s = "ERROR: Could not get nodeList from node"
		log.Error("ROUTINGNODE: I/F              " + s)
		return errors.New(s)
	}

	// LOAD THE nodeList
	loadNodeList(messageNodeList)

	// SEND - THANK YOU
	s = "-C send   - Thank you"
	log.Info("ROUTINGNODE: I/F   " + s)
	fmt.Fprintf(conn, "Thank You\n")

	// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "-C rcv    Message from Network Node: " + message
	log.Info("ROUTINGNODE: I/F   " + s)
	if message == "ERROR" {
		s = "ERROR: Waiting for Command"
		log.Trace("ROUTINGNODE: I/F              " + s)
		return errors.New(s)
	}

	// REQ - EOF (CLOSE CONNECTION)
	s = "-C req    - EOF (CLOSE CONNECTION)"
	log.Info("ROUTINGNODE: I/F   " + s)
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: I/F             " + s)
	s = "CLIENT - Closed a connection"
	log.Info("ROUTINGNODE: I/F             " + s)
	s = "----------------------------------------------------------------"
	log.Info("ROUTINGNODE: I/F             " + s)

	s = "END    RequestNodeList() -  Requests the nodeList from a Network Node"
	log.Trace("ROUTINGNODE: I/F      " + s)

	return nil

}

// NODE ******************************************************************************************************************

// GetNode - Gets a Node (via Index number) from the nodeList
func GetNode(id string) nodeStruct {

	s := "START  GetNode() - Gets a Node (via Index number) from the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	theNode := getNode(id)

	// RETURN NOT FOUND
	s = "END    GetNode() - Gets a Node (via Index number) from the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	return theNode

}

// AppendNewNode - Appends a new Node to the nodeList
func AppendNewNode(messageNewNode string) nodeStruct {

	s := "START  AppendNewNode() - Appends a new Node to the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	newNode := appendNewNode(messageNewNode)

	s = "END    AppendNewNode() - Appends a new Node to the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	return newNode

}

// THISNODE **************************************************************************************************************

// GetThisNode - Gets thisNode
func GetThisNode() nodeStruct {

	s := "START  GetThisNode() - Gets thisNode"
	log.Trace("ROUTINGNODE: I/F      " + s)

	theNode := getThisNode()

	s = "END    GetThisNode() - Gets thisNode"
	log.Trace("ROUTINGNODE: I/F      " + s)

	return theNode
}

// LoadThisNode - Loads thisNode
func LoadThisNode(ip string, httpPort string, tcpPort string, nodeName string, toolVersion string) {

	s := "START  LoadThisNode() - Loads thisNode"
	log.Trace("ROUTINGNODE: I/F      " + s)

	loadThisNode(ip, httpPort, tcpPort, nodeName, toolVersion)

	fmt.Printf("\nCongrats, you created your thisNode:\n\n")
	js, _ := json.MarshalIndent(thisNode, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END    LoadThisNode() - Loads thisNode"
	log.Trace("ROUTINGNODE: I/F      " + s)

}

// AppendThisNode - Appends thisNode to the nodeList
func AppendThisNode() {

	s := "START  AppendThisNode() - Appends thisNode to the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

	// DO YOU ALREADY HAVE thisNode IN THE nodeList?
	if !checkIfThisNodeinNodeList() {

		_ = appendThisNode()

	}

	s = "END    AppendThisNode() - Appends thisNode to the nodeList"
	log.Trace("ROUTINGNODE: I/F      " + s)

}

// BroadcastThisNode - Broadcasts thisNode to the Network
func BroadcastThisNode() error {

	s := "START  BroadcastThisNode() - Broadcasts thisNode to the Network"
	log.Trace("ROUTINGNODE: I/F      " + s)

	// DO YOU ALREADY HAVE thisNode IN THE nodeList?
	if checkIfThisNodeinNodeList() {

		s = "YOU ALREADY HAVE thisNode IN THE nodeList"
		log.Info("ROUTINGNODE: I/F             " + s)

		s = "END    BroadcastThisNode() - Broadcasts thisNode to the Network"
		log.Trace("ROUTINGNODE: I/F      " + s)

		return nil

	}

	s = "thisNode IN NOT THE nodeList"
	log.Info("ROUTINGNODE: I/F             " + s)

	theNodeList := getNodeList()

	// FOR EACH NODE IN NODELIST
	for _, item := range theNodeList {

		networkIP := item.IP
		networkTCPPort := item.TCPPort

		//  CONN - SETUP THE CONNECTION
		s = "----------------------------------------------------------------"
		log.Info("ROUTINGNODE: I/F             " + s)
		s = "CLIENT - Requesting a connection"
		log.Info("ROUTINGNODE: I/F             " + s)
		s = "----------------------------------------------------------------"
		log.Info("ROUTINGNODE: I/F             " + s)
		s = "-C conn   TCP Connection on " + networkIP + ":" + networkTCPPort
		log.Info("ROUTINGNODE: I/F   " + s)
		conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
		if err != nil {
			// The connection is down - Skip this node
			s := "ERROR - NODE DOWN (SKIP) " + networkIP + ":" + networkTCPPort
			log.Warn("ROUTINGNODE: I/F              " + s)
			continue
		}

		// RCV - GET THE RESPONSE MESSAGE (Waiting for command)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		s = "-C rcv    Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F   " + s)
		if message == "ERROR" {
			s = "ERROR: Waiting for command"
			log.Error("ROUTINGNODE: I/F              " + s)
			return errors.New(s)
		}

		// REQ - BROADCAST-ADD-NEW-NODE
		s = "-C req    - BROADCAST-ADD-NEW-NODE"
		log.Info("ROUTINGNODE: I/F   " + s)
		fmt.Fprintf(conn, "BROADCAST-ADD-NEW-NODE\n")

		// RCV - GET THE RESPONSE MESSAGE (ASKING TO SEND thisNode)
		message, _ = bufio.NewReader(conn).ReadString('\n')
		s = "-C rcv    Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F   " + s)
		if message == "ERROR" {
			s = "ERROR: ASKING TO SEND thisNode"
			log.Error("ROUTINGNODE: I/F              " + s)
			return errors.New(s)
		}

		// SEND - SEND thisNode
		s = "-C send   SEND thisNode"
		log.Info("ROUTINGNODE: I/F   " + s)
		thisNode := getThisNode()
		js, _ := json.Marshal(thisNode)
		s = string(js)
		fmt.Fprintf(conn, s+"\n")

		// RCV - GET THE RESPONSE MESSAGE (AppendedNode to the Nodelist)
		message, _ = bufio.NewReader(conn).ReadString('\n')
		s = "-C rcv    Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F   " + s)
		if message == "ERROR" {
			s = "ERROR: AppendedNode to the Nodelist"
			log.Trace("ROUTINGNODE: I/F              " + s)
			return errors.New(s)
		}

		// SEND - THANK YOU
		s = "-C send   - Thank you"
		log.Info("ROUTINGNODE: I/F   " + s)
		fmt.Fprintf(conn, "Thank You\n")

		// RCV - GET THE RESPONSE MESSAGE (Waiting for Command)
		message, _ = bufio.NewReader(conn).ReadString('\n')
		s = "-C rcv    Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F   " + s)
		if message == "ERROR" {
			s = "ERROR: Waiting for Command"
			log.Trace("ROUTINGNODE: I/F              " + s)
			return errors.New(s)
		}

		// REQ - EOF (CLOSE CONNECTION)
		s = "-C req    - EOF (CLOSE CONNECTION)"
		log.Info("ROUTINGNODE: I/F   " + s)
		fmt.Fprintf(conn, "EOF\n")
		time.Sleep(2 * time.Second)
		conn.Close()
		s = "----------------------------------------------------------------"
		log.Info("ROUTINGNODE: I/F             " + s)
		s = "CLIENT - Closed a connection"
		log.Info("ROUTINGNODE: I/F             " + s)
		s = "----------------------------------------------------------------"
		log.Info("ROUTINGNODE: I/F             " + s)

	}

	s = "END    BroadcastThisNode() - Broadcasts thisNode to the Network"
	log.Trace("ROUTINGNODE: I/F      " + s)

	return nil

}
