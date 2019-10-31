// jeffCoin routingnode-interface.go

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

// NODELIST ************************************************************************************************************

// GenesisNodeList - Creates the nodeList
func GenesisNodeList() {

	s := "START: GenesisNodeList - Creates the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	thisNode := appendThisNode()

	fmt.Printf("\nCongrats, your first Node in your nodeList is:\n\n")
	js, _ := json.MarshalIndent(thisNode, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END:   GenesisNodeList - Creates the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

}

// LoadNodeList -  Receives the nodeList from a Network Node
func LoadNodeList(networkIP string, networkTCPPort string) error {

	s := "START: LoadNewNode -  Receives the nodeList from a Network Node"
	log.Trace("ROUTINGNODE: I/F    " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("ROUTINGNODE: I/F           " + s)
	if message == "ERROR" {
		s = "ERROR: Could not setup connection"
		log.Trace("ROUTINGNODE: I/F            " + s)
		return errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "SEND-NODELIST\n")

	// GET THE nodeList
	messageNodeList, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Info("ROUTINGNODE: I/F           " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Error("ROUTINGNODE: I/F            " + s)
		return errors.New(s)
	}

	// LOAD THE nodeList
	loadNodeList(messageNodeList)

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END:   LoadNodeList -  Receives the nodeList from a Network Node"
	log.Trace("ROUTINGNODE: I/F    " + s)

	return nil

}

// GetNodeList -  Gets the nodeList
func GetNodeList() nodeSlice {

	s := "START: GetNodeList -  Gets the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	theNodeList := getNodeList()

	s = "END:   GetNodeList -  Gets the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	return theNodeList

}

// NODE ****************************************************************************************************************

// GetNode - Gets a Node (via Index number) from the nodeList
func GetNode(id string) nodeStruct {

	s := "START: GetNode - Gets a Node (via Index number) from the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	theNode := getNode(id)

	// RETURN NOT FOUND
	s = "END:   GetNode - Gets a Node (via Index number) from the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	return theNode

}

// AppendNewNode - Appends a New Node to the nodeList
func AppendNewNode(messageNewNode string) nodeStruct {

	s := "START: AppendNode - Appends a New Node to the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	newNode := appendNewNode(messageNewNode)

	s = "END:   AppendNewNode - Appends a New Node to the nodeList"
	log.Trace("ROUTINGNODE: I/F    " + s)

	return newNode

}

// THISNODE ************************************************************************************************************

// LoadThisNode - Loads thisNode
func LoadThisNode(ip string, tcpPort string) {

	s := "START: LoadThisNode - Loads thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	loadThisNode(ip, tcpPort)

	s = "END:   LoadThisNode - Loads thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

}

// GetThisNode - Gets thisNode
func GetThisNode() nodeStruct {

	s := "START: GetThisNode - Gets thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	theNode := getThisNode()

	s = "END:   GetThisNode - Gets thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return theNode
}

// AppendThisNode - Appends thisNode to the nodeList
func AppendThisNode() {

	s := "START: AppendThisNode - Appends thisNode to the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	// DO YOU ALREADY HAVE thisNode IN THE nodeList?
	if !checkIfThisNodeinNodeList() {

		_ = appendThisNode()

	}

	s = "END:   AppendThisNode - Appends thisNode to the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

}

// BroadcastThisNode - Broadcasts thisNode to the Network
func BroadcastThisNode() error {

	s := "START: BroadcastThisNode - Broadcasts thisNode to the Network"
	log.Trace("ROUTINGNODE: I/F    " + s)

	// DO YOU ALREADY HAVE thisNode IN THE nodeList?
	if checkIfThisNodeinNodeList() {

		s = "END:   BroadcastThisNode - Broadcasts thisNode to the Network"
		log.Trace("ROUTINGNODE: I/F    " + s)

		return nil

	}

	theNodeList := getNodeList()

	// FOR EACH NODE IN NODELIST
	for _, item := range theNodeList {

		networkIP := item.IP
		networkTCPPort := item.Port

		// SETUP THE CONNECTION
		conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
		if err != nil {
			// The connection is down - Skip this node
			s := "ERROR - NODE DOWN (SKIP) " + networkIP + ":" + networkTCPPort
			log.Warn("ROUTINGNODE: I/F            " + s)
			continue
		}

		// GET THE RESPONSE MESSAGE
		message, _ := bufio.NewReader(conn).ReadString('\n')
		s = "Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F           " + s)
		if message == "ERROR" {
			s = "ERROR: Could not setup connection"
			log.Error("ROUTINGNODE: I/F            " + s)
			return errors.New(s)
		}

		// SEND THE REQUEST
		fmt.Fprintf(conn, "BROADCAST-ADD-NEW-NODE\n")

		// GET THE RESPONSE MESSAGE
		message, _ = bufio.NewReader(conn).ReadString('\n')
		s = "Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F           " + s)
		if message == "ERROR" {
			s = "ERROR: Could not get blockchain from node"
			log.Error("ROUTINGNODE: I/F            " + s)
			return errors.New(s)
		}

		// SEND NODE
		thisNode := getThisNode()
		js, _ := json.Marshal(thisNode)
		s = string(js)
		fmt.Fprintf(conn, s+"\n")

		// GET THE RESPONSE MESSAGE
		message, _ = bufio.NewReader(conn).ReadString('\n')
		s = "Message from Network Node: " + message
		log.Info("ROUTINGNODE: I/F           " + s)
		if message == "ERROR" {
			s = "ERROR: Could not get blockchain from node"
			log.Error("ROUTINGNODE: I/F            " + s)
			return errors.New(s)
		}

		// CLOSE CONNECTION
		fmt.Fprintf(conn, "EOF\n")
		time.Sleep(2 * time.Second)
		conn.Close()

	}

	s = "END:   BroadcastThisNode - Broadcasts thisNode to the Network"
	log.Trace("ROUTINGNODE: I/F    " + s)

	return nil

}
